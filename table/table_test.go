package table_test

import (
	"fmt"
	"testing"

	"github.com/martindrlik/rex/box"
	"github.com/martindrlik/rex/table"
)

type T = map[string]any

func ExampleTable() {
	movies := table.New().Add(
		T{"title": "The Matrix", "year": 1999},
		T{"title": "Blade Runner: 2049", "year": 2017, "length": 164},
		T{"title": "Dune", "year": 2021, "length": 155})

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
		T{"title": "The Matrix", "year": 1999},
		T{"title": "The Matrix", "year": 1999}) // no duplicate

	fmt.Println(box.Table([]string{"title", "year"}, movies.Tuples()...))

	// Output:
	// ┏━━━━━━━━━━━━┯━━━━━━┓
	// ┃ title      │ year ┃
	// ┠────────────┼──────┨
	// ┃ The Matrix │ 1999 ┃
	// ┗━━━━━━━━━━━━┷━━━━━━┛
}

func TestContains(t *testing.T) {
	movies := table.New().Add(T{"title": "The Matrix", "year": 1999})
	moviesBox := box.Table([]string{"title", "year"}, movies.Tuples()...)
	if !movies.Tuples().Contains(T{"title": "The Matrix", "year": 1999}) {
		t.Errorf(
			"\nexpected\n%v\nto contain\n%v",
			moviesBox,
			T{"title": "The Matrix", "year": 1999})
	}

	matrixWithLength := T{"title": "The Matrix", "year": 1999, "length": 136}
	if movies.Tuples().Contains(matrixWithLength) {
		t.Errorf(
			"\nexpected\n%v\nnot to contain\n%v",
			moviesBox,
			matrixWithLength)
	}
}
