package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"maps"
	"os"
	"slices"
	"strings"

	"github.com/martindrlik/rex/box"
	"github.com/martindrlik/rex/persist"
	"github.com/martindrlik/rex/table"
)

func main() {
	defer func() {
		switch x := recover().(type) {
		case error:
			fmt.Fprintf(os.Stderr, "error: %v\n", x)
		case nil:
		default:
			panic(x)
		}
	}()
	must := func(t *table.Table, err error) *table.Table {
		if err != nil {
			panic(err)
		}
		return t
	}
	bind("union", "", func(a, b *table.Table) *table.Table { return must(a.Union(b)) })
	bind("difference", "", func(a, b *table.Table) *table.Table { return must(a.Difference(b)) })
	bind("natural-join", "", func(a, b *table.Table) *table.Table { return a.NaturalJoin(b) })
	exec(parse(os.Args[1:]))
}

func exec(op string, tables []*table.Table, outputFormat string, schema []string) {
	func(fn func([]*table.Table) []*table.Table) {
		for _, t := range fn(tables) {
			projection(t, outputFormat, schema...)
		}
	}(binaryOp(op))
}

func binaryOp(op string) func([]*table.Table) []*table.Table {
	return aggr(func(a, b *table.Table) *table.Table {
		if desc, ok := ops[op]; ok {
			return desc.fn(a, b)
		}
		panic("unreachable")
	})
}

func aggr(fn func(a, b *table.Table) *table.Table) func([]*table.Table) []*table.Table {
	return func(tables []*table.Table) []*table.Table {
		result := tables[0]
		for _, t := range tables[1:] {
			result = fn(result, t)
		}
		return []*table.Table{result}
	}
}

func projection(t *table.Table, outputFormat string, schema ...string) {
	if len(schema) == 0 {
		schema = slices.Collect(maps.Keys(t.Schema))
	}
	w, _ := t.Projection(schema...)
	switch outputFormat {
	case "json":
		if err := persist.WriteJson(os.Stdout, w); err != nil {
			fmt.Fprintf(os.Stderr, "unable to write json output: %v", err)
		}
	case "table":
		fmt.Println(box.Relation(schema, w.List()))
	}
}

func parse(args []string) (string, []*table.Table, string, []string) {
	if len(args) < 2 {
		usage(errors.New("missing arguments"))
	}
	fs := flag.NewFlagSet("", flag.ExitOnError)
	var (
		schemalessFilenames = stringsFlag{}
		schemalessInlines   = stringsFlag{}

		outputFormat = fs.String("of", "table", "table or json")
	)
	fs.Var(&schemalessFilenames, "fa", "name of file that contains array of tuples")
	fs.Var(&schemalessInlines, "ia", "inline array of tuples")

	op := args[0]
	_, ok := ops[op]
	if !ok {
		usage(fmt.Errorf("unknown op: %s", op))
	}

	fs.Parse(args[1:])
	if len(schemalessFilenames) == 0 && len(schemalessInlines) == 0 {
		usage(errors.New("missing table"))
	}

	tables := []*table.Table{}
	load := func(r io.Reader, fn func(io.Reader) (*table.Table, error)) error {
		t, err := fn(r)
		if err != nil {
			return err
		}
		tables = append(tables, t)
		return nil
	}
	loadFile := func(name string, fn func(io.Reader) (*table.Table, error)) error {
		f, err := os.Open(name)
		if err != nil {
			return err
		}
		defer f.Close()
		return load(f, fn)
	}
	loadFiles := func(filenames []string, fn func(io.Reader) (*table.Table, error)) {
		for _, name := range filenames {
			if err := loadFile(name, fn); err != nil {
				usage(fmt.Errorf("loading file: %w", err))
			}
		}
	}
	loadInline := func(inlines []string, fn func(io.Reader) (*table.Table, error)) {
		for _, inline := range inlines {
			t, err := fn(strings.NewReader(inline))
			if err != nil {
				usage(fmt.Errorf("loading inline %v: %w", inline, err))
			}
			tables = append(tables, t)
		}
	}

	loadFiles(schemalessFilenames, persist.Load)
	loadInline(schemalessInlines, persist.Load)

	return op, tables, *outputFormat, fs.Args()
}

type stringsFlag []string

func (s *stringsFlag) String() string {
	return fmt.Sprint(*s)
}

func (s *stringsFlag) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func usage(err error) {
	if err != nil {
		fmt.Println("Error:")
		fmt.Printf("	%v\n", err)
	}
	fmt.Println("Usage:")
	fmt.Println("	rex <command> <input> <options> [attribute ...]")
	fmt.Println("Commands:")
	names := slices.Collect(maps.Keys(ops))
	slices.Sort(names)
	for _, name := range names {
		fmt.Printf("	%s", name)
		desc := ops[name].desc
		if desc == "" {
			fmt.Println()
		} else {
			fmt.Printf("%s\n", desc)
		}
	}
	fmt.Println("Input:")
	fmt.Println("	-fa <file>   [-ta <file>   ...]: name of file that contains array of tuples")
	fmt.Println("	-ia <inline> [-ia <inline> ...]: inline array of tuples")
	fmt.Println("	-fs <file>   [-ts <file>   ...]: name of file that contains table object: schema and tuples")
	fmt.Println("	-is <inline> [-is <file>   ...]: inline table object: schema and tuples")
	fmt.Println("Options:")
	fmt.Println("	-of <format>: output format: table or json")

	fmt.Println("Note:")
	fmt.Println("	JSON is used as an input format")
	os.Exit(1)
}

type opDesc struct {
	desc string
	fn   func(a, b *table.Table) *table.Table
}

var ops = map[string]opDesc{}

func bind(name, desc string, fn func(a, b *table.Table) *table.Table) {
	ops[name] = opDesc{desc, fn}
}
