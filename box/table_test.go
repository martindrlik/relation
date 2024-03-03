package box_test

import (
	"fmt"

	"github.com/martindrlik/rex/box"
	"github.com/martindrlik/rex/require"
	"github.com/martindrlik/rex/table"
	"github.com/martindrlik/rex/tuple"
)

func ExampleTable() {

	t := require.NoError(table.New("title", "year"))
	require.NilError(t.Append(tuple.T{"title": "Adventure Time"}))
	require.NilError(t.Append(tuple.T{"title": "What We Do in the Shadows", "year": 2019}))
	fmt.Println(box.Table(t))
	fmt.Println(box.Table(t.Project("title")))
	empty := require.NoError(table.New("title", "year"))
	fmt.Println(box.Table(empty))

	// Output:
	// ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━┯━━━━━━┓
	// ┃ title                     │ year ┃
	// ┠───────────────────────────┼──────┨
	// ┃ Adventure Time            │ *    ┃
	// ┃ What We Do in the Shadows │ 2019 ┃
	// ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━┛
	//
	// ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
	// ┃ title                     ┃
	// ┠───────────────────────────┨
	// ┃ Adventure Time            ┃
	// ┃ What We Do in the Shadows ┃
	// ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
	//
	// ┏━━━━━━━┯━━━━━━┓
	// ┃ title │ year ┃
	// ┗━━━━━━━┷━━━━━━┛

}
