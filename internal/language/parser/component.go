package parser

import (
	"github.com/Flo-Leroux/quanty-lang/internal/language/ast"
	"github.com/Flo-Leroux/quanty-lang/internal/language/token"
)

// parseComponentStatement returns a COMPONENT Statement AST Node
func (p *Parser) parseComponentStatement() *ast.ComponentStatement {
	if !p.expectAndNext(token.IDENT) {
		return nil
	}

	cs := ast.NewComponentStatement(p.currentToken.Literal)

	p.wrapWith(
		BraceWrapper,
		func() {
			switch p.currentToken.Type {
			case token.STRING:
				cs.WithSelections(p.parseStringValue())
			case token.IDENT:
				cs.WithSelections(p.parseField())
			}
		},
	)

	return cs
}
