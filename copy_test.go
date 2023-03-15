package rex_test

import (
	"fmt"
	"strings"

	"github.com/martindrlik/rex"
)

func ExampleCopy() {
	a := rex.R{}
	a.InsertOne(strings.NewReader(`{"one": 1, "two": 2}`))
	a.InsertOne(strings.NewReader(`{"three": 3}`))
	b := a.Copy()
	b.InsertOne(strings.NewReader(`{"four": 4}`))
	fmt.Println(rex.Dump(a, 3, 5, 3))
	fmt.Println("--")
	fmt.Println(rex.Dump(b, 4, 3, 5, 3))
	// Output:
	// one | three | two
	// 1   | ✕     | 2
	// ✕   | 3     | ✕
	// --
	// four | one | three | two
	// 4    | ✕   | ✕     | ✕
	// ✕    | 1   | ✕     | 2
	// ✕    | ✕   | 3     | ✕
}
