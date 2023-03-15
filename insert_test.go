package rex_test

import (
	"fmt"
	"strings"

	"github.com/martindrlik/rex"
)

func ExampleInsertOne() {
	users := rex.R{}
	users.InsertOne(strings.NewReader(`{"name": "Jake"}`))
	users.InsertOne(strings.NewReader(`{"age": 35}`))
	users.InsertOne(strings.NewReader(`{"occupation": "developer"}`))
	users.InsertOne(strings.NewReader(`{"age": 35}`)) // duplicate is not inserted
	fmt.Println(rex.Dump(users, 3, 4, 10))
	// Output:
	// age | name | occupation
	// 35  | ✕    | ✕
	// ✕   | Jake | ✕
	// ✕   | ✕    | developer
}
