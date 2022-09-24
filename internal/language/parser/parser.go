package parser

import (
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
	stackWrapper   []token.Type
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
	// // p.registerPrefix(token.NOT, p.parsePrefixExpression)

	p.infixParseFunc = make(map[token.Type]infixParseFn)
	// // p.registerInfix(token.AND, p.parseInfixExpression)
	// p.registerInfix(token.OR, p.parseInfixExpression)

	p.next()
	p.next()
	return
}

// next -
func (p *Parser) next() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
	p.checkStackWrappers()
}

// checkStackWrappers -
func (p *Parser) checkStackWrappers() {
	if wrappers.isOpen(p.currentToken) {
		p.stackWrapper = append(p.stackWrapper, p.currentToken.Type)
	}

	if wrappers.isClose(p.currentToken) {
		if len(p.stackWrapper) == 0 {
			p.peekUnexpectedClosing(p.currentToken.Type)
		} else {
			last := p.stackWrapper[len(p.stackWrapper)-1:][0]

			expected := wrappers.findOpenWithClose(p.currentToken.Type)
			if last != expected {
				p.peekUnexpectedClosing(p.currentToken.Type)
			}

			p.stackWrapper = p.stackWrapper[:len(p.stackWrapper)-1]
		}
	}
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
	// // return error.New(errors.Validation).SetParams(map[string]interface{}{
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

// expectAndNext -
func (p *Parser) expectAndNext(t token.Type) bool {
	if p.peekTokenIs(t) {
		p.next()
		return true
	}
	p.peekError(t)
	return false
}

// parseIdentifier
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
}

// registerPrefix
func (p *Parser) registerPrefix(tokenType token.Type, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}
