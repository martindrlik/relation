package rex_test

import (
	"os"
	"strings"

	"github.com/martindrlik/rex"
)

func ExampleNaturalJoin() {
	names := rex.NewRelation().InsertManyJson(strings.NewReader(`[
		{"id": 1, "name": "Lee"},
		{"id": 2, "name": "Jake"},
		{"id": 3, "name": "Kristen"}
	]`))
	years := rex.NewRelation().InsertManyJson(strings.NewReader(`[
		{"id": 1, "bornYear": 1979},
		{"id": 2, "bornYear": 1980},
		{"id": 3, "bornYear": 1990}
	]`))
	names.NaturalJoin(years).Project("bornYear", "name").Serialize(os.Stdout)
	// Output:
	// [{"bornYear": 1979, "name": "Lee"},
	// {"bornYear": 1980, "name": "Jake"},
	// {"bornYear": 1990, "name": "Kristen"}]
}
