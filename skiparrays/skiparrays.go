package skiparrays

import (
	"github.com/emirpasic/gods/lists/arraylist"
)

type SkipArrayIf interface {
	Len() int
	Insert(pos int, value interface{})
	Remove(pos int)
	Values() []interface{}
}

type row struct {
	pos      int
	elements *arraylist.List
}

type SkipArray struct {
	rows   []*row
	rowLen int
}

func newRow(pos int) *row {
	r := &row{}
	r.pos = pos
	r.elements = arraylist.New()
	return r
}

var rowLen = 1024

func New(rowlen int, values ...interface{}) *SkipArray {
	a := &SkipArray{}
	if rowlen != 0 {
		a.rowLen = rowlen
	} else {
		a.rowLen = rowLen
	}
	numRows := (len(values) / a.rowLen) + 1
	a.rows = make([]*row, numRows)
	s := 0
	e := min(a.rowLen, len(values))
	for i := 0; i < numRows; i++ {
		a.rows[i] = newRow(s)
		if e > s {
			a.rows[i].elements.Add(values[s:e]...)
		}
		s = e
		e = min(s+a.rowLen, len(values))
	}
	return a
}

func (a *SkipArray) findRow(pos int) int {
	if len(a.rows) == 1 {
		return 0
	}
	s := 0
	e := len(a.rows) - 1

	for s < e {
		mid := s + (e-s)/2
		if a.rows[mid].pos <= pos && mid < e && pos < a.rows[mid+1].pos {
			return mid
		}
		if pos < a.rows[mid].pos {
			e = mid
		} else {
			s = mid + 1
		}
	}
	return s
}

func (a *SkipArray) updateIndexes(pos int, delta int) {
	for i := pos + 1; i < len(a.rows)-1; i++ {
		a.rows[i].pos += delta
	}
}

func (a *SkipArray) Insert(pos int, value interface{}) {
	if pos < 0 {
		return
	}
	idx := a.findRow(pos)
	a.rows[idx].elements.Insert(pos-a.rows[idx].pos, value)
	a.updateIndexes(idx, 1)
}

func (a *SkipArray) Set(pos int, value interface{}) {
	if pos < 0 {
		return
	}
	idx := a.findRow(pos)
	a.rows[idx].elements.Set(pos-a.rows[idx].pos, value)
}

func (a *SkipArray) Get(pos int) (interface{}, bool) {
	if pos < 0 {
		return nil, false
	}
	idx := a.findRow(pos)
	return a.rows[idx].elements.Get(pos - a.rows[idx].pos)
}

func (a *SkipArray) Add(value interface{}) {
	n := len(a.rows) - 1
	a.rows[n].elements.Add(value)
}

func (a *SkipArray) dropRow(idx int) {
	copy(a.rows[idx:], a.rows[idx+1:])
	a.rows = a.rows[:len(a.rows)-1]
}

func (a *SkipArray) Remove(pos int) {
	idx := a.findRow(pos)
	a.rows[idx].elements.Remove(pos - a.rows[idx].pos)
	a.updateIndexes(idx, -1)
	if a.rows[idx].elements.Size() == 0 && idx > 0 {
		a.dropRow(idx)
	}
}

func (a *SkipArray) Values() []interface{} {
	sz := a.Size()
	vals := make([]interface{}, sz)
	pos := 0
	for i := range a.rows {
		for _, v := range a.rows[i].elements.Values() {
			vals[pos] = v
			pos++
		}
	}
	return vals
}

func (a *SkipArray) Size() int {
	var sz int
	for i := range a.rows {
		sz += a.rows[i].elements.Size()
	}
	return sz
}

func (a *SkipArray) Iterator() *SkipArrayItr {
	return newIterator(a)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type SkipArrayItr struct {
	arr   *SkipArray
	crow  int
	cpos  int
	index int
}

func newIterator(arr *SkipArray) *SkipArrayItr {
	ski := &SkipArrayItr{arr: arr}
	return ski
}

func (ski *SkipArrayItr) Index() int {
	return ski.index
}

func (ski *SkipArrayItr) Value() interface{} {
	val, _ := ski.arr.Get(ski.index)
	return val
}

func (ski *SkipArrayItr) Next() bool {
	ski.index++
	return ski.index < ski.arr.Size()
}
