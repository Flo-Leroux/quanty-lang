package parser

import (
	. "quanty/ast"
	"quanty/lexer"
)

func (p *parser) parseField() *Field {
	var field Field
	field.Position = p.peekPos()
	field.IsTextContent = false

	if p.peekToken.Kind == lexer.String {
		field.Name = p.next().Value
		field.IsTextContent = true
	} else {
		field.Alias = p.parseName()

		if p.skip(lexer.Colon) {
			field.Name = p.parseName()
		} else {
			field.Name = field.Alias
			field.Alias = ""
		}

		field.Arguments = p.parseArguments(false)
		field.Directives = p.parseDirectives(false)
	}

	if p.peek().Kind == lexer.BraceL {
		field.SelectionSet = p.parseOptionalSelectionSet()
	}

	return &field
}
