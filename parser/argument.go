package parser

import (
	. "quanty/ast"
	"quanty/lexer"
)

func (p *parser) parseArguments(isConst bool) ArgumentList {
	var arguments ArgumentList
	p.many(lexer.ParenL, lexer.ParenR, func() {
		arguments = append(arguments, p.parseArgument(isConst))
	})

	return arguments
}

func (p *parser) parseArgument(isConst bool) *Argument {
	arg := Argument{}
	arg.Position = p.peekPos()
	arg.Name = p.parseName()
	p.expect(lexer.Colon)

	arg.Value = p.parseValueLiteral(isConst)
	return &arg
}
