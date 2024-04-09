package table_test

import (
	"fmt"
	"testing"

	"github.com/martindrlik/rex/box"
	"github.com/martindrlik/rex/table"
)

type T = map[string]any

func ExampleTable() {
	movies := table.New("title", "year", "length").Add(
		movie("The Matrix", 1999),
		withLength(movie("Blade Runner: 2049", 2017), 164),
		withLength(movie("Dune: Part One", 2021), 155))

	fmt.Println(box.Table(movies.SchemaOrder(), movies.Tuples()...))

	// Output:
	// ┏━━━━━━━━━━━━━━━━━━━━┯━━━━━━┯━━━━━━━━┓
	// ┃ title              │ year │ length ┃
	// ┠────────────────────┼──────┼────────┨
	// ┃ The Matrix         │ 1999 │ ?      ┃
	// ┃ Blade Runner: 2049 │ 2017 │ 164    ┃
	// ┃ Dune: Part One     │ 2021 │ 155    ┃
	// ┗━━━━━━━━━━━━━━━━━━━━┷━━━━━━┷━━━━━━━━┛
}

func ExampleTable_Add() {
	movies := table.New().Add(
		movie("The Matrix", 1999),
		movie("The Matrix", 1999)) // duplicate

	fmt.Println(box.Table([]string{"title", "year"}, movies.Tuples()...))

	// Output:
	// ┏━━━━━━━━━━━━┯━━━━━━┓
	// ┃ title      │ year ┃
	// ┠────────────┼──────┨
	// ┃ The Matrix │ 1999 ┃
	// ┗━━━━━━━━━━━━┷━━━━━━┛
}

func TestContains(t *testing.T) {
	matrixMovie := movie("The Matrix", 1999)
	movies := table.New().Add(matrixMovie)
	moviesBox := box.Table([]string{"title", "year"}, movies.Tuples()...)
	if !movies.Tuples().Contains(matrixMovie) {
		t.Errorf(
			"\nexpected\n%v\nto contain\n%v",
			moviesBox,
			matrixMovie)
	}

	matrixWithLength := withLength(matrixMovie, 136)
	if movies.Tuples().Contains(matrixWithLength) {
		t.Errorf(
			"\nexpected\n%v\nnot to contain\n%v",
			moviesBox,
			matrixWithLength)
	}
}

func movie(title string, year int) map[string]any {
	return map[string]any{"title": title, "year": year}
}

func withLength(tuple map[string]any, length int) map[string]any {
	x := map[string]any{}
	for k, v := range tuple {
		x[k] = v
	}
	x["length"] = length
	return x
}
