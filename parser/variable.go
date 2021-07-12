package parser

import (
	. "quanty/ast"
	"quanty/lexer"
)

func (p *parser) parseVariableDefinitions() VariableDefinitionList {
	var defs []*VariableDefinition
	p.many(lexer.ParenL, lexer.ParenR, func() {
		defs = append(defs, p.parseVariableDefinition())
	})

	return defs
}

func (p *parser) parseVariableDefinition() *VariableDefinition {
	var def VariableDefinition
	def.Position = p.peekPos()
	def.Variable = p.parseVariable()

	p.expect(lexer.Colon)

	def.Type = p.parseTypeReference()

	if p.skip(lexer.Equals) {
		def.DefaultValue = p.parseValueLiteral(true)
	}

	def.Directives = p.parseDirectives(false)

	return &def
}

func (p *parser) parseVariable() string {
	p.expect(lexer.Dollar)
	return p.parseName()
}
