/*
 */
package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/martindrlik/rex/box"
	"github.com/martindrlik/rex/load"
	"github.com/martindrlik/rex/table"
	"golang.org/x/exp/maps"
)

var (
	format = flag.String("format", "table", "output format: table, json")
)

func main() {
	flag.Usage = usage
	flag.Parse()

	rest := func(n int) []string {
		return flag.Args()[n:]
	}
	switch flag.Arg(0) {
	case "print":
		print(loadTables(rest(1)...))
	case "union":
		union(loadTables(rest(1)...))
	case "naturaljoin":
		naturaljoin(loadTables(rest(1)...))
	case "project":
		project(loadTables(flag.Arg(1)), rest(2)...)
	}
}

func usage() {
	fmt.Println("usage:")
	fmt.Println("	rex print [file ...]")
	fmt.Println("	rex union [file ...]")
	fmt.Println("	rex naturaljoin [file ...]")
	fmt.Println("	rex project file [attribute ...]")
	os.Exit(1)
}

func loadTables(names ...string) []*table.Table {
	tables, err := load.TableFiles(names...)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return tables
}

func print(tables []*table.Table, attributes ...string) {
	json := func(t *table.Table) {
		panic("not implemented")
	}
	for _, t := range tables {
		if len(attributes) == 0 {
			attributes = maps.Keys(t.Schema())
			sort.Strings(attributes)
		}
		if *format == "json" {
			json(t)
		} else {
			fmt.Println(box.Table(attributes, t.Tuples()...))
		}
	}
}

func union(tables []*table.Table) {
	t := tables[0]
	for _, u := range tables[1:] {
		t = t.Union(u)
	}
	print(tableSlice(t))
}

func naturaljoin(tables []*table.Table) {
	t := tables[0]
	for _, u := range tables[1:] {
		t = t.NaturalJoin(u)
	}
	print(tableSlice(t))
}

func project(tables []*table.Table, attributes ...string) {
	for _, t := range tables {
		print(tableSlice(t.Project(attributes...)))
	}
}

func tableSlice(tables ...*table.Table) []*table.Table {
	return tables
}
