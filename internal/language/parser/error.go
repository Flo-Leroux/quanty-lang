package parser

import (
	"errors"
	"fmt"

	"github.com/Flo-Leroux/quanty-lang/internal/language/token"
)

// peekError -
func (p *Parser) peekError(t token.Type) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	err := errors.New(msg)
	p.errors = append(p.errors, err)
}

// peekError -
func (p *Parser) peekUnexpectedClosing(t token.Type) {
	msg := fmt.Sprintf("unexpected closing token %s", t)
	err := errors.New(msg)
	p.errors = append(p.errors, err)
}
