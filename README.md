# rex

Experimental relational NoSQL database. It is my playground for ideas and API will change over time. There is a lot more to do before it can be even considered interesting.

## Example

``` go
package example

import (
	"fmt"

	"github.com/martindrlik/rex/box"
	"github.com/martindrlik/rex/table"
)

type T = map[string]any

func ExampleTable() {
	movies := table.New().Add(
		T{"title": "The Matrix", "year": 1999},
		T{"title": "Dune", "year": 2021, "length": 155},
		T{"title": "Blade Runner: 2049", "year": 2017, "length": 164})

	fmt.Println(box.Table([]string{"title", "year", "length"}, movies.Tuples()...))

	// Output:
	// ┏━━━━━━━━━━━━━━━━━━━━┯━━━━━━┯━━━━━━━━┓
	// ┃ title              │ year │ length ┃
	// ┠────────────────────┼──────┼────────┨
	// ┃ The Matrix         │ 1999 │ ?      ┃
	// ┃ Dune               │ 2021 │ 155    ┃
	// ┃ Blade Runner: 2049 │ 2017 │ 164    ┃
	// ┗━━━━━━━━━━━━━━━━━━━━┷━━━━━━┷━━━━━━━━┛
}

```