package box_test

import (
	"fmt"

	"github.com/martindrlik/rex/box"
)

func ExampleTable() {
	fmt.Println(box.Table(
		[]string{"title", "year"},
		map[string]any{"title": "Adventure Time", "year": 2010},
		map[string]any{"title": "What We Do in the Shadows", "year": 2019},
		map[string]any{"title": "The Last of Us"}))

	fmt.Println(box.Table([]string{"empty", "table"}))

	// Output:
	// ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━┯━━━━━━┓
	// ┃ title                     │ year ┃
	// ┠───────────────────────────┼──────┨
	// ┃ Adventure Time            │ 2010 ┃
	// ┃ What We Do in the Shadows │ 2019 ┃
	// ┃ The Last of Us            │ ?    ┃
	// ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━┛
	//
	// ┏━━━━━━━┯━━━━━━━┓
	// ┃ empty │ table ┃
	// ┗━━━━━━━┷━━━━━━━┛
}
