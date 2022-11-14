package complete

import (
	"fmt"

	IntervalSet "github.com/Flo-Leroux/quanty-lang/pkg/complete/interval-set"
	list "github.com/Flo-Leroux/quanty-lang/pkg/complete/list"
	mutablelist "github.com/Flo-Leroux/quanty-lang/pkg/complete/mutable-list"
	mutablemap "github.com/Flo-Leroux/quanty-lang/pkg/complete/mutable-map"
	mutableset "github.com/Flo-Leroux/quanty-lang/pkg/complete/mutable-set"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

var followSetsByATN *mutablemap.MutableMap[string, FollowSetsPerState] = mutablemap.New[string, FollowSetsPerState]()

type Vocabulary interface {
	GetDisplayName(index int) string
}

type CodeCompletionCore struct {
	atn                 *antlr.ATN
	vocabulary          Vocabulary
	ruleNames           []string
	languageName        string
	predicatesEvaluator antlr.Recognizer

	// Not dependent on showDebugOutput. Prints the collected rules + tokens to terminal.
	showResult bool
	// Enables printing ATN state info to terminal.
	showDebugOutput bool
	// Only relevant when showDebugOutput is true. Enables transition printing for a state.
	debugOutputWithTransitions bool
	// Also depends on showDebugOutput. Enables call stack printing for each rule recursion.
	showRuleStack bool

	// Tokens which should not appear in the candidates set.
	ignoredTokens []int
	// Rules which replace any candidate token they contain.
	preferredRules []int

	tokens          TokenList
	statesProcessed int // Used only for debugging
	tokensProvider  TokensProvider

	// A mapping of rule index + token stream position to end token positions.
	// A rule which has been visited before with the same input position will always produce the same output positions.
	shortcutMap *mutablemap.MutableMap[int, *mutablemap.MutableMap[int, RuleEndStatus]]
	candidates  *CandidatesCollection

	myl *list.List[string]
}

func FromParser(name string, parser antlr.Parser, vocabulary Vocabulary) *CodeCompletionCore {
	return NewCodeCompletionCore(
		parser.GetATN(),
		vocabulary,
		parser.GetRuleNames(),
		name,
		parser,
	)
}

func NewCodeCompletionCore(
	atn *antlr.ATN,
	vocabulary Vocabulary,
	ruleNames []string,
	languageName string,
	predicatesEvaluator antlr.Recognizer,
) *CodeCompletionCore {
	return &CodeCompletionCore{
		atn:                 atn,
		vocabulary:          vocabulary,
		ruleNames:           ruleNames,
		languageName:        languageName,
		predicatesEvaluator: predicatesEvaluator,

		showResult:                 false,
		showDebugOutput:            false,
		debugOutputWithTransitions: false,
		showRuleStack:              false,

		ignoredTokens:  make([]int, 0),
		preferredRules: make([]int, 0),

		tokens:          NewTokenList(),
		statesProcessed: 0,
		tokensProvider:  nil,

		shortcutMap: mutablemap.New[int, *mutablemap.MutableMap[int, RuleEndStatus]](),
		candidates:  NewCandidatesCollection(),

		myl: list.New(func(a, b string) bool { return a == b }),
	}
}

func (c *CodeCompletionCore) EnableDebug() {
	c.showDebugOutput = true
	c.showResult = true
	c.debugOutputWithTransitions = true
	c.showRuleStack = true
}

func (c *CodeCompletionCore) ShowResult() {
	c.showResult = true
}

func (c *CodeCompletionCore) SetPreferredRules(rules ...int) {
	c.preferredRules = append(c.preferredRules, rules...)
}

func (c *CodeCompletionCore) SetIgnoredTokens(tokens ...int) {
	c.ignoredTokens = append(c.ignoredTokens, tokens...)
}

func (c *CodeCompletionCore) CollectCandidates(
	tokenStream antlr.TokenStream,
	caretTokenIndex int,
	context antlr.ParserRuleContext,
) *CandidatesCollection {
	return c.collectCandidates(NewByStreamTokenProvider(tokenStream, caretTokenIndex, context))
}

/**
* This is the main entry point. The caret token index specifies the token stream index for the token which currently
* covers the caret (or any other position you want to get code completion candidates for).
* Optionally you can pass in a parser rule context which limits the ATN walk to only that or called rules. This can significantly
* speed up the retrieval process but might miss some candidates (if they are outside of the given context).
 */
func (c *CodeCompletionCore) collectCandidates(tokensProvider TokensProvider) *CandidatesCollection {
	c.tokensProvider = tokensProvider
	c.shortcutMap.Clear()
	c.candidates.Clear()
	c.statesProcessed = 0

	c.tokens = tokensProvider.Tokens()

	callStack := mutablelist.New(func(a, b int) bool { return a == b })
	startRule := tokensProvider.StartRuleIndex()

	ruleToStartState := GetUnexportedField[*antlr.ATN, []*antlr.RuleStartState](c.atn, "ruleToStartState")
	c.processRule(ruleToStartState[startRule], 0, callStack, "", nil)

	if c.showResult {
		fmt.Printf("States processed: %d", c.statesProcessed)
		c.candidates.Show(c.ruleNames, c.vocabulary)
	}

	return c.candidates
}

/**
* Check if the predicate associated with the given transition evaluates to true.
 */
func (c *CodeCompletionCore) checkPredicate(transition *antlr.PredicateTransition) bool {
	panic("Method Not Implemented")
}

/**
* Walks the rule chain upwards to see if that matches any of the preferred rules.
* If found, that rule is added to the collection candidates and true is returned.
 */
func (c *CodeCompletionCore) translateToRuleIndex(ruleStack RuleList) bool {
	if len(c.preferredRules) == 0 {
		return false
	}

	// Loop over the rule stack from highest to lowest rule level. This way we properly handle the higher rule
	// if it contains a lower one that is also a preferred rule.
	// ruleStack
	for i := 0; i <= ruleStack.Len(); i++ {
		if Contains(c.preferredRules, ruleStack.Get(i)) {
			// Add the rule to our candidates list along with the current rule path,
			// but only if there isn't already an entry like that.
			path := ruleStack.Slice(0, i)
			addNew := true
			c.candidates.rules.Range(func(key int, value RuleList) bool {
				if key != ruleStack.Get(i) || value.Len() != path.Len() {
					return true
				}

				// Found an entry for this rule. Same path? If so don't add a new (duplicate) entry.
				found := false
				for j, v := range path.ToArray() {
					if v == value.Get(j) {
						found = true
						break
					}
				}

				if found {
					addNew = false
					return false
				}

				return true
			})

			if addNew {
				c.candidates.rules.Set(ruleStack.Get(i), path)

				if c.showDebugOutput {
					c.myl.Add(c.ruleNames[i])
					fmt.Printf("=====> collected: %s", c.ruleNames[i])

					counter := c.myl.Filter(func(s string) bool {
						return s == c.ruleNames[i]
					})

					if counter.Len() > 10 {
						panic("LOOP")
					}
				}
			}

			return true
		}
	}

	return false
}

/**
* This method follows the given transition and collects all symbols within the same rule that directly follow it
* without intermediate transitions to other rules and only if there is a single symbol for a transition.
 */
func (c *CodeCompletionCore) getFollowingTokens(transition antlr.Transition) TokenList {
	result := NewTokenList()

	pipeline := list.New(func(a, b antlr.ATNState) bool { return a.Equals(b) })

	target := GetUnexportedField[antlr.Transition, antlr.ATNState](transition, "target")
	pipeline.Add(target)

	for pipeline.Len() > 0 {
		state := pipeline.Pop()

		for _, t := range state.GetTransitions() {
			switch tr := t.(type) {
			case *antlr.AtomTransition:
				isEpsilon := GetUnexportedField[*antlr.AtomTransition, bool](tr, "isEpsilon")
				target = GetUnexportedField[antlr.Transition, antlr.ATNState](tr, "target")

				if !isEpsilon {
					label := GetUnexportedField[*antlr.AtomTransition, *antlr.IntervalSet](tr, "intervalSet")

					list := IntervalSet.FromAntlr(label).ToList()
					if len(list) == 1 && !Contains(c.ignoredTokens, list[0]) {
						result.Add(list[0])
						pipeline.Add(target)
					}
				} else {
					pipeline.Add(target)
				}
			}
		}
	}

	return result
}

/**
* Entry point for the recursive follow set collection function.
 */
func (c *CodeCompletionCore) determineFollowSets(start, stop antlr.ATNState) FollowSetWithPathList {

	result := NewFollowSetWithPathList()
	seen := mutableset.New(func(a, b antlr.ATNState) bool { return a.Equals(b) })
	ruleStack := mutablelist.New(func(a, b int) bool { return a == b })

	c.collectFollowSets(start, stop, result, seen, ruleStack)

	return result
}

/**
* Collects possible tokens which could be matched following the given ATN state. This is essentially the same
* algorithm as used in the LL1Analyzer class, but here we consider predicates also and use no parser rule context.
 */
func (c *CodeCompletionCore) collectFollowSets(
	s, stopState antlr.ATNState,
	followSets FollowSetWithPathList,
	seen *mutableset.MutableSet[antlr.ATNState],
	ruleStack *mutablelist.MutableList[int],
) {
	if seen.Contains(s) {
		return
	}

	seen.Add(s)

	if s.Equals(stopState) || s.GetStateType() == antlr.ATNStateRuleStop {
		set := NewFollowSetWithPath()
		set.intervals = IntervalSet.Of(antlr.TokenEpsilon)
		set.path.AddList(ruleStack)
		followSets.Add(set)
		return
	}

	for _, transition := range s.GetTransitions() {
		serializationType := GetUnexportedField[antlr.Transition, int](transition, "serializationType")
		target := GetUnexportedField[antlr.Transition, antlr.ATNState](transition, "target")
		isEpsilon := GetUnexportedField[antlr.Transition, bool](transition, "isEpsilon")
		maxTokenType := GetUnexportedField[*antlr.ATN, int](c.atn, "maxTokenType")

		if serializationType == antlr.TransitionRULE {
			if ruleStack.IndexOf(target.GetRuleIndex()) != -1 {
				continue
			}

			ruleStack.Add(target.GetRuleIndex())
			c.collectFollowSets(target, stopState, followSets, seen, ruleStack)
			ruleStack.Pop()
		} else if serializationType == antlr.TransitionPREDICATE {
			if c.checkPredicate(transition.(*antlr.PredicateTransition)) {
				c.collectFollowSets(target, stopState, followSets, seen, ruleStack)
			}
		} else if isEpsilon {
			c.collectFollowSets(target, stopState, followSets, seen, ruleStack)
		} else if serializationType == antlr.TransitionWILDCARD {
			set := NewFollowSetWithPath()
			set.intervals = IntervalSet.FromStartStop(antlr.TokenMinUserTokenType, maxTokenType)
			set.path.AddList(ruleStack)
			followSets.Add(set)
		} else {
			label := GetUnexportedField[antlr.Transition, *antlr.IntervalSet](transition, "intervalSet")
			myLabel := IntervalSet.FromAntlr(label)
			if len(myLabel.GetIntervals()) != 0 {
				if serializationType == antlr.TransitionNOTSET {
					myLabel = myLabel.Complement(antlr.TokenMinUserTokenType, maxTokenType)
				}

				set := NewFollowSetWithPath()
				set.intervals = myLabel
				set.path.AddList(ruleStack)
				set.following = c.getFollowingTokens(transition)
				followSets.Add(set)
			}
		}
	}
}

func (c *CodeCompletionCore) setsPerState() FollowSetsPerState {
	setsPerState := followSetsByATN.Get(c.languageName)

	if setsPerState == nil {
		setsPerState = mutablemap.New[int, *FollowSetsHolder]()
		followSetsByATN.Set(c.languageName, setsPerState)
	}

	return setsPerState
}

func (c *CodeCompletionCore) followSets(startState antlr.ATNState) *FollowSetsHolder {
	setsPerState := c.setsPerState()
	followSets := setsPerState.Get(startState.GetStateNumber())
	if followSets == nil {
		followSets = NewFollowSetsHolder()
		setsPerState.Set(startState.GetStateNumber(), followSets)

		ruleToStopState := GetUnexportedField[*antlr.ATN, []*antlr.RuleStopState](c.atn, "ruleToStopState")
		stop := ruleToStopState[startState.GetRuleIndex()]
		followSets.sets = c.determineFollowSets(startState, stop)

		// Sets are split by path to allow translating them to preferred rules. But for quick hit tests
		// it is also useful to have a set with all symbols combined.
		combined := IntervalSet.NewIntervalSet()
		for _, set := range followSets.sets.ToArray() {
			// combined.AddAll(IntervalSet.OfAntlr(set.intervals))
			combined.AddSet(set.intervals)
		}
		followSets.combined = combined
	}
	return followSets
}

/**
* Walks the ATN for a single rule only. It returns the token stream position for each path that could be matched in this rule.
* The result can be empty in case we hit only non-epsilon transitions that didn't match the current input or if we
* hit the caret position.
 */
func (c *CodeCompletionCore) processRule(
	startState antlr.ATNState,
	tokenIndex int,
	callStack *mutablelist.MutableList[int],
	_indentation string,
	pathDescription *PathDescription,
) RuleEndStatus {

	if pathDescription == nil {
		pathDescription = NewPathDescription(startState, nil)
	}

	//println("PROCESSRULE state=${startState} tokenIndex=$tokenIndex callStack=$callStack")
	indentation := _indentation
	// Start with rule specific handling before going into the ATN walk.

	// Check first if we've taken this path with the same input before.
	positionMap := c.shortcutMap.Get(startState.GetRuleIndex())
	if positionMap == nil {
		positionMap = mutablemap.New[int, RuleEndStatus]()
		c.shortcutMap.Set(startState.GetRuleIndex(), positionMap)
	} else {
		if positionMap.Contains(tokenIndex) {
			if c.showDebugOutput {
				fmt.Println("=====> shortcut")
			}

			return positionMap.Get(tokenIndex)
		}
	}

	result := NewRuleEndStatus()

	// For rule start states we determine and cache the follow set, which gives us 3 advantages:
	// 1) We can quickly check if a symbol would be matched when we follow that rule. We can so check in advance
	//    and can save us all the intermediate steps if there is no match.
	// 2) We'll have all symbols that are collectable already together when we are at the caret when entering a rule.
	// 3) We get this lookup for free with any 2nd or further visit of the same rule, which often happens
	//    in non trivial grammars, especially with (recursive) expressions and of course when invoking code completion
	//    multiple times.
	followSets := c.followSets(startState)

	callStack.Add(startState.GetRuleIndex())

	currentSymbol := c.tokens.Get(tokenIndex)

	if tokenIndex >= c.tokens.Len()-1 { // At caret?
		if Contains(c.preferredRules, startState.GetRuleIndex()) {
			// No need to go deeper when collecting entries and we reach a rule that we want to collect anyway.
			c.translateToRuleIndex(callStack)
		} else {
			// Convert all follow sets to either single symbols or their associated preferred rule and add
			// the result to our candidates list.
			for _, set := range followSets.sets.ToArray() {
				fullPath := mutablelist.New(func(a, b int) bool { return a == b })
				fullPath.AddList(callStack)
				fullPath.AddList(set.path)

				if !c.translateToRuleIndex(fullPath) {
					for _, symbol := range set.intervals.GetIntervals() {
						if !Contains(c.ignoredTokens, symbol.Start) {
							if c.showDebugOutput {
								fmt.Printf("=====> collected: %s\n", c.vocabulary.GetDisplayName(symbol.Start))
							}
							if c.candidates.tokens.Contains(symbol.Start) {
								stack := list.New(func(a, b int) bool { return a == b })
								for _, cs := range callStack.ToArray() {
									stack.Add(cs)
								}
								for _, sp := range set.path.ToArray() {
									stack.Add(sp)
								}

								c.candidates.RecordToken(
									symbol.Start,
									set.following,
									stack,
								)
							} else {
								// More than one following list for the same symbol.
								if c.candidates.tokens.Get(symbol.Start) != set.following {
									stack := list.New(func(a, b int) bool { return a == b })
									for _, cs := range callStack.ToArray() {
										stack.Add(cs)
									}
									for _, sp := range set.path.ToArray() {
										stack.Add(sp)
									}

									c.candidates.RecordToken(
										symbol.Start,
										NewTokenList(),
										stack,
									)
								}
							}
						}
					}
				}
			}
		}

		callStack.Pop()

		//println("POPPING FROM CALLSTACK A $callStack")
		return result

	} else {
		// Process the rule if we either could pass it without consuming anything (epsilon transition)
		// or if the current input symbol will be matched somewhere after this entry point.
		// Otherwise stop here.
		if followSets.combined != nil {
			if !followSets.combined.Contains(antlr.TokenEpsilon) && !followSets.combined.Contains(currentSymbol) {
				callStack.Pop()
				//println("POPPING FROM CALLSTACK B $callStack")
				return result
			}
		}
	}

	// The current state execution pipeline contains all yet-to-be-processed ATN states in this rule.
	// For each such state we store the token index + a list of rules that lead to it.
	statePipeline := list.New(func(a, b *PipelineEntry) bool { return a.state.Equals(b.state) })
	var currentEntry *PipelineEntry

	// Bootstrap the pipeline.
	statePipeline.Add(NewPipelineEntry(
		startState,
		tokenIndex,
		callStack,
		nil,
	))

	processed := *list.New(func(a, b *PipelineEntry) bool { return a.state.Equals(b.state) })
pipelineLoop:
	for statePipeline.Len() != 0 {
		if statePipeline.Len() > 1000 {
			panic("State pipeline way too big")
		}
		currentEntry = statePipeline.Pop()

		if currentEntry == nil {
			continue
		}

		if processed.Contains(currentEntry) {
			continue
		} else {
			processed.Add(currentEntry)
		}

		c.statesProcessed++

		currentSymbol = c.tokens.Get(currentEntry.tokenIndex)

		atCaret := false
		if currentEntry != nil && currentEntry.tokenIndex >= c.tokens.Len()-1 {
			atCaret = true
		}

		if c.showDebugOutput {
			c.printDescription(
				indentation,
				currentEntry.state,
				c.generateBaseDescription(currentEntry.state),
				currentEntry.tokenIndex,
			)
			if c.showRuleStack {
				c.printRuleState(currentEntry.callStack.ToArray())
			}
		}

		//println("  Processing Pipeline Entry ${currentEntry.state} ${currentEntry.state.stateType} ${currentEntry.state.describe(ruleNames)} ${currentEntry.transitions}")

		myPathDescription := pathDescription

		if currentEntry.state.GetStateType() == antlr.ATNStateRuleStart {
			myPathDescription = pathDescription.followState(currentEntry.state)
		}

		switch currentEntry.state.GetStateType() {
		case antlr.ATNStateRuleStart: // Happens only for the first state in this rule, not subrules.
			indentation += "  "
		case antlr.ATNStateRuleStop:
			result.Add(currentEntry.tokenIndex)
			continue pipelineLoop
		}

		transitions := currentEntry.state.GetTransitions()
	myFor:
		for _, transition := range transitions {
			//println("    transition ${transition.serializationType}")
			serializationType := GetUnexportedField[antlr.Transition, int](transition, "serializationType")
			target := GetUnexportedField[antlr.Transition, antlr.ATNState](transition, "target")

			switch serializationType {
			case antlr.TransitionRULE:
				endStatus := c.processRule(
					target,
					currentEntry.tokenIndex,
					currentEntry.callStack,
					indentation,
					myPathDescription.follow(transition),
				)

				followState := GetUnexportedField[*antlr.RuleTransition, antlr.ATNState](
					transition.(*antlr.RuleTransition),
					"followState",
				)

				trList := list.New[antlr.Transition](func(a, b antlr.Transition) bool { return a == b })
				trList.AddList(currentEntry.transitions)
				trList.Add(transition)

				for _, position := range endStatus.ToArray() {
					statePipeline.Add(
						NewPipelineEntry(
							followState,
							position,
							currentEntry.callStack,
							trList,
						),
					)
				}

			case antlr.TransitionPREDICATE:
				if c.checkPredicate(transition.(*antlr.PredicateTransition)) {
					currentEntry.transitions.Add(transition)

					statePipeline.Add(
						NewPipelineEntry(
							target,
							currentEntry.tokenIndex,
							currentEntry.callStack,
							currentEntry.transitions,
						),
					)
				}

			case antlr.TransitionWILDCARD:
				if atCaret {
					if !c.translateToRuleIndex(currentEntry.callStack) {
						maxTokentType := GetUnexportedField[*antlr.ATN, int](c.atn, "maxTokenType")
						list := IntervalSet.FromStartStop(antlr.TokenMinUserTokenType, maxTokentType).ToList()

						for _, item := range list {
							if !Contains(c.ignoredTokens, item) {
								if currentEntry != nil {
									c.candidates.RecordToken(
										item,
										NewTokenList(),
										currentEntry.UpdatedCallStack(),
									)
								} else {
									c.candidates.RecordToken(
										item,
										NewTokenList(),
										nil,
									)
								}
							}
						}
					}
				} else {

					trList := list.New[antlr.Transition](func(a, b antlr.Transition) bool { return a == b })
					trList.AddList(currentEntry.transitions)
					trList.Add(transition)

					statePipeline.Add(
						NewPipelineEntry(
							target,
							currentEntry.tokenIndex+1,
							currentEntry.callStack,
							trList,
						),
					)
				}

			default:
				isEpsilon := GetUnexportedField[antlr.Transition, bool](transition, "isEpsilon")
				if isEpsilon {
					// Jump over simple states with a single outgoing epsilon transition.

					trList := list.New(func(a, b antlr.Transition) bool { return a == b })
					trList.AddList(currentEntry.transitions)
					trList.Add(transition)

					statePipeline.Add(
						NewPipelineEntry(
							target,
							currentEntry.tokenIndex,
							currentEntry.callStack,
							trList,
						),
					)

					continue myFor
				}

				base := GetUnexportedField[antlr.Transition, *antlr.IntervalSet](transition, "intervalSet")
				set := IntervalSet.FromAntlr(base)
				if len(set.GetIntervals()) != 0 {
					serializationType := GetUnexportedField[antlr.Transition, int](transition, "serializationType")
					if serializationType == antlr.TransitionNOTSET {
						maxTokentType := GetUnexportedField[*antlr.ATN, int](c.atn, "maxTokentType")
						set = set.Complement(antlr.TokenMinUserTokenType, maxTokentType)
					}

					//println("    transition atCaret? $atCaret tokenIndex ${currentEntry!!.tokenIndex} tokens ${tokens}")
					if atCaret {
						if !c.translateToRuleIndex(currentEntry.callStack) {
							list := set.ToList()

							addFollowing := len(list) == 1
							for _, symbol := range list {
								if !Contains(c.ignoredTokens, symbol) {
									if c.showDebugOutput {
										fmt.Printf("=====> collected: %s\n", c.vocabulary.GetDisplayName(symbol))
									}

									if addFollowing {
										c.candidates.RecordToken(
											symbol,
											c.getFollowingTokens(transition),
											currentEntry.UpdatedCallStack(),
										)
									} else {
										c.candidates.RecordToken(
											symbol,
											NewTokenList(),
											currentEntry.UpdatedCallStack(),
										)
									}
								}
							}
						}
					} else {
						if set.Contains(currentSymbol) {
							if c.showDebugOutput {
								fmt.Printf("=====> consumed: %s\n", c.vocabulary.GetDisplayName(currentSymbol))
							}

							trList := list.New[antlr.Transition](func(a, b antlr.Transition) bool { return a == b })
							trList.AddList(currentEntry.transitions)
							trList.Add(transition)

							statePipeline.Add(
								NewPipelineEntry(
									target,
									currentEntry.tokenIndex+1,
									currentEntry.callStack,
									trList,
								),
							)
						}
					}
				}
			}
		}
	}

	callStack.Pop()

	// Cache the result, for later lookup to avoid duplicate walks.
	positionMap.Set(tokenIndex, result)

	return result
}

func (c *CodeCompletionCore) generateBaseDescription(state antlr.ATNState) string {
	stateValue := "Invalid"
	if state.GetStateNumber() != antlr.ATNStateInvalidStateNumber {
		stateValue = string(state.GetStateNumber())
	}

	return fmt.Sprintf(
		"[%s %s] in %s",
		stateValue, atnStateTypeMap[state.GetStateType()], c.ruleNames[state.GetRuleIndex()],
	)
}

func (c *CodeCompletionCore) printDescription(
	currentIndent string,
	state antlr.ATNState,
	baseDescription string,
	tokenIndex int,
) {
	output := currentIndent

	transitionDescription := ""
	if c.debugOutputWithTransitions {
		for _, transition := range state.GetTransitions() {
			labels := ""

			symbols := make([]int, 0)

			base := GetUnexportedField[antlr.Transition, *antlr.IntervalSet](transition, "intervalSet")
			if base != nil {
				set := IntervalSet.FromAntlr(base)
				symbols = set.ToList()
			}

			if len(symbols) > 2 {
				// Only print start and end symbols to avoid large lists in debug output.
				labels = c.vocabulary.GetDisplayName(symbols[0]) + ".." + c.vocabulary.GetDisplayName(symbols[len(symbols)-1])
			} else {
				for _, symbol := range symbols {

					if labels != "" {
						labels += ", "
					}

					labels += c.vocabulary.GetDisplayName(symbol)
				}
			}

			if labels == "" {
				labels = "ε"
			}

			target := GetUnexportedField[antlr.Transition, antlr.ATNState](transition, "target")

			transitionDescription += fmt.Sprintf(
				"\n%s\t(%s) [%d] in %s",
				currentIndent, labels, target.GetStateType(), c.ruleNames[target.GetRuleIndex()],
			)
		}
	}

	if tokenIndex >= c.tokens.Len()-1 {
		output += fmt.Sprintf(
			"<<%d%d>> ",
			c.tokensProvider.TokensStartIndex(), tokenIndex,
		)
	} else {
		output += fmt.Sprintf(
			"<%d%d> ",
			c.tokensProvider.TokensStartIndex(), tokenIndex,
		)
	}

	fmt.Println(output + "Current state: " + baseDescription + transitionDescription)
}

func (c *CodeCompletionCore) printRuleState(stack []int) {
	if len(stack) == 0 {
		fmt.Println("<empty stack>")
		return
	}

	for _, rule := range stack {
		fmt.Printf("(R) %s\n", c.ruleNames[rule])
	}
}

// // The main class for doing the collection process.
// class CodeCompletionCore(val atn: ATN, val vocabulary: Vocabulary, val ruleNames: Array<String>,
// 	val languageName : String,
// 	val predicatesEvaluator : Recognizer<*, *>? = null) {

// }
