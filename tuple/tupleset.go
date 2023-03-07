package tuple

import (
	"maps"
	"slices"
)

type TupleSet []Tuple

func (ts *TupleSet) Add(t Tuple) {
	*ts = append(*ts, t)
}

func (ts *TupleSet) Has(t Tuple) bool {
	_, ok := ts.index(t)
	return ok
}

func (ts *TupleSet) Delete(t Tuple) {
	i, ok := ts.index(t)
	if ok {
		*ts = slices.Delete(*ts, i, i+1)
	}
}

func (ts *TupleSet) index(t Tuple) (int, bool) {
	for i, t0 := range *ts {
		if maps.Equal(t0, t) {
			return i, true
		}
	}
	return 0, false
}
