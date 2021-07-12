package parser

import (
	"quanty/lexer"
)

func (p *parser) parseName() string {
	token := p.expect(lexer.Name)

	return token.Value
}
