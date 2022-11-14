package complete

import (
	"fmt"

	mutablemap "github.com/Flo-Leroux/quanty-lang/pkg/complete/mutable-map"
	mutableset "github.com/Flo-Leroux/quanty-lang/pkg/complete/mutable-set"
)

// All the candidates which have been found. Tokens and rules are separated (both use a numeric value).
// Token entries include a list of tokens that directly follow them (see also the "following" member in the FollowSetWithPath class).
type CandidatesCollection struct {

	// The map keys are the lexer tokens
	// The list consists of further token ids which directly follow the given token in the grammar (if any)
	tokens *mutablemap.MutableMap[TokenKind, TokenList]

	// The map keys are the rule indices
	// The list represents the call stack at which the given rule was found during evaluation
	// This allows to determine a context for rules that are used in different places
	rules         *mutablemap.MutableMap[RuleIndex, RuleList]
	tokensContext *mutablemap.MutableMap[TokenKind, ParserStack]
}

func NewCandidatesCollection() *CandidatesCollection {
	return &CandidatesCollection{
		tokens:        mutablemap.New[TokenKind, TokenList](),
		rules:         mutablemap.New[RuleIndex, RuleList](),
		tokensContext: mutablemap.New[TokenKind, ParserStack](),
	}
}

func (c *CandidatesCollection) RecordToken(tokenKind TokenKind, followers TokenList, parserStack ParserStack) {
	c.tokens.Set(tokenKind, followers)
	c.tokensContext.Set(tokenKind, parserStack)
}

func (c *CandidatesCollection) String() string {
	return fmt.Sprintf("CandidatesCollection(tokens=%s, rules=%s, tokensContext=%s)", c.tokens, c.rules, c.tokensContext)
}

func (c *CandidatesCollection) Clear() {
	c.rules.Clear()
	c.tokens.Clear()
	c.tokensContext.Clear()
}

func (c *CandidatesCollection) Tokens() *mutablemap.MutableMap[TokenKind, TokenList] {
	return c.tokens
}

func (c *CandidatesCollection) Rules() *mutablemap.MutableMap[RuleIndex, RuleList] {
	return c.rules
}

func (c *CandidatesCollection) Show(ruleNames []string, vocabulary Vocabulary) {

	if c.rules.Len() > 0 {
		fmt.Printf("\n\nCollected rules:\n")
		for key, value := range c.rules.ToMap() {
			path := ""
			for _, token := range value.ToArray() {
				path += ruleNames[token] + " "
			}
			fmt.Printf("  - %s, path: %s", ruleNames[key], path)
		}
	}

	if c.tokens.Len() > 0 {
		sortedTokens := mutableset.New(func(a, b string) bool { return a == b })
		for key, _ := range c.tokens.ToMap() {
			value := vocabulary.GetDisplayName(key)
			// for _, following := range value1.ToArray() {
			// 	value += " => " + vocabulary.GetDisplayName(following)
			// }
			sortedTokens.Add(value)
		}

		fmt.Print("\n\nCollected tokens:\n")
		for _, symbol := range sortedTokens.ToArray() {
			fmt.Printf("  - %s\n", symbol)
		}
	}
	fmt.Printf("\n")
}
