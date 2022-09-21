package parser

import (
	"github.com/Flo-Leroux/quanty-lang/internal/language/ast"
	"github.com/Flo-Leroux/quanty-lang/internal/language/token"
)

// parseStatement method based on defined token types
func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.COMPONENT:
		return p.parseComponentStatement()
	default:
		return nil
	}
}
