package lexer

import (
	"testing"

	"quanty/token"
)

func TestNextToken(t *testing.T) {
	input := `
		component Main {
			div {
				p {
					"Hello World"
				}
			}
		}
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.COMPONENT, "component"},
		{token.IDENT, "Main"},
		{token.LBRACE, "{"},
		{token.IDENT, "div"},
		{token.LBRACE, "{"},
		{token.IDENT, "p"},
		{token.LBRACE, "{"},
		{token.DQUOTE, "\""},
		{token.IDENT, "Hello"},
		{token.IDENT, "World"},
		{token.DQUOTE, "\""},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
