package rex_test

import (
	"fmt"

	"github.com/martindrlik/rex"
)

func ExampleInsertOne() {
	users := rex.R{}
	must(users.InsertOne(rex.String(`{"name": "Jake"}`)))
	must(users.InsertOne(rex.String(`{"age": 35}`)))
	must(users.InsertOne(rex.String(`{"occupation": "developer"}`)))
	must(users.InsertOne(rex.String(`{"age": 35}`))) // duplicate is not inserted
	fmt.Println(rex.Dump(users, 3, 4, 10))
	// Output:
	// age | name | occupation
	// 35  | ✕    | ✕
	// ✕   | Jake | ✕
	// ✕   | ✕    | developer
}
