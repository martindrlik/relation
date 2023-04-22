package rex

import (
	"fmt"
	"io"
	"sort"
)

func (r *Relation) Serialize(w io.Writer) *Relation {
	k := keys(r.relations)
	sort.Strings(k)
	a := keys(r.attributes())
	sort.Strings(a)
	must(fmt.Fprint(w, "["))
	for _, k := range k {
		r := r.relations[k]
		tks := keys(r.tuples)
		sort.Strings(tks)
		for tki, tk := range tks {
			ts := r.tuples[tk]
			for ti, t := range ts {
				must(fmt.Fprint(w, "{"))
				writeTuple(w, a, t)
				if tki == len(tks)-1 && ti == len(ts)-1 {
					must(fmt.Fprint(w, "}"))
				} else {
					must(fmt.Fprint(w, "},\n"))
				}
			}
		}
	}
	must(fmt.Fprint(w, "]"))
	return r
}

func writeTuple(w io.Writer, attributes []string, t tuple) {
	first := true
	for _, a := range attributes {
		v, ok := t[a]
		if !ok {
			continue
		}
		if first {
			first = false
		} else {
			must(fmt.Fprint(w, ", "))
		}
		must(fmt.Fprintf(w, "%q: ", a))
		writeValue(w, v)
	}
}

func writeValue(w io.Writer, v any) {
	switch x := v.(type) {
	case string:
		must(fmt.Fprintf(w, "%q", x))
	case *Relation:
		x.Serialize(w)
	default:
		must(fmt.Fprintf(w, "%v", x))
	}
}

func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
