package rex

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"sync"
)

var bufPool = sync.Pool{
	New: func() any { return new(bytes.Buffer) },
}

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

	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	defer func() { bufPool.Put(b) }()

	b.WriteString(dumpattrs(o, pad))

	for _, k := range r.keyOrder() {
		r := r[k]
		ai := r.attri()
		for _, t := range r.tuples {
			b.WriteRune('\n')
			b.WriteString(dumptuple(o, ai, t, pad))
		}
	}
	return b.String()
}

func dumpattrs(o []string, pad func(pi int, v string) string) string {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	defer func() { bufPool.Put(b) }()

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
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	defer func() { bufPool.Put(b) }()

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
