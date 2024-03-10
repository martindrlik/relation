package table_test

import (
	"fmt"

	"github.com/martindrlik/rex/box"
	"github.com/martindrlik/rex/table"
)

func ExampleTable_NaturalJoin() {
	movies := table.New().Add(
		T{"title": "The Matrix", "year": 1999},
		T{"title": "Dune", "year": 2021})
	actors := table.New().Add(
		T{"actor": "Keanu Reeves", "title": "The Matrix"},
		T{"actor": "Carrie-Anne Moss", "title": "The Matrix"},
		T{"actor": "Laurence Fishburne", "title": "The Matrix"},
		T{"actor": "Timothée Chalamet", "title": "Dune"},
		T{"actor": "Rebecca Ferguson", "title": "Dune"},
		T{"actor": "Zendaya", "title": "Dune"})

	fmt.Println(box.Table([]string{"title", "year", "actor"}, movies.NaturalJoin(actors).Tuples()...))

	// Output:
	// ┏━━━━━━━━━━━━┯━━━━━━┯━━━━━━━━━━━━━━━━━━━━┓
	// ┃ title      │ year │ actor              ┃
	// ┠────────────┼──────┼────────────────────┨
	// ┃ The Matrix │ 1999 │ Keanu Reeves       ┃
	// ┃ The Matrix │ 1999 │ Carrie-Anne Moss   ┃
	// ┃ The Matrix │ 1999 │ Laurence Fishburne ┃
	// ┃ Dune       │ 2021 │ Timothée Chalamet  ┃
	// ┃ Dune       │ 2021 │ Rebecca Ferguson   ┃
	// ┃ Dune       │ 2021 │ Zendaya            ┃
	// ┗━━━━━━━━━━━━┷━━━━━━┷━━━━━━━━━━━━━━━━━━━━┛
}
