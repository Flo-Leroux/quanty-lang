package complete

import (
	"fmt"

	IntervalSet "github.com/Flo-Leroux/quanty-lang/pkg/complete/interval-set"
)

// A list of follow sets (for a given state number) + all of them combined for quick hit tests.
// This data is static in nature (because the used ATN states are part of a static struct: the ATN).
// Hence it can be shared between all C3 instances, however it dependes on the actual parser class (type).
type FollowSetsHolder struct {
	sets     FollowSetWithPathList
	combined *IntervalSet.IntervalSet
}

func NewFollowSetsHolder() *FollowSetsHolder {
	return &FollowSetsHolder{
		sets:     NewFollowSetWithPathList(),
		combined: nil,
	}
}

func (f *FollowSetsHolder) String() string {
	return fmt.Sprintf("FollowSetsHolder(sets=%s, combined=%s)", f.sets, f.combined)
}
