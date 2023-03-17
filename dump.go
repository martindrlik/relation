package rex

import (
	"bytes"
	"fmt"
	"strings"
)

func Dump(r R, options ...func(*DumpOptions)) string {
	opt := buildDumpOptions(options...)
	oa := r.attributes()
	pad := func(a, v string) string {
		var p int
		if n, ok := opt.padding[a]; ok {
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

	for i := 0; i < len(oa); i++ {
		if i > 0 {
			b.WriteString(" | ")
		}
		b.WriteString(oa[i])
		if i < len(oa)-1 {
			b.WriteString(pad(oa[i], oa[i]))
		}
	}

	nested := []R{}
	ref := func(a any) any {
		if r, ok := a.(R); ok {
			nested = append(nested, r)
			return fmt.Sprintf("*R%d", len(nested))
		}
		return a
	}

	for _, k := range r.keyOrder() {
		r := r[k]
		ai := r.attri()
		for _, t := range r.tuples {
			b.WriteRune('\n')
			for i := 0; i < len(oa); i++ {
				if i > 0 {
					b.WriteString(" | ")
				}
				var v string
				if i, ok := ai[oa[i]]; ok {
					v = fmt.Sprintf("%v", ref(t[i]))
				} else {
					v = "✕"
				}
				b.WriteString(v)
				if i < len(oa)-1 {
					b.WriteString(pad(oa[i], v))
				}
			}
		}
	}

	for i, r := range nested {
		b.WriteString(fmt.Sprintf("\n-- R%d:\n", i+1))
		b.WriteString(Dump(r, options...))
	}
	return b.String()
}
