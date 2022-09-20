package parser

import (
	"errors"
	"fmt"

	"github.com/Flo-Leroux/quanty-lang/internal/language/ast"
	"github.com/Flo-Leroux/quanty-lang/internal/language/lexer"
	"github.com/Flo-Leroux/quanty-lang/internal/language/token"
)

const (
	_ int = iota

	LOWEST
	LOGIC
	PREFIX // not IDENT
)

// var precedences = map[token.Type]int{
// 	// token.AND: LOGIC,
// 	// token.OR:  LOGIC,
// }

// Parser -
type Parser struct {
	l              *lexer.Lexer
	currentToken   token.Token
	peekToken      token.Token
	errors         []error
	prefixParseFns map[token.Type]prefixParseFn
	infixParseFunc map[token.Type]infixParseFn
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

// NewParser -
func NewParser(str string) (p *Parser) {
	p = &Parser{
		l:      lexer.NewLexer(str),
		errors: make([]error, 0),
	}

	p.prefixParseFns = make(map[token.Type]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	// p.registerPrefix(token.NOT, p.parsePrefixExpression)

	p.infixParseFunc = make(map[token.Type]infixParseFn)
	// p.registerInfix(token.AND, p.parseInfixExpression)
	// p.registerInfix(token.OR, p.parseInfixExpression)

	p.next()
	p.next()
	return
}

// next -
func (p *Parser) next() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// currentTokenIs -
func (p *Parser) currentTokenIs(t token.Type) bool {
	return p.currentToken.Type == t
}

// peekTokenIs -
func (p *Parser) peekTokenIs(t token.Type) bool {
	return p.peekToken.Type == t
}

// Error -
func (p *Parser) Error() error {
	if len(p.errors) == 0 {
		return nil
	}
	panic(p.errors[0])
	// return error.New(errors.Validation).SetParams(map[string]interface{}{
	// 	"schema": strings.Join(p.errors, ","),
	// })
}

// Errors -
func (p *Parser) Errors() []error {
	return p.errors
}

// Parse -
func (p *Parser) Parse() *ast.Schema {
	schema := &ast.Schema{}
	schema.Statements = []ast.Statement{}

	for !p.currentTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			schema.Statements = append(schema.Statements, stmt)
		}
		p.next()
	}

	return schema
}

// parseStatement method based on defined token types
func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.COMPONENT:
		return p.parseComponentStatement()
	default:
		return nil
	}
}

func (p *Parser) parseWrapByTokens(enter, exit token.Type, fn func()) {
	if !p.expectAndNext(enter) {
		return
	}

	for !p.currentTokenIs(exit) {
		fn()
		p.next()

		if p.currentTokenIs(token.EOF) {
			p.peekError(exit)
			return
		}
	}
}

// parseComponentStatement returns a COMPONENT Statement AST Node
func (p *Parser) parseComponentStatement() *ast.ComponentStatement {
	stmt := &ast.ComponentStatement{Token: p.currentToken, Fields: []*ast.Field{}}
	if !p.expectAndNext(token.IDENT) {
		return nil
	}
	stmt.Name = p.currentToken

	// if !p.expectAndNext(token.LBRACE) {
	// 	return nil
	// }

	// for !p.currentTokenIs(token.RBRACE) {
	// 	switch p.currentToken.Type {
	// 	case token.IDENT:
	// 		stmt.Fields = append(stmt.Fields, p.parseField())
	// 	}
	// 	p.next()

	// 	if p.currentTokenIs(token.EOF) {
	// 		p.peekError(token.RBRACE)
	// 		return nil
	// 	}
	// }

	p.parseWrapByTokens(
		token.LBRACE, token.RBRACE,
		func() {
			switch p.currentToken.Type {
			case token.IDENT:
				stmt.Fields = append(stmt.Fields, p.parseField())
			}
		},
	)

	// if p.peekTokenIs(token.OPTION) {
	// 	p.next()
	// 	stmt.Option = p.currentToken
	// }

	return stmt
}

func (p *Parser) parseField() *ast.Field {
	f := &ast.Field{
		Name:       p.currentToken,
		Selections: []*ast.Field{},
	}

	if p.peekTokenIs(token.LBRACE) {
		p.parseWrapByTokens(
			token.LBRACE, token.RBRACE,
			func() {
				switch p.currentToken.Type {
				case token.IDENT:
					f.Selections = append(f.Selections, p.parseField())
				}
			},
		)
	}

	return f
}

// // parseRelationStatement -
// func (p *Parser) parseRelationStatement() *ast.RelationStatement {
// 	stmt := &ast.RelationStatement{Token: p.currentToken}
// 	if !p.expectAndNext(token.IDENT) {
// 		return nil
// 	}
// 	stmt.Name = p.currentToken

// 	//&& !p.currentTokenIs(token.RELATION)
// 	for p.peekTokenIs(token.SIGN) && !p.peekTokenIs(token.OPTION) {
// 		stmt.RelationTypes = append(stmt.RelationTypes, p.parseRelationTypeStatement())
// 	}

