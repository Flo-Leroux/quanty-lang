package parser

import (
	"github.com/Flo-Leroux/quanty-lang/internal/errors"
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/lexer"
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/models"
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/models/node"
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/token"
)

// Parser -
type Parser struct {
	l            *lexer.Lexer
	CurrentToken token.Token
	PeekToken    token.Token
	errors       []errors.Error

	nodeParsers      map[models.Name]RegisterFnNode
	statementParsers map[models.Name]RegisterFnStatement
	// errors         []errors.Error
	// prefixParseFns map[token.Type]prefixParseFn
	// infixParseFunc map[token.Type]infixParseFn
	stackWrapper []token.Type
}

// NewParser -
func New(str string) (p *Parser) {
	p = &Parser{
		l:      lexer.New(str),
		errors: make([]errors.Error, 0),

		nodeParsers:      make(map[models.Name]func(*Parser) node.Node),
		statementParsers: make(map[models.Name]func(*Parser) node.Statement),
	}

	// p.prefixParseFns = make(map[token.Type]prefixParseFn)
	// // p.registerPrefix(token.IDENT, p.parseIdentifier)
	// // p.registerPrefix(token.NOT, p.parsePrefixExpression)

	// p.infixParseFunc = make(map[token.Type]infixParseFn)
	// // p.registerInfix(token.AND, p.parseInfixExpression)
	// p.registerInfix(token.OR, p.parseInfixExpression)

	p.Next()
	p.Next()
	return
}

// next -
func (p *Parser) Next() {
	p.CurrentToken = p.PeekToken
	p.PeekToken = p.l.NextToken()
	p.checkStackWrappers()
}

// checkStackWrappers -
func (p *Parser) checkStackWrappers() {
	if Wrappers.isOpen(p.CurrentToken) {
		p.stackWrapper = append(p.stackWrapper, p.CurrentToken.Type)
	}

	if Wrappers.isClose(p.CurrentToken) {
		if len(p.stackWrapper) == 0 {
			p.PeekUnexpectedClosing(p.CurrentToken.Type)
		} else {
			last := p.stackWrapper[len(p.stackWrapper)-1:][0]

			expected := Wrappers.findOpenWithClose(p.CurrentToken.Type)
			if last != expected {
				p.PeekUnexpectedClosing(p.CurrentToken.Type)
			}

			p.stackWrapper = p.stackWrapper[:len(p.stackWrapper)-1]
		}
	}
}

// currentTokenIs -
func (p *Parser) CurrentTokenIs(t token.Type) bool {
	return p.CurrentToken.Type == t
}

// peekTokenIs -
func (p *Parser) PeekTokenIs(t token.Type) bool {
	return p.PeekToken.Type == t
}

// Error -
func (p *Parser) Error() error {
	if len(p.errors) == 0 {
		return nil
	}
	panic(p.errors[0])
}

// Errors -
func (p *Parser) Errors() []errors.Error {
	return p.errors
}

// Parse -
// func (p *Parser) Parse() *ast.Schema {
// 	schema := ast.NewSchema()

// 	// for !p.currentTokenIs(token.EOF) {
// 	// 	stmt := p.parseStatement()
// 	// 	if stmt != nil {
// 	// 		schema.WithStatements(stmt)
// 	// 	}
// 	// 	p.next()
// 	// }

// 	return schema
// }

// expectAndNext -
func (p *Parser) ExpectAndNext(t token.Type) bool {
	if p.PeekTokenIs(t) {
		p.Next()
		return true
	}
	p.PeekError(t)
	return false
}
