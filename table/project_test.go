package table_test

import (
	"fmt"

	"github.com/martindrlik/rex/box"
	"github.com/martindrlik/rex/table"
)

func ExampleTable_Project() {
	movies := table.New().Add(
		map[string]any{"title": "The Matrix", "year": 1999},
		map[string]any{"title": "Dune", "year": 2021, "length": 155},
		map[string]any{"title": "Blade Runner: 2049", "year": 2017, "length": 164})

	fmt.Println(box.Table([]string{"title"}, movies.Project("title").Tuples()...))

	// Output:
	// ┏━━━━━━━━━━━━━━━━━━━━┓
	// ┃ title              ┃
	// ┠────────────────────┨
	// ┃ The Matrix         ┃
	// ┃ Dune               ┃
	// ┃ Blade Runner: 2049 ┃
	// ┗━━━━━━━━━━━━━━━━━━━━┛
}
