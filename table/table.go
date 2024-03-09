package table

type Table struct {
	tuples []map[string]any
}

func New() *Table {
	return &Table{}
}

func (t *Table) Add(tuples ...map[string]any) *Table {
	for _, tuple := range tuples {
		if len(tuple) != 0 && !t.Contains(tuple) {
			t.tuples = append(t.tuples, tuple)
		}
	}
	return t
}

func (t *Table) Contains(tuple map[string]any) bool {
	for _, t := range t.tuples {
		if tupleEqual(t, tuple) {
			return true
		}
	}
	return false
}

func (t *Table) Tuples() []map[string]any {
	return t.tuples
}
