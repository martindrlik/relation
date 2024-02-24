package rex_test

import (
	"fmt"

	"github.com/martindrlik/rex"
)

func ExampleBoxTable() {

	t1 := rex.NewTable("title", "year").
		Add(rex.T{"title": "Adventure Time"}).
		Add(rex.T{"title": "What We Do in the Shadows", "year": 2019})

	fmt.Println(rex.BoxTable(t1.SchemaInOrder(), t1.Relations()))

	t2 := t1.Pick("title")
	fmt.Println(rex.BoxTable(t2.SchemaInOrder(), t2.Relations()))

	empty := rex.NewTable("title", "year")
	fmt.Println(rex.BoxTable(empty.SchemaInOrder(), empty.Relations()))

	// Output:
	// ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━┯━━━━━━┓
	// ┃ title                     │ year ┃
	// ┠───────────────────────────┼──────┨
	// ┃ Adventure Time            │ *    ┃
	// ┃ What We Do in the Shadows │ 2019 ┃
	// ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━┛
	//
	// ┏━━━━━━━━━━━━━━━━┓
	// ┃ title          ┃
	// ┠────────────────┨
	// ┃ Adventure Time ┃
	// ┗━━━━━━━━━━━━━━━━┛
	//
	// ┏━━━━━━━┯━━━━━━┓
	// ┃ title │ year ┃
	// ┗━━━━━━━┷━━━━━━┛

}
