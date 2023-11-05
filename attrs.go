package rex

import (
	"sort"
	"strings"
)

type (
	attrk string
	attrs []string
)

func (a attrs) key() attrk {
	return attrk(strings.Join(a, "|"))
}

func (a attrs) intersection(b attrs) map[string]struct{} {
	ma := map[string]struct{}{}
	for _, a := range a {
		ma[a] = struct{}{}
	}
	m := map[string]struct{}{}
	for _, b := range b {
		if _, ok := ma[b]; ok {
			m[b] = struct{}{}
		}
	}
	return m
}

func (a attrs) join(b attrs) attrs {
	m := map[string]struct{}{}
	for _, a := range []attrs{a, b} {
		for _, v := range a {
			m[v] = struct{}{}
		}
	}
	s := make(attrs, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	sort.Strings(s)
	return s
}

func (a attrk) split() attrs {
	return strings.Split(string(a), "|")
}
