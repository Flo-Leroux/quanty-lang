package parser

import (
	"fmt"

	"github.com/Flo-Leroux/quanty-lang/internal/errors"
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/token"
)

var parserError = errors.New(errors.LanguageError).SetSubKind("Parser")

// peekError -
func (p *Parser) PeekError(t token.Type) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.PeekToken.Type)
	p.errors = append(
		p.errors,
		parserError.
			SetMessage(msg).
			AddParam("expected", t).
			AddParam("got", p.PeekToken.Type),
	)
}

// peekError -
func (p *Parser) PeekUnexpectedClosing(t token.Type) {
	msg := fmt.Sprintf("unexpected closing token %s", t)
	p.errors = append(
		p.errors,
		parserError.
			SetMessage(msg).
			AddParam("unexpected", t),
	)
}
