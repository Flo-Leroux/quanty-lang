package parser

import (
	. "quanty/ast"
	"quanty/lexer"
)

func (p *parser) parseOptionalSelectionSet() SelectionSet {
	var selections []Selection
	p.some(lexer.BraceL, lexer.BraceR, func() {
		selections = append(selections, p.parseSelection())
	})

	return SelectionSet(selections)
}

func (p *parser) parseRequiredSelectionSet() SelectionSet {
	if p.peek().Kind != lexer.BraceL {
		p.error(p.peek(), "Expected %s, found %s", lexer.BraceL, p.peek().Kind.String())
		return nil
	}

	var selections []Selection
	p.some(lexer.BraceL, lexer.BraceR, func() {
		selections = append(selections, p.parseSelection())
	})

	return SelectionSet(selections)
}

func (p *parser) parseSelection() Selection {
	if p.peek().Kind == lexer.Spread {
		return p.parseFragment()
	}
	return p.parseField()
}
