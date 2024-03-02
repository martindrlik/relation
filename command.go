package main

import (
	"fmt"
	"os"
	"path/filepath"
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

func (rex *rex) makeRunCommand() func(args) {
	m := map[string]command{}
	d := command{}
	bind := func(name string, f func(name string) command) {
		m[name] = f(name)
	}
	bindHelp := func(name string) {
		d = rex.makeHelpCommand(name, m)
		m[name] = d
	}

	bindHelp("?")
	bind(".", rex.makePrintTableCommand)
	bind("l", rex.makeLoadCommand)
	bind("u", rex.makeUnionCommand)
	bind("x", rex.makeExitCommand)

	return func(a args) {
		if len(a) == 0 {
			d.action(a)
			return
		}
		c, ok := m[a.Head()]
		if !ok {
			d.action(args{})
			return
		}
		c.action(a.Tail())
	}
}

func (rex *rex) makeHelpCommand(name string, m map[string]command) command {
	helpUsage := func() {
		fmt.Printf("usage: %s\n", name)
		fmt.Println("  print help")
		fmt.Printf("usage: %s command\n", name)
		fmt.Println("  print help for command")
	}
	help := func(a args) {
		if len(a) == 1 {
			c, ok := m[a.Head()]
			if ok {
				c.usage()
			}
			return
		}
		x := require.Must(table.New("name", "description"))
		sorted := maps.Keys(m)
		sort.Strings(sorted)
		for _, k := range sorted {
			c := m[k]
			x.Append(tuple.T{"name": k, "description": c.desc})
		}
		rex.printTable(x)
	}
	return command{"help", helpUsage, help}
}

func (rex *rex) makePrintTableCommand(name string) command {
	tablesUsage := func() {
		fmt.Printf("usage: %s\n", name)
		fmt.Println("  print table names")
		fmt.Printf("usage: %s table\n", name)
		fmt.Println("  print table")
		fmt.Printf("usage: %s table1 table2 ...\n", name)
		fmt.Println("  print table1 then table2 then all other listed tables")
	}
	tables := func(a args) {
		if len(a) == 0 {
			x := require.Must(table.New("name"))
			for name := range rex.ts {
				x.Append(tuple.T{"name": name})
			}
			rex.printTable(x)
			return
		}
		for _, name := range a {
			rex.printTableByName(name)
		}
	}
	return command{"print", tablesUsage, tables}
}

func (rex *rex) makeLoadCommand(name string) command {
	loadUsage := func() {
		fmt.Printf("usage: %s filename\n", name)
		fmt.Println("  load table from file filename")
		fmt.Printf("usage: %s filename1 filename2 ...\n", name)
		fmt.Println("  load tables from files filename1 filename2 and from all other listed file names")
	}
	load := func(a args) {
		for _, name := range a {
			names, err := filepath.Glob(name)
			if err != nil {
				fmt.Printf("failed to load %q: %v\n", name, err)
				return
			}
			for _, name := range names {
				rex.loadFile(name)
			}
		}
	}
	return command{"file loader", loadUsage, load}
}

func (rex *rex) makeUnionCommand(name string) command {
	unionUsage := func() {
		fmt.Printf("usage: %s newtable table1 table2\n", name)
		fmt.Println("  create newtable as the union of table1 and table2")
	}
	union := func(a args) {
		if len(a) < 3 {
			unionUsage()
			return
		}
		findTable := func(name string) *table.Table {
			t, ok := rex.ts[name]
			if !ok {
				fmt.Printf("table %q not found\n", name)
			}
			return t
		}
		t := findTable(a[1])
		v := findTable(a[2])
		if t == nil || v == nil {
			return
		}
		x, err := t.Union(v)
		if err != nil {
			fmt.Printf("union failed: %v\n", err)
		}
		rex.ts[a[0]] = x
	}
	return command{"table union", unionUsage, union}
}

func (rex *rex) makeExitCommand(name string) command {
	exitUsage := func() {
		fmt.Printf("usage: %s\n", name)
		fmt.Println("  exit the program")
	}
	exit := func(a args) {
		os.Exit(0)
	}
	return command{"exit", exitUsage, exit}
}
