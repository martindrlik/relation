package rex

import (
	"fmt"
	"sort"
	"strings"
)

type (
	tuple  map[string]any
	tuples map[string][]tuple
)

func (t tuple) shallowCopy() tuple {
	v := map[string]any{}
	for k, w := range t {
		v[k] = w
	}
	return v
}

func (t tuple) key() (key string, isPartial bool) {
	k := strings.Builder{}
	a := keys(t)
	sort.Strings(a)
	for _, a := range a {
		v := t[a]
		if k.Len() > 0 {
			k.WriteString("\f")
		}
		if r, ok := v.(*Relation); ok {
			k.WriteString(r.key())
			isPartial = true
		} else {
			k.WriteString(fmt.Sprint(v))
		}
	}
	return k.String(), isPartial
}

func (t tuples) first() tuple {
	for _, t := range t {
		for _, t := range t {
			return t
		}
	}
	return tuple{}
}
