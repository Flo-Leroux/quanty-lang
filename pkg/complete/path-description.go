package complete

import (
	"github.com/Flo-Leroux/quanty-lang/pkg/complete/list"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type PathDescription struct {
	initialState antlr.ATNState
	transitions  *list.List[antlr.ATNState]
}

func NewPathDescription(initialState antlr.ATNState, transitions *list.List[antlr.ATNState]) *PathDescription {
	p := &PathDescription{
		initialState: initialState,
		transitions:  list.New[antlr.ATNState](func(a, b antlr.ATNState) bool { return a.Equals(b) }),
	}

	if transitions != nil {
		p.transitions = transitions
	}

	return p
}

func (p *PathDescription) follow(transition antlr.Transition) *PathDescription {
	nl := list.New[antlr.ATNState](func(a, b antlr.ATNState) bool { return a.Equals(b) })
	nl.AddList(p.transitions)

	target := GetUnexportedField[antlr.Transition, antlr.ATNState](transition, "target")
	nl.Add(target)

	return NewPathDescription(p.initialState, nl)
}

func (p *PathDescription) followState(state antlr.ATNState) *PathDescription {
	nl := list.New[antlr.ATNState](func(a, b antlr.ATNState) bool { return a.Equals(b) })
	nl.AddList(p.transitions)

	nl.Add(state)

	return NewPathDescription(p.initialState, nl)
}
