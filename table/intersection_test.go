package table_test

import (
	"fmt"

	"github.com/martindrlik/rex/box"
	"github.com/martindrlik/rex/table"
)

func ExampleTable_Intersection() {
	u := table.New().Add(
		T{"title": "Dune"},
		T{"title": "Dune: Part Two", "year": 2024})
	v := table.New().Add(
		T{"title": "Dune"})

	fmt.Println(box.Table([]string{"title"}, u.Intersection(v).Tuples()...))

	// Output:
	// ┏━━━━━━━┓
	// ┃ title ┃
	// ┠───────┨
	// ┃ Dune  ┃
	// ┗━━━━━━━┛
}
