package field

import (
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/models"
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/models/node"
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/parser"
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/token"
)

var Parser parser.RegisterFnStatement = func(p *parser.Parser) node.Statement {
	f := New(p.CurrentToken.Literal)

	if p.PeekTokenIs(token.LBRACE) {
		p.WrapWith(
			parser.BraceWrapper,
			func() {
				switch p.CurrentToken.Type {
				case token.STRING:
					f.WithSelections(p.ParseStatement(models.String))
				case token.IDENT:
					f.WithSelections(p.ParseStatement(models.Field))
				}
			},
		)
	}

	return f
}
