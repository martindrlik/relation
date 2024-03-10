package table_test

import (
	"fmt"

	"github.com/martindrlik/rex/box"
	"github.com/martindrlik/rex/table"
)

func ExampleTable() {
	movies := table.New().Add(
		map[string]any{"title": "The Matrix", "year": 1999},
		map[string]any{"title": "Blade Runner: 2049", "year": 2017, "length": 164},
		map[string]any{"title": "Dune", "year": 2021, "length": 155})

	fmt.Println(box.Table([]string{"title", "year", "length"}, movies.Tuples()...))

	// Output:
	// ┏━━━━━━━━━━━━━━━━━━━━┯━━━━━━┯━━━━━━━━┓
	// ┃ title              │ year │ length ┃
	// ┠────────────────────┼──────┼────────┨
	// ┃ The Matrix         │ 1999 │ ?      ┃
	// ┃ Blade Runner: 2049 │ 2017 │ 164    ┃
	// ┃ Dune               │ 2021 │ 155    ┃
	// ┗━━━━━━━━━━━━━━━━━━━━┷━━━━━━┷━━━━━━━━┛
}

func ExampleTable_Add() {
	movies := table.New().Add(
		map[string]any{"title": "The Matrix", "year": 1999},
		map[string]any{"title": "The Matrix", "year": 1999}) // no duplicate

	fmt.Println(box.Table([]string{"title", "year"}, movies.Tuples()...))

	// Output:
	// ┏━━━━━━━━━━━━┯━━━━━━┓
	// ┃ title      │ year ┃
	// ┠────────────┼──────┨
	// ┃ The Matrix │ 1999 ┃
	// ┗━━━━━━━━━━━━┷━━━━━━┛
}
