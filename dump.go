package rex

import (
	"bytes"
	"fmt"
	"strings"
)

type DumpOptions struct {
	padding map[string]int
}

func Pad(name string, pad int) func(*DumpOptions) {
	return func(do *DumpOptions) {
		do.padding[name] = pad
	}
}

func Dump(r R, options ...func(*DumpOptions)) string {
	do := &DumpOptions{
		padding: map[string]int{},
	}
	for _, option := range options {
		option(do)
	}
	o := r.attributes()
	pad := func(a, v string) string {
		var p int
		if n, ok := do.padding[a]; ok {
			p = n
		} else {
			p = len([]rune(a))
		}
		count := p - len([]rune(v))
		if count <= 0 {
			return ""
		}
		return strings.Repeat(" ", count)
	}

	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	defer func() { bufPool.Put(b) }()

	b.WriteString(dumpattrs(o, pad))

	inner := []R{}
	ref := func(a any) any {
		if r, ok := a.(R); ok {
			inner = append(inner, r)
			return fmt.Sprintf("*r%d", len(inner))
		}
		return a
	}

	for _, k := range r.keyOrder() {
		r := r[k]
		ai := r.attri()
		for _, t := range r.tuples {
			b.WriteRune('\n')
			b.WriteString(dumptuple(
				o,
				ai,
				t,
				ref,
				pad))
		}
	}

	for i, r := range inner {
		b.WriteString(fmt.Sprintf("\n-- r%d:\n", i+1))
		b.WriteString(Dump(r, options...))
	}
	return b.String()
}

func dumpattrs(o []string, pad func(a, v string) string) string {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	defer func() { bufPool.Put(b) }()

	for i := 0; i < len(o); i++ {
		if b.Len() > 0 {
			b.WriteString(" | ")
		}
		b.WriteString(o[i])
		if i < len(o)-1 {
			b.WriteString(pad(o[i], o[i]))
		}
	}
	return b.String()
}

func dumptuple(
	o []string,
	ai map[string]int,
	t []any,
	ref func(any) any,
	pad func(a, v string) string) string {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	defer func() { bufPool.Put(b) }()

	for pi := 0; pi < len(o); pi++ {
		if b.Len() > 0 {
			b.WriteString(" | ")
		}
		var v string
		if i, ok := ai[o[pi]]; ok {
			v = fmt.Sprintf("%v", ref(t[i]))
		} else {
			v = "✕"
		}
		b.WriteString(v)
		if pi < len(o)-1 {
			b.WriteString(pad(o[pi], v))
		}
	}
	return b.String()
}
