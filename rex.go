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
	ts  map[string]*table.Table
	run func(args)
}

func newRex() *rex {
	rex := &rex{}
	rex.ts = map[string]*table.Table{}
	rex.run = rex.makeRunCommand()
	return rex
}

func (rex *rex) exec(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		rex.run(args(strings.Fields(scanner.Text())))
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
	fmt.Print(box.Table(x))
}
