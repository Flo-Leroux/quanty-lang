package parser

import (
	. "quanty/ast"
	"quanty/lexer"
)

func (p *parser) parseFragment() Selection {
	p.expect(lexer.Spread)

	if peek := p.peek(); peek.Kind == lexer.Name && peek.Value != "on" {
		return &FragmentSpread{
			Position:   p.peekPos(),
			Name:       p.parseFragmentName(),
			Directives: p.parseDirectives(false),
		}
	}

	var def InlineFragment
	def.Position = p.peekPos()
	if p.peek().Value == "on" {
		p.next() // "on"

		def.TypeCondition = p.parseName()
	}

	def.Directives = p.parseDirectives(false)
	def.SelectionSet = p.parseRequiredSelectionSet()
	return &def
}

func (p *parser) parseFragmentDefinition() *FragmentDefinition {
	var def FragmentDefinition
	def.Position = p.peekPos()
	p.expectKeyword("fragment")

	def.Name = p.parseFragmentName()
	def.VariableDefinition = p.parseVariableDefinitions()

	p.expectKeyword("on")

	def.TypeCondition = p.parseName()
	def.Directives = p.parseDirectives(false)
	def.SelectionSet = p.parseRequiredSelectionSet()
	return &def
}

func (p *parser) parseFragmentName() string {
	if p.peek().Value == "on" {
		p.unexpectedError()
		return ""
	}

	return p.parseName()
}
