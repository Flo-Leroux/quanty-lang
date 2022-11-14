package complete

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type ByListTokenProvider struct {
	tokens  TokenList
	context antlr.ParserRuleContext
}

func NewByListTokenProvider(tokens []TokenKind, context antlr.ParserRuleContext) *ByListTokenProvider {
	p := &ByListTokenProvider{
		tokens:  NewTokenList(),
		context: context,
	}

	for _, tok := range tokens {
		p.tokens.Add(tok)
	}

	return p
}

func (p *ByListTokenProvider) TokensStartIndex() int {
	if p.context != nil {
		return p.context.GetStart().GetTokenIndex()
	}

	return 0
}

func (p *ByListTokenProvider) Tokens() TokenList {
	return p.tokens
}

func (p *ByListTokenProvider) StartRuleIndex() int {
	if p.context != nil {
		return p.context.GetRuleIndex()
	}

	return 0
}
