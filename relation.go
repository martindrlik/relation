package rex

import (
	"strings"
)

type (
	Relation map[attrsKey]relation
	relation []tuplex

	attrsKey string
	attrs    []string
)

// NewRelation returns new empty relation.
func NewRelation() *Relation { return &Relation{} }

// Contains returns true if relation contains tuple given by tm map.
func (r *Relation) Contains(tm map[string]any) bool {
	k, tx := tupleMap(tm).ktx()
	return (*r)[k].contains(tx)
}

// Insert inserts tuples given by variadic a.
func (r *Relation) Insert(a ...map[string]any) *Relation {
	for _, tm := range a {
		r.insertTuplex(tupleMap(tm).ktx())
	}
	return r
}

// Each calls f for each tuple. If f returns false, iteration stops.
func (r *Relation) Each(f func(map[string]any, bool) bool) {
	for k, r := range *r {
		a := k.split()
		for _, tx := range r {
			if !f(tx.toMap(a), tx.isPossible) {
				break
			}
		}
	}
}

func (r relation) contains(tx tuplex) bool {
	b := false
	r.each(func(i int, ctx tuplex) bool {
		if tx.equal(ctx.tuple) {
			b = true
			return false
		}
		return true
	})
	return b
}

func (r relation) each(f func(int, tuplex) bool) {
	for i, t := range r {
		if !f(i, t) {
			break
		}
	}
}

func (r *Relation) insertTuplex(k attrsKey, tx tuplex) *Relation {
	updated := false
	(*r)[k].each(func(i int, ctx tuplex) bool {
		if tx.equal(ctx.tuple) {
			ctx.meta = tx.meta
			(*r)[k][i] = ctx
			updated = true
			return false
		}
		return true
	})
	if !updated {
		(*r)[k] = append((*r)[k], tx)
	}
	return r
}

func (a attrs) key() attrsKey {
	return attrsKey(strings.Join(a, "|"))
}

func (a attrs) indexMap() map[string]int {
	m := make(map[string]int)
	for i, v := range a {
		m[v] = i
	}
	return m
}

func (a attrsKey) split() attrs {
	return strings.Split(string(a), "|")
}
