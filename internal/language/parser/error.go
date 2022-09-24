package parser

import (
	"fmt"

	"github.com/Flo-Leroux/quanty-lang/internal/errors"
	"github.com/Flo-Leroux/quanty-lang/internal/language/token"
)

var parserError = errors.New(errors.LanguageError).SetSubKind("Parser")

// peekError -
func (p *Parser) peekError(t token.Type) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(
		p.errors,
		parserError.
			SetMessage(msg).
			AddParam("expected", t).
			AddParam("got", p.peekToken.Type),
	)
}

// peekError -
func (p *Parser) peekUnexpectedClosing(t token.Type) {
	msg := fmt.Sprintf("unexpected closing token %s", t)
	p.errors = append(
		p.errors,
		parserError.
			SetMessage(msg).
			AddParam("unexpected", t),
	)
}
