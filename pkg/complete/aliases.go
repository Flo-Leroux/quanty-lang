package complete

import (
	list "github.com/Flo-Leroux/quanty-lang/pkg/complete/list"
	mutablelist "github.com/Flo-Leroux/quanty-lang/pkg/complete/mutable-list"
	mutablemap "github.com/Flo-Leroux/quanty-lang/pkg/complete/mutable-map"
	mutableset "github.com/Flo-Leroux/quanty-lang/pkg/complete/mutable-set"
)

type (
	TokenKind = int
	RuleIndex = int
	TokenList = *mutablelist.MutableList[TokenKind]
	RuleList  = *mutablelist.MutableList[RuleIndex]

	FollowSetsPerState = *mutablemap.MutableMap[int, *FollowSetsHolder]

	RuleEndStatus = *mutableset.MutableSet[int]

	TokensProvider interface {
		Tokens() TokenList
		TokensStartIndex() int
		StartRuleIndex() int
	}

	RuleKind         = int
	RuleKindList     = *list.List[RuleKind]
	ParserStack      = *list.List[RuleKind]
	CompletionOption struct {
		TokenKind   TokenKind
		ParserStack ParserStack
	}
)

func NewTokenList() TokenList {
	return mutablelist.New[TokenKind](func(a, b int) bool { return a == b })
}

func NewRuleList() RuleList {
	return mutablelist.New[RuleIndex](func(a, b int) bool { return a == b })
}

func NewParserStack() ParserStack {
	return list.New[RuleKind](func(a, b int) bool { return a == b })
}

func NewRuleEndStatus() RuleEndStatus {
	return mutableset.New[int](func(a, b int) bool { return a == b })
}

var atnStateTypeMap = []string{
	"invalid",
	"basic",
	"rule start",
	"block start",
	"plus block start",
	"star block start",
	"token start",
	"rule stop",
	"block end",
	"star loop back",
	"star loop entry",
	"plus loop back",
	"loop end",
}
