package complete

import (
	list "github.com/Flo-Leroux/quanty-lang/pkg/complete/list"
	mutablelist "github.com/Flo-Leroux/quanty-lang/pkg/complete/mutable-list"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type PipelineEntry struct {
	state       antlr.ATNState
	tokenIndex  int
	callStack   *mutablelist.MutableList[RuleKind]
	transitions *list.List[antlr.Transition]
}

func NewPipelineEntry(
	state antlr.ATNState,
	tokenIndex int,
	callStack *mutablelist.MutableList[RuleKind],
	transitions *list.List[antlr.Transition],
) *PipelineEntry {
	p := &PipelineEntry{
		state:       state,
		tokenIndex:  tokenIndex,
		callStack:   callStack,
		transitions: transitions,
	}

	if transitions == nil {
		p.transitions = list.New[antlr.Transition](func(a, b antlr.Transition) bool { return a == b })
	}

	return p
}

func (p *PipelineEntry) UpdatedCallStack() ParserStack {

	parserStack := NewParserStack()

	parserStack.AddMany(p.callStack.ToArray())

	transitions := p.transitions.Filter(func(t antlr.Transition) bool {
		target := GetUnexportedField[antlr.Transition, antlr.ATNState](t, "target")
		return target.GetStateType() == antlr.ATNStateRuleStart
	})

	for _, transition := range transitions.ToArray() {
		target := GetUnexportedField[antlr.Transition, antlr.ATNState](transition, "target")

		parserStack.Add(target.GetRuleIndex())
	}

	return parserStack
}
