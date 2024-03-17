package table

type Table struct {
	schema []string
	tuples []map[string]any
}

func New(schema ...string) *Table {
	return &Table{schema: schema}
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
	ch := make(chan string)
	go func() {
		defer close(ch)
		for _, k := range t.schema {
			ch <- k
		}
		if len(t.schema) != 0 {
			return
		}
		for _, tuple := range t.tuples {
			for k := range tuple {
				ch <- k
			}
		}
	}()
	x := map[string]struct{}{}
	for attribute := range ch {
		x[attribute] = struct{}{}
	}
	return x
}

func (t *Table) Tuples() Tuples {
	return t.tuples
}

func (t *Table) CompleteTuples() Tuples {
	x := []map[string]any{}
	isComplete := t.isCompleteTuple()
	for _, tuple := range t.tuples {
		if isComplete(tuple) {
			x = append(x, tuple)
		}
	}
	return x
}

func (t *Table) isCompleteTuple() func(tuple map[string]any) bool {
	schema := t.Schema()
	return func(tuple map[string]any) bool {
		return len(schema) == len(tuple)
	}
}
