package box

import (
	"fmt"
	"io"
	"strings"
	"unicode/utf8"
)

type boxTable struct {
	schema []string
	rows   []map[string]string
	max    map[string]int
}

func Table(schema []string, tuples ...map[string]any) interface{ String() string } {
	t := &boxTable{
		schema: schema,
		rows:   []map[string]string{},
		max:    map[string]int{},
	}
	for _, s := range t.schema {
		t.max[s] = utf8.RuneCountInString(s)
	}
	for _, tuple := range tuples {
		t.addRow(tuple)
	}
	return t
}

func (t *boxTable) addRow(tuple map[string]any) {
	str := func(v any) string { return fmt.Sprintf("%v", v) }
	row := map[string]string{}
	for k, v := range tuple {
		s := str(v)
		if l := utf8.RuneCountInString(s); t.max[k] < l {
			t.max[k] = l
		}
		row[k] = s
	}
	t.rows = append(t.rows, row)
}

func (t *boxTable) String() string {
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

func (t *boxTable) writeTop(w io.Writer) {
	// ┏━━━━━━┯━━━━━━┓
	t.writeRow(w, "┏", "┯", "┓", func(s string) string {
		return strings.Repeat("━", t.max[s]+2)
	})
}

func (t *boxTable) writeHeader(w io.Writer) {
	// ┃    x │    y ┃
	t.writeRow(w, "┃", "│", "┃", func(s string) string {
		return fmt.Sprintf(" %s ", t.pad(s, s))
	})
}

func (t *boxTable) writeSeparator(w io.Writer) {
	// ┠──────┼──────┨
	t.writeRow(w, "┠", "┼", "┨", func(s string) string {
		return strings.Repeat("─", t.max[s]+2)
	})
}

func (t *boxTable) writeRows(w io.Writer) {
	for _, row := range t.rows {
		// ┃ 2023 │ 2024 ┃
		t.writeRow(w, "┃", "│", "┃", func(s string) string {
			v, ok := row[s]
			return fmt.Sprintf(" %s ", t.pad(s, val(v, ok)))
		})
	}
}

func (t *boxTable) writeBottom(w io.Writer) {
	// ┗━━━━━━┷━━━━━━┛
	t.writeRow(w, "┗", "┷", "┛", func(s string) string {
		return strings.Repeat("━", t.max[s]+2)
	})
}

func val(v string, ok bool) string {
	if ok {
		return v
	}
	return "?"
}

func (t *boxTable) writeRow(w io.Writer, left, middle, right string, valueFunc func(string) string) {
	fmt.Fprint(w, left)
	for i, s := range t.schema {
		if i > 0 {
			fmt.Fprint(w, middle)
		}
		fmt.Fprint(w, valueFunc(s))
	}
	fmt.Fprintln(w, right)
}

func (bt *boxTable) pad(s, v string) string {
	return fmt.Sprintf("%s%s", v, strings.Repeat(" ", bt.max[s]-utf8.RuneCountInString(v)))
}
