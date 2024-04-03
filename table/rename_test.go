package table_test

import (
	"fmt"

	"github.com/martindrlik/rex/box"
	"github.com/martindrlik/rex/table"
)

func ExampleTable_Rename() {
	movie := func(title string, year int) map[string]any {
		return map[string]any{"title": title, "year": year}
	}
	movies := table.New("title", "year").Add(
		movie("The Matrix", 1999),
		movie("Dune", 2021))
	movies = movies.Rename(map[string]string{
		"title": "movie_title",
		"year":  "release_year"})
	fmt.Println(box.Table(
		[]string{"movie_title", "release_year"},
		movies.Tuples()...))

	// Output:
	// ┏━━━━━━━━━━━━━┯━━━━━━━━━━━━━━┓
	// ┃ movie_title │ release_year ┃
	// ┠─────────────┼──────────────┨
	// ┃ The Matrix  │ 1999         ┃
	// ┃ Dune        │ 2021         ┃
	// ┗━━━━━━━━━━━━━┷━━━━━━━━━━━━━━┛
}
