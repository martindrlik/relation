package rex

import (
	"fmt"
	"io"
	"strings"

	"github.com/martindrlik/rex/schema"
)

type boxTable struct {
	schema []string
	rows   []map[string]string
	max    map[string]int
}

func BoxTable[T any](schema map[string]T, relations []*R) interface{ String() string } {
	return makeBoxTable(schema, relations)
}

func makeBoxTable[T any](s map[string]T, rs []*R) interface{ String() string } {
	bt := &boxTable{
		schema: schema.Slice(s),
		rows:   []map[string]string{},
		max:    map[string]int{},
	}
	for _, s := range bt.schema {
		bt.max[s] = len(s)
	}
	for _, r := range rs {
		for _, t := range *r {
			bt.addRow(t)
		}
	}
	return bt
}

func (bt *boxTable) addRow(t T) {
	str := func(v any) string { return fmt.Sprintf("%v", v) }
	row := map[string]string{}
	for k, v := range t {
		s := str(v)
		if l := len(s); bt.max[k] < l {
			bt.max[k] = l
		}
		row[k] = s
	}
	bt.rows = append(bt.rows, row)
}

func (bt *boxTable) String() string {
	sb := &strings.Builder{}
	bt.writeHeader(sb)
	bt.writeRows(sb)
	return sb.String()
}

func (bt *boxTable) writeHeader(w io.Writer) {
	// ┏━━━━━━┯━━━━━━┓
	bt.writeRow(w, "┏", "┯", "┓", bt.schema, func(s string) string {
		return strings.Repeat("━", bt.max[s]+2)
	})

	// ┃    x │    y ┃
	bt.writeRow(w, "┃", "│", "┃", bt.schema, func(s string) string {
		return fmt.Sprintf(" %s ", bt.pad(s, s))
	})

	// ┠──────┼──────┨
	bt.writeRow(w, "┠", "┼", "┨", bt.schema, func(s string) string {
		return strings.Repeat("─", bt.max[s]+2)
	})
}

func (bt *boxTable) writeRows(w io.Writer) {
	for _, row := range bt.rows {
		// ┃ 2023 │ 2024 ┃
		bt.writeRow(w, "┃", "│", "┃", bt.schema, func(s string) string {
			v, ok := row[s]
			return fmt.Sprintf(" %s ", bt.pad(s, val(v, ok)))
		})
	}
	// ┗━━━━━━┷━━━━━━┛
	bt.writeRow(w, "┗", "┷", "┛", bt.schema, func(s string) string {
		return strings.Repeat("━", bt.max[s]+2)
	})
}

func val(v string, ok bool) string {
	if ok {
		return v
	}
	return "*"
}

func (bt *boxTable) writeRow(w io.Writer, left, middle, right string, schema []string, valueFunc func(string) string) {
	fmt.Fprint(w, left)
	for i, s := range bt.schema {
		if i > 0 {
			fmt.Fprint(w, middle)
		}
		fmt.Fprint(w, valueFunc(s))
	}
	fmt.Fprintln(w, right)
}

func (bt *boxTable) pad(s, v string) string {
	return fmt.Sprintf("%s%s", v, strings.Repeat(" ", bt.max[s]-len(v)))
}
