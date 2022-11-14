package complete

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type ByStreamTokenProvider struct {
	tokenStream     antlr.TokenStream
	caretTokenIndex int
	context         antlr.ParserRuleContext
}

func NewByStreamTokenProvider(
	tokenStream antlr.TokenStream,
	caretTokenIndex int,
	context antlr.ParserRuleContext,
) *ByStreamTokenProvider {
	return &ByStreamTokenProvider{
		tokenStream:     tokenStream,
		caretTokenIndex: caretTokenIndex,
		context:         context,
	}
}

func (p *ByStreamTokenProvider) TokensStartIndex() int {
	if p.context != nil {
		return p.context.GetStart().GetTokenIndex()
	}

	return 0
}

func (p *ByStreamTokenProvider) StartRuleIndex() int {
	if p.context != nil {
		return p.context.GetRuleIndex()
	}

	return 0
}

func (p *ByStreamTokenProvider) Tokens() TokenList {
	tokenStartIndex := p.TokensStartIndex()

	p.tokenStream.Seek(tokenStartIndex)
	tokens := NewTokenList()
	offset := 0

	for {
		offset++
		token := p.tokenStream.LT(offset)

		tokens.Add(token.GetTokenType())

		if token.GetTokenIndex() >= p.caretTokenIndex || token.GetTokenType() == antlr.TokenEOF {
			break
		}
	}

	currentIndex := p.tokenStream.Index()
	if currentIndex == -1 {
		panic("CurrentIndex should be not -1")
	}

	p.tokenStream.Seek(currentIndex)
	return tokens

}
