package rex_test

import (
	"fmt"

	"github.com/martindrlik/rex"
)

func ExampleDump() {
	g := rex.R{}
	must(g.InsertOne(rex.String(`{"greeting": "Hello", "weight": 1.5}`)))
	must(g.InsertOne(rex.String(`{"age": 35}`)))
	fmt.Println(rex.Dump(g, 3, 8, 6, 4))
	// Output:
	// age | greeting | weight
	// 35  | ✕        | ✕
	// ✕   | Hello    | 1.5
}
