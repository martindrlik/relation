package box_test

import (
	"fmt"

	"github.com/martindrlik/rex/box"
	"github.com/martindrlik/rex/table"
	"github.com/martindrlik/rex/tuple"
)

func ExampleTable() {
	require := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	t := table.NewTable("title", "year")
	require(t.Append(tuple.Tuple{"title": "Adventure Time"}))
	require(t.Append(tuple.Tuple{"title": "What We Do in the Shadows", "year": 2019}))

	fmt.Println(box.Table(t.Schema().Attributes(), t.Relations()))

	v := t.Project("title")
	fmt.Println(box.Table(v.Schema().Attributes(), v.Relations()))

	empty := table.NewTable("title", "year")
	fmt.Println(box.Table(empty.Schema().Attributes(), empty.Relations()))

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