// 	if p.peekTokenIs(token.OPTION) {
// 		p.next()
// 		stmt.Option = p.currentToken
// 	}

// 	return stmt
// }

// // parseRelationTypeStatement -
// func (p *Parser) parseRelationTypeStatement() *ast.RelationTypeStatement {
// 	if !p.expectAndNext(token.SIGN) {
// 		return nil
// 	}
// 	stmt := &ast.RelationTypeStatement{Sign: p.currentToken}
// 	if !p.expectAndNext(token.IDENT) {
// 		return nil
// 	}
// 	stmt.Token = p.currentToken
// 	return stmt
// }

// // parseActionStatement -
// func (p *Parser) parseActionStatement() ast.Statement {
// 	stmt := &ast.ActionStatement{Token: p.currentToken}

// 	if !p.expectAndNext(token.IDENT) {
// 		return nil
// 	}

// 	stmt.Name = p.currentToken

// 	if !p.expectAndNext(token.ASSIGN) {
// 		return nil
// 	}

// 	p.next()

// 	stmt.ExpressionStatement = p.parseExpressionStatement()

// 	return stmt
// }

// // parseExpressionStatement -
// func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
// 	stmt := &ast.ExpressionStatement{}
// 	stmt.Expression = p.parseExpression(LOWEST)

// 	if p.currentTokenIs(token.NEWLINE) {
// 		p.next()
// 	}

// 	return stmt
// }

// expectAndNext -
func (p *Parser) expectAndNext(t token.Type) bool {
	if p.peekTokenIs(t) {
		p.next()
		return true
	}
	p.peekError(t)
	return false
}

// // parseExpression -
// func (p *Parser) parseExpression(precedence int) ast.Expression {
// 	if p.currentTokenIs(token.LPAREN) {
// 		p.next()
// 		return p.parseInnerParen()
// 	}

// 	prefix := p.prefixParseFns[p.currentToken.Type]
// 	if prefix == nil {
// 		p.noPrefixParseFnError(p.currentToken.Type)
// 		return nil
// 	}
// 	left := prefix()

// 	for !p.peekTokenIs(token.NEWLINE) && precedence < p.peekPrecedence() {
// 		infix := p.infixParseFunc[p.peekToken.Type]
// 		if infix == nil {
// 			return left
// 		}
// 		p.next()
// 		left = infix(left)
// 	}

// 	return left
// }

// // parseInnerParen -
// func (p *Parser) parseInnerParen() ast.Expression {
// 	if p.currentTokenIs(token.LPAREN) {
// 		return p.parseExpression(LOWEST)
// 	}

// 	prefix := p.prefixParseFns[p.currentToken.Type]
// 	if prefix == nil {
// 		p.noPrefixParseFnError(p.currentToken.Type)
// 		return nil
// 	}
// 	left := prefix()

// 	for !p.currentTokenIs(token.RPAREN) {
// 		infix := p.infixParseFunc[p.peekToken.Type]
// 		if infix == nil {
// 			return left
// 		}
// 		p.next()
// 		left = infix(left)
// 	}

// 	return left
// }

// // parsePrefixExpression -
// func (p *Parser) parsePrefixExpression() ast.Expression {
// 	expression := &ast.PrefixExpression{
// 		Token:    p.currentToken,
// 		Operator: p.currentToken.Literal,
// 	}
// 	p.next()
// 	expression.Value = p.currentToken.Literal
// 	return expression
// }

// // parseInfixExpression
// func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
// 	expression := &ast.InfixExpression{
// 		Token:    p.currentToken, // and, or
// 		Left:     left,
// 		Operator: p.currentToken.Literal,
// 	}
// 	precedence := p.currentPrecedence()
// 	p.next()
// 	expression.Right = p.parseExpression(precedence)
// 	return expression
// }

// // peekPrecedence -
// func (p *Parser) peekPrecedence() int {
// 	if p, ok := precedences[p.peekToken.Type]; ok {
// 		return p
// 	}
// 	return LOWEST
// }

// // peekPrecedence -
// func (p *Parser) currentPrecedence() int {
// 	if p, ok := precedences[p.currentToken.Type]; ok {
// 		return p
// 	}
// 	return LOWEST
// }

// parseIdentifier
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
}

// registerPrefix
func (p *Parser) registerPrefix(tokenType token.Type, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

// // registerInfix
// func (p *Parser) registerInfix(tokenType token.Type, fn infixParseFn) {
// 	p.infixParseFunc[tokenType] = fn
// }

// // noPrefixParseFnError -
// func (p *Parser) noPrefixParseFnError(t token.Type) {
// 	msg := fmt.Sprintf("no prefix parse function for %s found", t)
// 	p.errors = append(p.errors, msg)
// }

// peekError -
func (p *Parser) peekError(t token.Type) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	err := errors.New(msg)
	p.errors = append(p.errors, err)
}
