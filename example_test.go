package rex_test

import (
	"fmt"

	"github.com/martindrlik/rex"
)

func ExampleDump() {
	users := rex.R{}
	must(users.Insert(rex.String(`{"name": "Jake"}`)))
	must(users.Insert(rex.String(`{"age": 35}`)))
	must(users.Insert(rex.String(`{"occupation": "developer"}`)))
	must(users.Insert(rex.String(`{"age": 35}`))) // duplicate is not inserted
	fmt.Println(rex.Dump(users))
	// Output:
	// age | name | occupation
	// 35  | ✕    | ✕
	// ✕   | Jake | ✕
	// ✕   | ✕    | developer
}

func ExampleDumpPad() {
	users := rex.R{}
	must(users.Insert(rex.String(`{"name": "Gwendolyn"}`)))
	must(users.Insert(rex.String(`{"occupation": "developer"}`)))
	fmt.Println(rex.Dump(users, rex.Pad("name", 9)))
	// Output:
	// name      | occupation
	// Gwendolyn | ✕
	// ✕         | developer
}

func ExampleInnerRelation() {
	users := rex.R{}
	must(users.Insert(rex.String(`{"name": "Jake", "address": {"city": "New York", "street": "Broadway"}}}`)))
	fmt.Println(rex.Dump(users, rex.Pad("city", 8)))
	// Output:
	// address | name
	// *r1     | Jake
	// -- r1:
	// city     | street
	// New York | Broadway
}

func ExampleCopy() {
	a := rex.R{}
	must(a.Insert(rex.String(`{"one": 1, "two": 2}`)))
	must(a.Insert(rex.String(`{"three": 3}`)))
	b := a.Copy()
	must(b.Insert(rex.String(`{"four": 4}`)))
	fmt.Println(rex.Dump(a))
	fmt.Println("--")
	fmt.Println(rex.Dump(b))
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

func ExampleUnion() {
	r := rex.R{}
	s := rex.R{}
	must(r.Insert(rex.String(`{"name": "Jake", "age": 24}`)))
	must(r.Insert(rex.String(`{"city": "Olomouc"}`)))
	must(s.Insert(rex.String(`{"name": "Aya", "age": 30}`)))
	must(s.Insert(rex.String(`{"city": "Prague"}`)))
	t := r.Union(s)
	fmt.Println(rex.Dump(t, rex.Pad("city", 7)))
	// Output:
	// age | city    | name
	// 24  | ✕       | Jake
	// 30  | ✕       | Aya
	// ✕   | Olomouc | ✕
	// ✕   | Prague  | ✕
}
