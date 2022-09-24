package parser

import (
	"github.com/Flo-Leroux/quanty-lang/internal/language/ast"
	"github.com/Flo-Leroux/quanty-lang/internal/language/token"
)

func (p *Parser) parseField() *ast.Field {
	f := &ast.Field{
		Name:       p.currentToken,
		Selections: ast.SelectionList{},
	}

	if p.peekTokenIs(token.LBRACE) {
		p.wrapWith(
			BraceWrapper,
			func() {
				switch p.currentToken.Type {
				case token.STRING:
					f.Selections = append(f.Selections, p.parseStringValue())
				case token.IDENT:
					f.Selections = append(f.Selections, p.parseField())
				}
			},
		)
	}

	return f
}

func (p *Parser) parseStringValue() *ast.StringValue {
	return &ast.StringValue{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}
}
