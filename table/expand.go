package table

import (
	"github.com/martindrlik/rex/relation"
	"github.com/martindrlik/rex/require"
	"github.com/martindrlik/rex/schema"
)

func (t *Table) Expand(domains *Table) (*Table, error) {
	attributes := t.Schema().Attributes()
	return expand(t.Project(attributes[0], attributes[1:]...), domains)
}

func expand(x, domains *Table) (*Table, error) {
	replace := func(old, new *relation.Relation) {
		for i, r := range x.r {
			if r == old {
				x.r[i] = new
				break
			}
		}
	}

	for _, a := range domains.Schema().Attributes() {
		ds := domains.Relations(Matching(schema.New(a)))
		if len(ds) != 1 {
			return nil, ErrDomainTable(a)
		}
		rs := x.Relations(Missing(a))
		for _, r := range rs {
			replace(r, require.NoError(r.NaturalJoin(ds[0])))
		}
	}
	return x, nil
}
