package rex

import (
	"fmt"
	"sort"
	"strings"
)

func Dump(r R, paddings ...int) string {
	m := map[string]struct{}{}
	for _, r := range r {
		for _, a := range r.attributes {
			m[a] = struct{}{}
		}
	}
	if len(m) > len(paddings) {
		panic(fmt.Sprintf("missing padding for %d attribute(s)", len(m)-len(paddings)))
	}
	pad := func(pi int, v string) string {
		count := paddings[pi] - len([]rune(v))
		if count <= 0 {
			return ""
		}
		return strings.Repeat(" ", count)
	}
	o := make([]string, 0, len(m))
	for a := range m {
		o = append(o, a)
	}
	sort.Strings(o)
	var b strings.Builder
	b.WriteString(dumpattrs(o, pad))
	b.WriteRune('\n')
	for _, r := range r {
		ai := r.attri()
		for _, t := range r.tuples {
			b.WriteString(dumptuple(o, ai, t, pad))
			b.WriteRune('\n')
		}
	}
	return b.String()
}

func dumpattrs(o []string, pad func(pi int, v string) string) string {
	var b strings.Builder
	for i := 0; i < len(o); i++ {
		if b.Len() > 0 {
			b.WriteString(" | ")
		}
		b.WriteString(o[i])
		if i < len(o)-1 {
			b.WriteString(pad(i, o[i]))
		}
	}
	return b.String()
}

func dumptuple(o []string, ai map[string]int, t []any, pad func(pi int, v string) string) string {
	var b strings.Builder
	for pi := 0; pi < len(o); pi++ {
		if b.Len() > 0 {
			b.WriteString(" | ")
		}
		var v string
		if i, ok := ai[o[pi]]; ok {
			v = fmt.Sprintf("%v", t[i])
		} else {
			v = "✕"
		}
		b.WriteString(v)
		if pi < len(o)-1 {
			b.WriteString(pad(pi, v))
		}
	}
	return b.String()
}
