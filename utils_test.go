package rex_test

import (
	"time"

	"github.com/martindrlik/rex"
)

func show(show string) func() (string, any) { return pair("show", show) }
func name(name string) func() (string, any) { return pair("name", name) }

func born(year, month, day int) func() (string, any) {
	return pair("born", time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))
}

func pair(name string, value any) func() (string, any) {
	return func() (string, any) {
		return name, value
	}
}

func glue(a ...func() (string, any)) map[string]any {
	m := map[string]any{}
	for _, f := range a {
		k, v := f()
		m[k] = v
	}
	return m
}

func gluem(a ...map[string]any) map[string]any {
	m := map[string]any{}
	for _, mm := range a {
		for k, v := range mm {
			m[k] = v
		}
	}
	return m
}

func in(a ...map[string]any) *rex.Relation {
	nr := rex.NewRelation()
	return nr.Insert(a...)
}

func take2(r *rex.Relation) (tuplex, tuplex) {
	txs := take(r, 2)
	return txs[0], txs[1]
}

func take1(r *rex.Relation) tuplex {
	txs := take(r, 1)
	return txs[0]
}

func take(r *rex.Relation, n int) []tuplex {
	txs := make([]tuplex, 0, n)
	r.Each(func(m map[string]any, isPossible bool) bool {
		txs = append(txs, tuplex{m, isPossible})
		n--
		return n > 0
	})
	return txs
}

func attr(fn func() (string, any)) string {
	k, _ := fn()
	return k
}

type tuplex struct {
	m          map[string]any
	isPossible bool
}
