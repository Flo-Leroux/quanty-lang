package parser

import (
	. "quanty/ast"
	"quanty/lexer"
)

func (p *parser) parseDirectives(isConst bool) []*Directive {
	var directives []*Directive

	for p.peek().Kind == lexer.At {
		if p.err != nil {
			break
		}
		directives = append(directives, p.parseDirective(isConst))
	}
	return directives
}

func (p *parser) parseDirective(isConst bool) *Directive {
	p.expect(lexer.At)

	return &Directive{
		Position:  p.peekPos(),
		Name:      p.parseName(),
		Arguments: p.parseArguments(isConst),
	}
}
