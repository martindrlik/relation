package box

import (
	"fmt"
	"io"
	"strings"

	"github.com/martindrlik/rex/relation"
	"github.com/martindrlik/rex/tuple"
)

type table struct {
	schema []string
	rows   []map[string]string
	max    map[string]int
}

func Table(schema []string, relations []*relation.Relation) interface{ String() string } {
	t := &table{
		schema: schema,
		rows:   []map[string]string{},
		max:    map[string]int{},
	}
	for _, s := range t.schema {
		t.max[s] = len(s)
	}
	for _, r := range relations {
		for _, u := range r.Tuples() {
			t.addRow(u)
		}
	}
	return t
}

func (t *table) addRow(u tuple.T) {
	str := func(v any) string { return fmt.Sprintf("%v", v) }
	row := map[string]string{}
	for k, v := range u {
		s := str(v)
		if l := len(s); t.max[k] < l {
			t.max[k] = l
		}
		row[k] = s
	}
	t.rows = append(t.rows, row)
}

func (t *table) String() string {
	sb := &strings.Builder{}
	t.writeTop(sb)
	t.writeHeader(sb)
	if len(t.rows) > 0 {
		t.writeSeparator(sb)
		t.writeRows(sb)
	}
	t.writeBottom(sb)
	return sb.String()
}

func (t *table) writeTop(w io.Writer) {
	// ┏━━━━━━┯━━━━━━┓
	t.writeRow(w, "┏", "┯", "┓", func(s string) string {
		return strings.Repeat("━", t.max[s]+2)
	})
}

func (t *table) writeHeader(w io.Writer) {
	// ┃    x │    y ┃
	t.writeRow(w, "┃", "│", "┃", func(s string) string {
		return fmt.Sprintf(" %s ", t.pad(s, s))
	})
}

func (t *table) writeSeparator(w io.Writer) {
	// ┠──────┼──────┨
	t.writeRow(w, "┠", "┼", "┨", func(s string) string {
		return strings.Repeat("─", t.max[s]+2)
	})
}

func (t *table) writeRows(w io.Writer) {
	for _, row := range t.rows {
		// ┃ 2023 │ 2024 ┃
		t.writeRow(w, "┃", "│", "┃", func(s string) string {
			v, ok := row[s]
			return fmt.Sprintf(" %s ", t.pad(s, val(v, ok)))
		})
	}
}

func (t *table) writeBottom(w io.Writer) {
	// ┗━━━━━━┷━━━━━━┛
	t.writeRow(w, "┗", "┷", "┛", func(s string) string {
		return strings.Repeat("━", t.max[s]+2)
	})
}

func val(v string, ok bool) string {
	if ok {
		return v
	}
	return "*"
}

func (t *table) writeRow(w io.Writer, left, middle, right string, valueFunc func(string) string) {
	fmt.Fprint(w, left)
	for i, s := range t.schema {
		if i > 0 {
			fmt.Fprint(w, middle)
		}
		fmt.Fprint(w, valueFunc(s))
	}
	fmt.Fprintln(w, right)
}

func (bt *table) pad(s, v string) string {
	return fmt.Sprintf("%s%s", v, strings.Repeat(" ", bt.max[s]-len(v)))
}
