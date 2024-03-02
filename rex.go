package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/martindrlik/rex/box"
	"github.com/martindrlik/rex/maps"
	"github.com/martindrlik/rex/require"
	"github.com/martindrlik/rex/table"
	"github.com/martindrlik/rex/tuple"
)

func main() {
	newRex().exec(os.Stdin)
}

type args []string

func (a args) Head() string {
	return a[0]
}

func (a args) Tail() args {
	return a[1:]
}

type cmd struct {
	desc string
	f    func(args)
}

type rex struct {
	ts map[string]*table.Table
	cs map[string]cmd
}

func newRex() *rex {
	rex := &rex{
		ts: map[string]*table.Table{},
	}
	rex.cs = map[string]cmd{
		"?":  {"print help", rex.printHelp},
		"ld": {"load file", rex.load},
		"q":  {"quit", func(args) { os.Exit(0) }},
		".":  {"print tables", rex.printTables},
		"u":  {"union tables", rex.union},
	}
	return rex
}

func (rex *rex) exec(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) == 0 {
			continue
		}
		a := args(fields)
		cmd, ok := rex.cs[a.Head()]
		if !ok {
			fmt.Println("unknown command")
			continue
		}
		cmd.f(a.Tail())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("failed to scan: %v\n", err)
	}
}

func (rex *rex) printHelp(args) {
	x := require.Must(table.NewTable("cmd", "description"))
	order := maps.Keys(rex.cs)
	for _, k := range order {
		c := rex.cs[k]
		x.Append(tuple.Tuple{"cmd": k, "description": c.desc})
	}
	rex.printTable(x)
}

func (rex *rex) printTables(a args) {
	printTableNames := func() {
		x := require.Must(table.NewTable("name"))
		for name := range rex.ts {
			x.Append(tuple.Tuple{"name": name})
		}
		rex.printTable(x)
	}

	if len(a) == 0 {
		printTableNames()
	} else {
		rex.printTableByName(a.Head())
	}
}

func (rex *rex) load(a args) {
	if len(a) == 0 {
		fmt.Println("usage: ld <filename>")
		return
	}
	rex.loadFile(a.Head())
}

func (rex *rex) printTableByName(name string) {
	x, ok := rex.ts[name]
	if ok {
		rex.printTable(x)
	} else {
		fmt.Println("table not found")
	}
}

func (rex *rex) printTable(x *table.Table) {
	fmt.Print(box.Table(x.Schema().Attributes(), x.Relations()))
}

func (rex *rex) union(a args) {
	if len(a) < 3 {
		fmt.Println("usage: u table1 table2 newtable")
		return
	}
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
