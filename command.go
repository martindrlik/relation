package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/martindrlik/rex/maps"
	"github.com/martindrlik/rex/require"
	"github.com/martindrlik/rex/table"
	"github.com/martindrlik/rex/tuple"
)

type command struct {
	desc   string
	usage  func()
	action func(args)
}

func (rex *rex) makeHelpCommand(name string) command {
	helpUsage := func() {
		fmt.Printf("usage: %s command\n", name)
		fmt.Println("  prints command's usage")
		fmt.Println("  note with no argument prints help for all commands")
	}
	help := func(a args) {
		if len(a) == 1 {
			c, ok := rex.cs[a.Head()]
			if ok {
				c.usage()
			}
			return
		}
		x := require.Must(table.NewTable("name", "description"))
		sorted := maps.Keys(rex.cs)
		sort.Strings(sorted)
		for _, k := range sorted {
			c := rex.cs[k]
			x.Append(tuple.Tuple{"name": k, "description": c.desc})
		}
		rex.printTable(x)
	}
	return command{"prints help", helpUsage, help}
}

func (rex *rex) makePrintTableCommand(name string) command {
	tablesUsage := func() {
		fmt.Printf("usage: %s tablename1 [tablename2 ...]\n", name)
		fmt.Println("  prints table or tables")
		fmt.Println("  note with no argument prints all table names")
	}
	tables := func(a args) {
		if len(a) == 0 {
			x := require.Must(table.NewTable("name"))
			for name := range rex.ts {
				x.Append(tuple.Tuple{"name": name})
			}
			rex.printTable(x)
			return
		}
		for _, name := range a {
			rex.printTableByName(name)
		}
	}
	return command{"prints tables", tablesUsage, tables}
}

func (rex *rex) makeLoadCommand(name string) command {
	loadUsage := func() {
		fmt.Printf("usage: %s filename1 [filename2 ...]\n", name)
		fmt.Println("  loads tables from files")
	}
	load := func(a args) {
		for _, name := range a {
			rex.loadFile(name)
		}
	}
	return command{"loads tables from files", loadUsage, load}
}

func (rex *rex) makeUnionCommand(name string) command {
	unionUsage := func() {
		fmt.Printf("usage: %s newtablename tablename1 tablename2 [tablename3 ...]\n", name)
		fmt.Println("  union tables")
	}
	union := func(a args) {
		findTable := func(name string) *table.Table {
			t, ok := rex.ts[name]
			if !ok {
				fmt.Printf("table %q not found\n", name)
			}
			return t
		}
		t := findTable(a[0])
		v := findTable(a[1])
		if t == nil || v == nil {
			return
		}
		x, err := t.Union(v)
		if err != nil {
			fmt.Printf("union failed: %v\n", err)
		}
		rex.ts[a[2]] = x
	}
	return command{"union tables", unionUsage, union}
}

func (rex *rex) makeExitCommand(name string) command {
	exitUsage := func() {
		fmt.Printf("usage: %s\n", name)
		fmt.Println("  exits the program")
	}
	exit := func(a args) {
		os.Exit(0)
	}
	return command{"exits the program", exitUsage, exit}
}
