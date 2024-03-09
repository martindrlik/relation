package table_test

import (
	"fmt"

	"github.com/martindrlik/rex/box"
	"github.com/martindrlik/rex/table"
)

func ExampleTable_Restrict() {
	movies := table.New().Add(
		map[string]any{"title": "Die Hard", "year": 1988},
		map[string]any{"title": "The Matrix", "year": 1999},
		map[string]any{"title": "Guardians of the Galaxy", "year": 2014},
		map[string]any{"title": "Blade Runner: 2049", "year": 2017},
		map[string]any{"title": "Dune", "year": 2021})
	year := func(f func(int) bool) func(tuple map[string]any) bool {
		return func(tuple map[string]any) bool {
			return f(tuple["year"].(int))
		}
	}

	fmt.Println(box.Table([]string{"title", "year"}, movies.Restrict(year(func(x int) bool { return x < 2000 })).Tuples()...))
	fmt.Println(box.Table([]string{"title", "year"}, movies.Restrict(year(func(x int) bool { return x > 2000 })).Tuples()...))

	// Output:
	// ┏━━━━━━━━━━━━┯━━━━━━┓
	// ┃ title      │ year ┃
	// ┠────────────┼──────┨
	// ┃ Die Hard   │ 1988 ┃
	// ┃ The Matrix │ 1999 ┃
	// ┗━━━━━━━━━━━━┷━━━━━━┛
	//
	// ┏━━━━━━━━━━━━━━━━━━━━━━━━━┯━━━━━━┓
	// ┃ title                   │ year ┃
	// ┠─────────────────────────┼──────┨
	// ┃ Guardians of the Galaxy │ 2014 ┃
	// ┃ Blade Runner: 2049      │ 2017 ┃
	// ┃ Dune                    │ 2021 ┃
	// ┗━━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━┛
}
