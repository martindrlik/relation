package rex_test

import (
	"fmt"

	"github.com/martindrlik/rex"
)

func ExampleUnion() {

	t1 := rex.NewTable("name", "age").Add(rex.T{"name": "John", "age": 42})
	t2 := rex.NewTable("name", "age").Add(rex.T{"name": "Jake"})
	t3 := rex.Union(t1, t2)
	fmt.Println(rex.BoxTable(t3.Schema(), t3.Relations()))
	// Output:
	// ┏━━━━━┯━━━━━━┓
	// ┃ age │ name ┃
	// ┠─────┼──────┨
	// ┃ 42  │ John ┃
	// ┃ *   │ Jake ┃
	// ┗━━━━━┷━━━━━━┛

}
