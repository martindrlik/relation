package table_test

import (
	"fmt"

	"github.com/martindrlik/rex/box"
	"github.com/martindrlik/rex/table"
)

func ExampleTable_Difference() {
	available := table.New().Add(
		T{"title": "Dune"},
		T{"title": "Dune: Part Two", "year": 2024})
	seen := table.New().Add(
		T{"title": "Dune"})

	fmt.Println(box.Table([]string{"title"}, available.Difference(seen).Tuples()...))

	// Output:
	// ┏━━━━━━━━━━━━━━━━┓
	// ┃ title          ┃
	// ┠────────────────┨
	// ┃ Dune           ┃
	// ┃ Dune: Part Two ┃
	// ┗━━━━━━━━━━━━━━━━┛
}
