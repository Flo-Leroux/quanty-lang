package complete

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type ByParserClassTokenProvider struct {
	lexerFn func(input antlr.CharStream) antlr.Lexer
	parser  antlr.Parser
	code    string
}

func NewByParserClassTokenProvider(
	lexerFn func(input antlr.CharStream) antlr.Lexer,
	parser antlr.Parser,
	code string,
) *ByParserClassTokenProvider {
	return &ByParserClassTokenProvider{
		lexerFn: lexerFn,
		parser:  parser,
		code:    code,
	}
}

func (p *ByParserClassTokenProvider) TokensStartIndex() int {
	return 0
}

func (p *ByParserClassTokenProvider) StartRuleIndex() int {
	return 0
}

func (p *ByParserClassTokenProvider) Tokens() TokenList {
	input := antlr.NewInputStream(p.code)
	lexer := p.lexerFn(input)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	tokens := NewTokenList()
	offset := 1

	for {
		offset++
		token := stream.LT(offset)
		tokens.Add(token.GetTokenType())

		if token.GetTokenType() == antlr.TokenEOF {
			break
		}
	}

	return tokens
}
