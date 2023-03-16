package rex_test

import (
	"fmt"

	"github.com/martindrlik/rex"
)

func ExampleUnion() {
	r := rex.R{}
	s := rex.R{}
	must(r.InsertOne(rex.String(`{"name": "Jake", "age": 24}`)))
	must(r.InsertOne(rex.String(`{"city": "Olomouc"}`)))
	must(s.InsertOne(rex.String(`{"name": "Aya", "age": 30}`)))
	must(s.InsertOne(rex.String(`{"city": "Prague"}`)))
	t := r.Union(s)
	fmt.Println(rex.Dump(t, 3, 7, 4))
	// Output:
	// age | city    | name
	// 24  | ✕       | Jake
	// 30  | ✕       | Aya
	// ✕   | Olomouc | ✕
	// ✕   | Prague  | ✕
}
