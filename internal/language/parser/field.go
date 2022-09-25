package parser

import (
	"github.com/Flo-Leroux/quanty-lang/internal/language/ast"
	"github.com/Flo-Leroux/quanty-lang/internal/language/token"
)

func (p *Parser) parseField() *ast.Field {
	f := ast.NewField(p.currentToken.Literal)

	if p.peekTokenIs(token.LBRACE) {
		p.wrapWith(
			BraceWrapper,
			func() {
				switch p.currentToken.Type {
				case token.STRING:
					f.WithSelections(p.parseStringValue())
				case token.IDENT:
					f.WithSelections(p.parseField())
				}
			},
		)
	}

	return f
}

func (p *Parser) parseStringValue() *ast.StringValue {
	return ast.NewStringValue(p.currentToken.Literal)
}
