package parser

import (
	. "quanty/ast"
	"quanty/lexer"
)

func (p *parser) parseTypeReference() *Type {
	var typ Type

	if p.skip(lexer.BracketL) {
		typ.Position = p.peekPos()
		typ.Elem = p.parseTypeReference()
		p.expect(lexer.BracketR)
	} else {
		typ.Position = p.peekPos()
		typ.NamedType = p.parseName()
	}

	if p.skip(lexer.Bang) {
		typ.Position = p.peekPos()
		typ.NonNull = true
	}
	return &typ
}

func (p *parser) parseName() string {
	token := p.expect(lexer.Name)

	return token.Value
}
