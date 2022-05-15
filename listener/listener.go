package listener

import (
	"fmt"
	"quanty/antlr/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type QuantyListener struct {
	*parser.BaseParserListener
}

func NewQuantyListener() *QuantyListener {
	return new(QuantyListener)
}

func (ql *QuantyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}
