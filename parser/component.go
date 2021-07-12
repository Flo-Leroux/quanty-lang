package parser

import (
	. "quanty/ast"
)

func (p *parser) parseComponent() *ComponentDefinition {
	p.expectKeyword("component")

	var def ComponentDefinition
	def.Kind = Component
	def.Position = p.peekPos()
	def.Name = p.parseName()

	def.VariableDefinitions = p.parseVariableDefinitions()
	def.Directives = p.parseDirectives(false)
	def.SelectionSet = p.parseRequiredSelectionSet()

	return &def
}
