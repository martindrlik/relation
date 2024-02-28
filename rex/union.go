package main

import (
	"fmt"

	"github.com/martindrlik/rex"
)

func (state *state) unionCmd(a args) {
	if len(a) < 3 {
		fmt.Println("usage: u table1 table2 newtable")
		return
	}
	findTable := func(name string) *rex.Table {
		t, ok := state.tables[name]
		if !ok {
			fmt.Printf("table %q not found\n", name)
		}
		return t
	}
	t1 := findTable(a[0])
	t2 := findTable(a[1])
	if t1 == nil || t2 == nil {
		return
	}
	u, err := rex.Union(t1, t2)
	if err != nil {
		fmt.Printf("union failed: %v\n", err)
	}
	state.tables[a[2]] = u
}
