package parser

import (
	"fmt"
	"quanty/ast"
	"quanty/lexer"
	"quanty/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token

	errors []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) ParseDocument() *ast.Document {
	document := &ast.Document{}
	document.Definitions = []ast.Definition{}

	for p.curToken.Type != token.EOF {
		def := p.parseDefinition()
		if def != nil {
			document.Definitions = append(document.Definitions, def)
		}
		p.nextToken()
	}

	return document
}

func (p *Parser) parseDefinition() ast.Definition {
	switch p.curToken.Type {
	case token.COMPONENT:
		return p.parseComponentDefinition()
	default:
		return nil
	}
}

func (p *Parser) parseComponentDefinition() *ast.ComponentDefinition {
	comp := &ast.ComponentDefinition{Kind: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	comp.Name = &ast.Identifier{Kind: p.curToken, Value: p.curToken.Literal}

	if !p.peekTokenIs(token.LBRACE) {
		return nil
	}

	comp.SelectionSet = p.parseSelectionSet()

	// if !p.peekTokenIs(token.RBRACE) {
	// return nil
	// }

	return comp
}

func (p *Parser) parseSelectionSet() *ast.SelectionSet {
	set := &ast.SelectionSet{Kind: p.curToken}
	set.Selection = []ast.Definition{}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	for !p.peekTokenIs(token.RBRACE) {

		p.nextToken()
		switch p.curToken.Type {
		case token.IDENT:
			tag := &ast.TagDefinition{
				Kind: token.Token{Type: token.TAG, Literal: "TAG"},
				Name: &ast.Identifier{
					Kind:  p.curToken,
					Value: p.curToken.Literal,
				},
			}

			if p.peekTokenIs(token.LBRACE) {
				tag.SelectionSet = p.parseSelectionSet()
			}

			set.Selection = append(set.Selection, tag)
		case token.STRING:
			set.Selection = append(set.Selection, &ast.StringLiteral{
				Kind:  p.curToken,
				Value: p.curToken.Literal,
			})
		default:
			return nil
		}

		// if p.peekTokenIs(token.STRING) {
		// }
	}

	if !p.expectPeek(token.RBRACE) {
		return nil
	}

	return set
}
