package rex

type Table struct {
	schema    []string
	relations []*R
}

func NewTable(schema ...string) *Table { return newTable(schema) }
func newTable(s []string) *Table       { return &Table{schema: s} }

func (tbl *Table) SchemaInOrder() []string            { return tbl.schema }
func (tbl *Table) Schema() Schema                     { return newSchema(tbl.schema...) }
func (tbl *Table) Relations() []*R                    { return tbl.relations }
func (tbl *Table) Projection(schema ...string) *Table { return tbl.projection(schema) }

func (tbl *Table) projection(schema []string) *Table {
	if len(schema) == 0 {
		return tbl
	}
	p := newSchema(schema...)
	rs := []*R{}
	for _, r := range tbl.relations {
		if r.Schema().IsEqual(p) || r.Schema().IsSubset(p) {
			rs = append(rs, r)
		}
	}

	nt := newTable(schema)
	nt.relations = rs
	return nt
}

func relationsSchema(rs []*R) map[string]struct{} {
	m := map[string]struct{}{}
	for _, r := range rs {
		for k := range r.Schema() {
			m[k] = struct{}{}
		}
	}
	return m
}

func (tbl *Table) Equal(other *Table) bool {
	if len(tbl.relations) != len(other.relations) ||
		!tbl.Schema().IsEqual(other.Schema()) {
		return false
	}
	for i, r := range tbl.relations {
		if !r.Equal(other.relations[i]) {
			return false
		}
	}
	return true
}

func (tbl *Table) Add(t T) *Table {
	if !(t.Schema().IsEqual(tbl.Schema()) ||
		t.Schema().IsSubset(tbl.Schema())) {
		panic("schema mismatch")
	}

	r := tbl.tryFindCompatible(t)
	if r == nil {
		r = &R{}
		tbl.relations = append(tbl.relations, r)
	}
	r.Add(t)
	return tbl
}

func (tbl *Table) tryFindCompatible(t T) *R {
	for _, r := range tbl.relations {
		if t.Schema().IsEqual((*r)[0].Schema()) {
			return r
		}
	}
	return nil
}

func (tbl *Table) forEach(f func(t T)) {
	for _, r := range tbl.relations {
		for _, t := range *r {
			f(t)
		}
	}
}
