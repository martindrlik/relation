package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/martindrlik/rex"
)

type state struct {
	tables map[string]*rex.Table
	domain map[string][]any
	cmds   map[string]func(args args)
}

func newState() *state {
	state := &state{
		tables: map[string]*rex.Table{},
		domain: map[string][]any{},
	}
	state.initTables()
	state.cmds = map[string]func(args args){
		"?":  state.help,
		"ld": state.loadCmd,
		"q":  func(args args) { os.Exit(0) },
		"t":  state.tableCmd,
		"ts": state.tablesCmd,
		"u":  state.unionCmd,
	}
	return state
}

func (state *state) help(a args) {
	printTable(state.tables["help"])
}

func (state *state) exec(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) == 0 {
			continue
		}
		a := args(fields)
		cmd, ok := state.cmds[a.first()]
		if !ok {
			fmt.Println("unknown command")
			continue
		}
		cmd(a.rest())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("failed to scan: %v\n", err)
	}
}

func (state *state) initTables() {
	table := rex.NewTable("cmd", "argument", "description")
	table.Add(rex.T{"cmd": "ts", "description": "print all table names"})
	table.Add(rex.T{"cmd": "t", "argument": "table name", "description": "print table"})
	table.Add(rex.T{"cmd": "ld", "argument": "file name", "description": "load file"})
	table.Add(rex.T{"cmd": "u", "argument": "table1 table2 newtable", "description": "union table1 and table2 into newtable"})
	state.tables["help"] = table
}
