package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"

	"github.com/martindrlik/rex/box"
	"github.com/martindrlik/rex/persist"
	"github.com/martindrlik/rex/table"
	"golang.org/x/exp/maps"
)

func main() {
	bind("union", "", func(a, b *table.Table) *table.Table { return a.Union(b) })
	bind("difference", "", func(a, b *table.Table) *table.Table { return a.SetDifference(b) })
	bind("natural-join", "", func(a, b *table.Table) *table.Table { return a.NaturalJoin(b) })
	exec(parse(os.Args[1:]))
}

func exec(op string, tables []*table.Table, attributes []string) {
	func(fn func([]*table.Table) []*table.Table) {
		for _, t := range fn(tables) {
			project(t, attributes)
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

func project(table *table.Table, attributes []string) {
	if len(attributes) == 0 {
		attributes = maps.Keys(table.Schema())
		slices.Sort(attributes)
	}
	fmt.Println(box.Table(attributes, table.Project(attributes...).Tuples()...))
}

func parse(args []string) (string, []*table.Table, []string) {
	if len(args) < 2 {
		usage(errors.New("missing arguments"))
	}
	fs := flag.NewFlagSet("", flag.ExitOnError)
	var (
		names = stringsFlag{}
		jsons = stringsFlag{}

		namesStrict = stringsFlag{}
		jsonsStrict = stringsFlag{}
	)
	fs.Var(&names, "t", "table file")
	fs.Var(&jsons, "j", "inline json")
	fs.Var(&namesStrict, "ts", "table file")
	fs.Var(&jsonsStrict, "js", "inline json")

	op := args[0]
	_, ok := ops[op]
	if !ok {
		usage(fmt.Errorf("unknown op: %s", op))
	}

	fs.Parse(args[1:])
	if len(names) == 0 && len(namesStrict) == 0 && len(jsons) == 0 && len(jsonsStrict) == 0 {
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
	loadFiles := func(names []string, fn func(io.Reader) (*table.Table, error)) {
		for _, name := range names {
			if err := loadFile(name, fn); err != nil {
				usage(fmt.Errorf("loading table: %w", err))
			}
		}
	}
	loadJsons := func(jsons []string, fn func(io.Reader) (*table.Table, error)) {
		for _, j := range jsons {
			t, err := fn(strings.NewReader(j))
			if err != nil {
				usage(fmt.Errorf("loading json: %w", err))
			}
			tables = append(tables, t)
		}
	}

	loadFiles(names, persist.Load)
	loadFiles(namesStrict, persist.LoadStrict)

	loadJsons(jsons, persist.Load)
	loadJsons(jsonsStrict, persist.LoadStrict)

	return op, tables, fs.Args()
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
	fmt.Println("	rex <command> -t filename [-t filename ...] [attribute ...]")
	fmt.Println("	rex <command> -j inlinejson [-j inlinejson ...] [attribute ...]")
	fmt.Println("Commands:")
	names := maps.Keys(ops)
	slices.Sort(names)
	for _, name := range names {
		fmt.Printf("	%s: %s\n", name, ops[name].desc)
	}
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
