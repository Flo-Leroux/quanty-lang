package complete

import (
	"fmt"

	IntervalSet "github.com/Flo-Leroux/quanty-lang/pkg/complete/interval-set"
	list "github.com/Flo-Leroux/quanty-lang/pkg/complete/list"
)

// A record for a follow set along with the path at which this set was found.
// If there is only a single symbol in the interval set then we also collect and store tokens which follow
// this symbol directly in its rule (i.e. there is no intermediate rule transition). Only single label transitions
// are considered. This is useful if you have a chain of tokens which can be suggested as a whole, because there is
// a fixed sequence in the grammar.
type FollowSetWithPath struct {
	intervals *IntervalSet.IntervalSet
	path      RuleList
	following TokenList
}

func NewFollowSetWithPath() *FollowSetWithPath {
	return &FollowSetWithPath{
		intervals: nil,
		path:      NewRuleList(),
		following: NewTokenList(),
	}
}

func (f *FollowSetWithPath) String() string {
	return fmt.Sprintf("FollowSetWithPath(intervals=%s, path=%s, following=%s)", f.intervals, f.path, f.following)
}

type FollowSetWithPathList = *list.List[*FollowSetWithPath]

func NewFollowSetWithPathList() FollowSetWithPathList {
	return list.New[*FollowSetWithPath](func(a, b *FollowSetWithPath) bool { return a.String() == b.String() })
}
