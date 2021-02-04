package main

import (
	"fmt"
	"io"
)

// Root represents a component.
type Root struct {
	Components []*Component
}

// Parser represents a parser.
type Parser struct {
	s   *Scanner
	buf struct {
		tok Token  // Last read token
		lit string // Last read literal
		n   int    // Buffer size (max=1)
	}
}

// NewParser returns a nex instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{
		s: NewScanner(r),
	}
}

// Scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() {
	p.buf.n = 1
}

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()

	if tok == WS {
		tok, lit = p.scan()
	}

	return
}

// Parse parses a SQL SELECT statement.
func (p *Parser) Parse() (*Root, error) {
	root := &Root{}

	tok, lit := p.scanIgnoreWhitespace()

	switch tok {
	case COMPONENT:
		p.unscan()
		component, err := p.scanComponent()
		if err != nil {
			return nil, err
		}
		root.Components = append(root.Components, component)
		break

	default:
		return nil, fmt.Errorf("Found %q, expected component", lit)
	}

	return root, nil
}
