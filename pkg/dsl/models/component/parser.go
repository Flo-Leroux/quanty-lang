package component

import (
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/models"
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/models/node"
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/parser"
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/token"
)

var Parser parser.RegisterFnStatement = func(p *parser.Parser) node.Statement {
	if !p.ExpectAndNext(token.IDENT) {
		return nil
	}

	cs := New(p.CurrentToken.Literal)

	p.WrapWith(
		parser.BraceWrapper,
		func() {
			switch p.CurrentToken.Type {
			case token.STRING:
				cs.WithSelections(p.ParseStatement(models.String))
			case token.IDENT:
				cs.WithSelections(p.ParseStatement(models.Field))
			}
		},
	)

	return cs
}
