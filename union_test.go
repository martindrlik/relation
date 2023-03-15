package rex_test

import (
	"fmt"
	"strings"

	"github.com/martindrlik/rex"
)

func ExampleUnion() {
	r := rex.R{}
	s := rex.R{}
	r.InsertOne(strings.NewReader(`{"name": "Jake", "age": 24}`))
	r.InsertOne(strings.NewReader(`{"city": "Olomouc"}`))
	s.InsertOne(strings.NewReader(`{"name": "Aya", "age": 30}`))
	s.InsertOne(strings.NewReader(`{"city": "Prague"}`))
	t := r.Union(s)
	fmt.Println(rex.Dump(t, 3, 7, 4))
	// Output:
	// age | city    | name
	// 24  | ✕       | Jake
	// 30  | ✕       | Aya
	// ✕   | Olomouc | ✕
	// ✕   | Prague  | ✕
}
