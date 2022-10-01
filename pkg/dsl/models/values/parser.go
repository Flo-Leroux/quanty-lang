package values

import (
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/models/node"
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/parser"
)

var ParserString parser.RegisterFnStatement = func(p *parser.Parser) node.Statement {
	return NewString(p.CurrentToken.Literal)
}
