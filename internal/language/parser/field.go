package parser

import (
	"github.com/Flo-Leroux/quanty-lang/internal/language/ast"
	"github.com/Flo-Leroux/quanty-lang/internal/language/token"
)

func (p *Parser) parseField() *ast.Field {
	f := &ast.Field{
		Name:       p.currentToken,
		Selections: []*ast.Field{},
	}

	if p.peekTokenIs(token.LBRACE) {
		p.wrapWith(
			BraceWrapper,
			func() {
				switch p.currentToken.Type {
				case token.IDENT:
					f.Selections = append(f.Selections, p.parseField())
				}
			},
		)
	}

	return f
}
