package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/martindrlik/rex/box"
	"github.com/martindrlik/rex/table"
)

type rex struct {
	ts map[string]*table.Table
	cs map[string]command
}

func newRex() *rex {
	rex := &rex{
		ts: map[string]*table.Table{},
	}
	rex.cs = map[string]command{}
	bind := func(name string, f func(name string) command) {
		rex.cs[name] = f(name)
	}
	bind("?", rex.makeHelpCommand)
	bind(".", rex.makePrintTableCommand)
	bind("l", rex.makeLoadCommand)
	bind("u", rex.makeUnionCommand)
	bind("x", rex.makeExitCommand)
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
		cmd.action(a.Tail())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("failed to scan: %v\n", err)
	}
}

func (rex *rex) printTableByName(name string) {
	x, ok := rex.ts[name]
	if ok {
		rex.printTable(x)
	} else {
		fmt.Printf("table %q not found\n", name)
	}
}

func (rex *rex) printTable(x *table.Table) {
	fmt.Print(box.Table(x.Schema().Attributes(), x.Relations()))
}
