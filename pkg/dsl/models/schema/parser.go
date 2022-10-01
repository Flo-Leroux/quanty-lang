package schema

import (
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/models"
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/models/node"
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/parser"
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/token"
)

var Parser parser.RegisterFnNode = func(p *parser.Parser) node.Node {
	schema := New()

	for !p.CurrentTokenIs(token.EOF) {
		stmt := parseStatement(p)
		if stmt != nil {
			schema.WithStatements(stmt)
		}
		p.Next()
	}

	return schema
}

func parseStatement(p *parser.Parser) node.Statement {
	switch p.CurrentToken.Type {
	case token.COMPONENT:
		return p.ParseStatement(models.Componnt)
	default:
		return nil
	}
}
