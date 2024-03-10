package table_test

import (
	"fmt"

	"github.com/martindrlik/rex/box"
	"github.com/martindrlik/rex/table"
)

func ExampleTable_SetDifference() {
	available := table.New().Add(
		map[string]any{"title": "Dune"},
		map[string]any{"title": "Dune: Part Two", "year": 2024})
	seen := table.New().Add(
		map[string]any{"title": "Dune"})

	fmt.Println(box.Table([]string{"title"}, available.SetDifference(seen).Tuples()...))

	// Output:
	// ┏━━━━━━━━━━━━━━━━┓
	// ┃ title          ┃
	// ┠────────────────┨
	// ┃ Dune           ┃
	// ┃ Dune: Part Two ┃
	// ┗━━━━━━━━━━━━━━━━┛
}
