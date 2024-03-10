package table

type Table struct {
	tuples []map[string]any
}

func New() *Table {
	return &Table{}
}

func (t *Table) Add(tuples ...map[string]any) *Table {
	for _, tuple := range tuples {
		if len(tuple) != 0 && !t.Tuples().Contains(tuple) {
			t.tuples = append(t.tuples, tuple)
		}
	}
	return t
}

func (t *Table) Schema() map[string]struct{} {
	x := map[string]struct{}{}
	for _, tuple := range t.tuples {
		for k := range tuple {
			x[k] = struct{}{}
		}
	}
	return x
}

func (t *Table) Tuples() Tuples {
	return t.tuples
}

func (t *Table) isCompleteTuple() func(tuple map[string]any) bool {
	schema := t.Schema()
	return func(tuple map[string]any) bool {
		return len(schema) == len(tuple)
	}
}
