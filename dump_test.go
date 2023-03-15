package rex_test

import (
	"fmt"
	"strings"

	"github.com/martindrlik/rex"
)

func ExampleDump() {
	g := rex.R{}
	g.InsertOne(strings.NewReader(`{"greeting": "Hello", "weight": 1.5}`))
	g.InsertOne(strings.NewReader(`{"age": 35}`))
	fmt.Println(rex.Dump(g, 3, 8, 6, 4))
	// Output:
	// age | greeting | weight
	// 35  | ✕        | ✕
	// ✕   | Hello    | 1.5
}
