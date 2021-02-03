package main

import (
	"fmt"
)

// Component represents a component.
type Component struct {
	Name       string
	Properties []*Property
}

type Property struct {
	Name  string
	Type  string
	Value string
}

// Parse parses a SQL SELECT statement.
func (p *Parser) scanComponent() (*Component, error) {
	comp := &Component{}

	tok, lit := p.scanIgnoreWhitespace()
	if tok != COMPONENT {
		return nil, fmt.Errorf("Found %q, expected component", lit)
	}

	tok, lit = p.scanIgnoreWhitespace()
	if tok != IDENT {
		return nil, fmt.Errorf("Found %q, expected component name", lit)
	}
	comp.Name = lit

	tok, lit = p.scanIgnoreWhitespace()
	if tok != OPEN_BRACKET {
		return nil, fmt.Errorf("Found %q, expected open bracket", lit)
	}

	for {
		// Read content
		tok, lit := p.scanIgnoreWhitespace()

		if tok == CLOSE_BRACKET {
			p.unscan()
			break
		}

		switch tok {
		case PROP:
			p.unscan()
			prop, err := p.scanPropety()
			if err != nil {
				return nil, err
			}
			comp.Properties = append(comp.Properties, prop)
			break

		default:
			return nil, fmt.Errorf("Found %q, expected field or close bracket", lit)
		}
	}

	tok, lit = p.scanIgnoreWhitespace()
	if tok != CLOSE_BRACKET {
		return nil, fmt.Errorf("Found %q, expected close bracket", lit)
	}

	return comp, nil
}

// Parse parses a SQL SELECT statement.
func (p *Parser) scanPropety() (*Property, error) {
	prop := &Property{}

	tok, lit := p.scanIgnoreWhitespace()
	if tok != PROP {
		return nil, fmt.Errorf("Found %q, expected prop", lit)
	}

	tok, lit = p.scanIgnoreWhitespace()
	if tok != IDENT {
		return nil, fmt.Errorf("Found %q, expected property name", lit)
	}
	prop.Name = lit

	tok, lit = p.scanIgnoreWhitespace()
	if tok != IDENT {
		return nil, fmt.Errorf("Found %q, expected property type", lit)
	}
	prop.Type = lit

	tok, lit = p.scanIgnoreWhitespace()
	if tok != END && tok != EQUAL {
		return nil, fmt.Errorf("Found %q, expected property assignation or ;", lit)
	}

	if tok == EQUAL {
		tok, lit = p.scanIgnoreWhitespace()

		if tok != IDENT {
			return nil, fmt.Errorf("Found %q, expected property value", lit)
		}

		prop.Value = lit

		tok, lit = p.scanIgnoreWhitespace()
	}

	if tok != END {
		return nil, fmt.Errorf("Found %q, expected ;", lit)
	}

	return prop, nil
}
