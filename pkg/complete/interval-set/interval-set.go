// Copyright (c) 2012-2022 The ANTLR Project. All rights reserved.
// Use of this file is governed by the BSD 3-clause license that
// can be found in the LICENSE.txt file in the project root.

package IntervalSet

import (
	"reflect"
	"strconv"
	"strings"
	"unsafe"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type Interval struct {
	Start int
	Stop  int
}

/* stop is not included! */
// func antlr.NewInterval(start, stop int) *antlr.Interval {
// 	i := new(Interval)

// 	i.Start = start
// 	i.Stop = stop
// 	return i
// }

// func (i *antlr.Interval) Contains(item int) bool {
// 	return item >= i.Start && item < i.Stop
// }

// func (i *antlr.Interval) String() string {
// 	if i.Start == i.Stop-1 {
// 		return strconv.Itoa(i.Start)
// 	}

// 	return strconv.Itoa(i.Start) + ".." + strconv.Itoa(i.Stop-1)
// }

// func (i *antlr.Interval) length() int {
// 	return i.Stop - i.Start
// }

type IntervalSet struct {
	intervals []*antlr.Interval
	readOnly  bool
}

func NewIntervalSet() *IntervalSet {

	i := new(IntervalSet)

	i.intervals = nil
	i.readOnly = false

	return i
}

func FromAntlr(base *antlr.IntervalSet) *IntervalSet {

	set := NewIntervalSet()

	set.intervals = base.GetIntervals()
	set.readOnly = GetUnexportedField[*antlr.IntervalSet, bool](base, "readOnly")

	return set
}

func FromStartStop(start, stop int) *IntervalSet {
	set := NewIntervalSet()

	set.AddRange(start, stop)

	return set
}

func Of(item int) *IntervalSet {
	set := NewIntervalSet()

	set.AddOne(item)

	return set
}

func (i *IntervalSet) ToList() []int {
	list := make([]int, 0)
	if len(i.intervals) == 0 {
		return list
	}

	for _, interval := range i.intervals {
		for i := interval.Start; i < interval.Stop; i++ {
			list = append(list, i)
		}
	}

	return list
}

func (i *IntervalSet) First() int {
	if len(i.intervals) == 0 {
		return antlr.TokenInvalidType
	}

	return i.intervals[0].Start
}

func (i *IntervalSet) AddOne(v int) {
	i.AddInterval(antlr.NewInterval(v, v+1))
}

func (i *IntervalSet) AddRange(l, h int) {
	i.AddInterval(antlr.NewInterval(l, h+1))
}

func (i *IntervalSet) AddInterval(v *antlr.Interval) {
	if i.intervals == nil {
		i.intervals = make([]*antlr.Interval, 0)
		i.intervals = append(i.intervals, v)
	} else {
		// find insert pos
		for k, interval := range i.intervals {
			// distinct range -> insert
			if v.Stop < interval.Start {
				i.intervals = append(i.intervals[0:k], append([]*antlr.Interval{v}, i.intervals[k:]...)...)
				return
			} else if v.Stop == interval.Start {
				i.intervals[k].Start = v.Start
				return
			} else if v.Start <= interval.Stop {
				i.intervals[k] = antlr.NewInterval(intMin(interval.Start, v.Start), intMax(interval.Stop, v.Stop))

				// if not applying to end, merge potential overlaps
				if k < len(i.intervals)-1 {
					l := i.intervals[k]
					r := i.intervals[k+1]
					// if r contained in l
					if l.Stop >= r.Stop {
						i.intervals = append(i.intervals[0:k+1], i.intervals[k+2:]...)
					} else if l.Stop >= r.Start { // partial overlap
						i.intervals[k] = antlr.NewInterval(l.Start, r.Stop)
						i.intervals = append(i.intervals[0:k+1], i.intervals[k+2:]...)
					}
				}
				return
			}
		}
		// greater than any exiting
		i.intervals = append(i.intervals, v)
	}
}

func (i *IntervalSet) AddSet(other *IntervalSet) *IntervalSet {
	if other.intervals != nil {
		for k := 0; k < len(other.intervals); k++ {
			i2 := other.intervals[k]
			i.AddInterval(antlr.NewInterval(i2.Start, i2.Stop))
		}
	}
	return i
}

func (i *IntervalSet) Complement(start int, stop int) *IntervalSet {
	result := NewIntervalSet()
	result.AddInterval(antlr.NewInterval(start, stop+1))
	for j := 0; j < len(i.intervals); j++ {
		result.RemoveRange(i.intervals[j])
	}
	return result
}

func (i *IntervalSet) Contains(item int) bool {
	if i.intervals == nil {
		return false
	}
	for k := 0; k < len(i.intervals); k++ {
		if i.intervals[k].Contains(item) {
			return true
		}
	}
	return false
}

func (i *IntervalSet) Length() int {
	len := 0

	for _, v := range i.intervals {
		len += intervalLen(v)
	}

	return len
}

func (i *IntervalSet) RemoveRange(v *antlr.Interval) {
	if v.Start == v.Stop-1 {
		i.RemoveOne(v.Start)
	} else if i.intervals != nil {
		k := 0
		for n := 0; n < len(i.intervals); n++ {
			ni := i.intervals[k]
			// intervals are ordered
			if v.Stop <= ni.Start {
				return
			} else if v.Start > ni.Start && v.Stop < ni.Stop {
				i.intervals[k] = antlr.NewInterval(ni.Start, v.Start)
				x := antlr.NewInterval(v.Stop, ni.Stop)
				// i.intervals.splice(k, 0, x)
				i.intervals = append(i.intervals[0:k], append([]*antlr.Interval{x}, i.intervals[k:]...)...)
				return
			} else if v.Start <= ni.Start && v.Stop >= ni.Stop {
				//                i.intervals.splice(k, 1)
				i.intervals = append(i.intervals[0:k], i.intervals[k+1:]...)
				k = k - 1 // need another pass
			} else if v.Start < ni.Stop {
				i.intervals[k] = antlr.NewInterval(ni.Start, v.Start)
			} else if v.Stop < ni.Stop {
				i.intervals[k] = antlr.NewInterval(v.Stop, ni.Stop)
			}
			k++
		}
	}
}

func (i *IntervalSet) RemoveOne(v int) {
	if i.intervals != nil {
		for k := 0; k < len(i.intervals); k++ {
			ki := i.intervals[k]
			// intervals i ordered
			if v < ki.Start {
				return
			} else if v == ki.Start && v == ki.Stop-1 {
				//				i.intervals.splice(k, 1)
				i.intervals = append(i.intervals[0:k], i.intervals[k+1:]...)
				return
			} else if v == ki.Start {
				i.intervals[k] = antlr.NewInterval(ki.Start+1, ki.Stop)
				return
			} else if v == ki.Stop-1 {
				i.intervals[k] = antlr.NewInterval(ki.Start, ki.Stop-1)
				return
			} else if v < ki.Stop-1 {
				x := antlr.NewInterval(ki.Start, v)
				ki.Start = v + 1
				//				i.intervals.splice(k, 0, x)
				i.intervals = append(i.intervals[0:k], append([]*antlr.Interval{x}, i.intervals[k:]...)...)
				return
			}
		}
	}
}

func (i *IntervalSet) String() string {
	return i.StringVerbose(nil, nil, false)
}

func (i *IntervalSet) StringVerbose(literalNames []string, symbolicNames []string, elemsAreChar bool) string {

	if i.intervals == nil {
		return "{}"
	} else if literalNames != nil || symbolicNames != nil {
		return i.toTokenString(literalNames, symbolicNames)
	} else if elemsAreChar {
		return i.toCharString()
	}

	return i.toIndexString()
}

func (i *IntervalSet) GetIntervals() []*antlr.Interval {
	return i.intervals
}

func (i *IntervalSet) toCharString() string {
	names := make([]string, len(i.intervals))

	var sb strings.Builder

	for j := 0; j < len(i.intervals); j++ {
		v := i.intervals[j]
		if v.Stop == v.Start+1 {
			if v.Start == antlr.TokenEOF {
				names = append(names, "<EOF>")
			} else {
				sb.WriteByte('\'')
				sb.WriteRune(rune(v.Start))
				sb.WriteByte('\'')
				names = append(names, sb.String())
				sb.Reset()
			}
		} else {
			sb.WriteByte('\'')
			sb.WriteRune(rune(v.Start))
			sb.WriteString("'..'")
			sb.WriteRune(rune(v.Stop - 1))
			sb.WriteByte('\'')
			names = append(names, sb.String())
			sb.Reset()
		}
	}
	if len(names) > 1 {
		return "{" + strings.Join(names, ", ") + "}"
	}

	return names[0]
}

func (i *IntervalSet) toIndexString() string {

	names := make([]string, 0)
	for j := 0; j < len(i.intervals); j++ {
		v := i.intervals[j]
		if v.Stop == v.Start+1 {
			if v.Start == antlr.TokenEOF {
				names = append(names, "<EOF>")
			} else {
				names = append(names, strconv.Itoa(v.Start))
			}
		} else {
			names = append(names, strconv.Itoa(v.Start)+".."+strconv.Itoa(v.Stop-1))
		}
	}
	if len(names) > 1 {
		return "{" + strings.Join(names, ", ") + "}"
	}

	return names[0]
}

func (i *IntervalSet) toTokenString(literalNames []string, symbolicNames []string) string {
	names := make([]string, 0)
	for _, v := range i.intervals {
		for j := v.Start; j < v.Stop; j++ {
			names = append(names, i.elementName(literalNames, symbolicNames, j))
		}
	}
	if len(names) > 1 {
		return "{" + strings.Join(names, ", ") + "}"
	}

	return names[0]
}

func (i *IntervalSet) elementName(literalNames []string, symbolicNames []string, a int) string {
	if a == antlr.TokenEOF {
		return "<EOF>"
	} else if a == antlr.TokenEpsilon {
		return "<EPSILON>"
	} else {
		if a < len(literalNames) && literalNames[a] != "" {
			return literalNames[a]
		}

		return symbolicNames[a]
	}
}

func intMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func intMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func GetUnexportedField[E any, R any](entity E, fieldName string) R {
	field := reflect.ValueOf(entity).Elem().FieldByName(fieldName)
	return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Interface().(R)
}

func intervalLen(interval *antlr.Interval) int {
	return interval.Stop - interval.Start
}
