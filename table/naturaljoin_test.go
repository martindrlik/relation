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
	casts := table.New().Add(
		T{"cast": "Keanu Reeves", "title": "The Matrix"},
		T{"cast": "Carrie-Anne Moss", "title": "The Matrix"},
		T{"cast": "Laurence Fishburne", "title": "The Matrix"},
		T{"cast": "Timothée Chalamet", "title": "Dune"},
		T{"cast": "Rebecca Ferguson", "title": "Dune"},
		T{"cast": "Zendaya", "title": "Dune"})

	fmt.Println(box.Table([]string{"title", "year", "cast"}, movies.NaturalJoin(casts).Tuples()...))

	numbers := table.New().Add(
		T{"number": 1},
		T{"number": 2},
		T{"number": 3})
	letters := table.New().Add(
		T{"letter": "a"},
		T{"letter": "b"},
		T{"letter": "c"})

	fmt.Println("Cartesian product:")
	fmt.Println(box.Table([]string{"number", "letter"}, numbers.NaturalJoin(letters).Tuples()...))

	// Output:
	// ┏━━━━━━━━━━━━┯━━━━━━┯━━━━━━━━━━━━━━━━━━━━┓
	// ┃ title      │ year │ cast               ┃
	// ┠────────────┼──────┼────────────────────┨
	// ┃ The Matrix │ 1999 │ Keanu Reeves       ┃
	// ┃ The Matrix │ 1999 │ Carrie-Anne Moss   ┃
	// ┃ The Matrix │ 1999 │ Laurence Fishburne ┃
	// ┃ Dune       │ 2021 │ Timothée Chalamet  ┃
	// ┃ Dune       │ 2021 │ Rebecca Ferguson   ┃
	// ┃ Dune       │ 2021 │ Zendaya            ┃
	// ┗━━━━━━━━━━━━┷━━━━━━┷━━━━━━━━━━━━━━━━━━━━┛
	//
	// Cartesian product:
	// ┏━━━━━━━━┯━━━━━━━━┓
	// ┃ number │ letter ┃
	// ┠────────┼────────┨
	// ┃ 1      │ a      ┃
	// ┃ 1      │ b      ┃
	// ┃ 1      │ c      ┃
	// ┃ 2      │ a      ┃
	// ┃ 2      │ b      ┃
	// ┃ 2      │ c      ┃
	// ┃ 3      │ a      ┃
	// ┃ 3      │ b      ┃
	// ┃ 3      │ c      ┃
	// ┗━━━━━━━━┷━━━━━━━━┛
}
