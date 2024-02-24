package rex

import (
	"fmt"
	"io"
	"strings"
)

type boxTable struct {
	schema []string
	rows   []map[string]string
	max    map[string]int
}

func BoxTable(schema []string, relations []*R) interface{ String() string } {
	return makeBoxTable(schema, relations)
}

func makeBoxTable(s []string, rs []*R) interface{ String() string } {
	bt := &boxTable{
		schema: s,
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
	bt.writeTop(sb)
	bt.writeHeader(sb)
	if len(bt.rows) > 0 {
		bt.writeSeparator(sb)
		bt.writeRows(sb)
	}
	bt.writeBottom(sb)
	return sb.String()
}

func (bt *boxTable) writeTop(w io.Writer) {
	// ┏━━━━━━┯━━━━━━┓
	bt.writeRow(w, "┏", "┯", "┓", bt.schema, func(s string) string {
		return strings.Repeat("━", bt.max[s]+2)
	})
}

func (bt *boxTable) writeHeader(w io.Writer) {
	// ┃    x │    y ┃
	bt.writeRow(w, "┃", "│", "┃", bt.schema, func(s string) string {
		return fmt.Sprintf(" %s ", bt.pad(s, s))
	})
}

func (bt *boxTable) writeSeparator(w io.Writer) {
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
}

func (bt *boxTable) writeBottom(w io.Writer) {
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
