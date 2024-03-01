package main

import (
	"fmt"

	"github.com/martindrlik/rex"
)

func (state *state) tablesCmd(a args) {
	tables := rex.NewTable("name")
	for name := range state.tables {
		tables.Add(rex.T{"name": name})
	}
	printTable(tables)
}

func (state *state) tableCmd(a args) {
	switch len(a) {
	case 0:
	case 1:
		state.tablePrint(a.first())
	}
}

func printTable(table *rex.Table) {
	fmt.Print(rex.BoxTable(table.Schema(), table.Relations()))
}

func (state *state) tablePrint(name string) {
	table, ok := state.tables[name]
	if ok {
		printTable(table)
	} else {
		fmt.Println("table not found")
	}
}

func (state *state) tableSchema(name string, schema ...string) *rex.Table {
	table := rex.NewTable(schema...)
	state.tables[name] = table
	return table
}
