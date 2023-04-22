package rex_test

import (
	"os"

	"github.com/martindrlik/rex"
)

func ExampleSerialize() {
	rex.NewRelation().
		InsertOne(name("Jake"), bornYear(1980)).
		InsertOne(name("Lee"), bornYear(1979)).
		InsertOne(name("Kristen"), bornYear(1990)).
		Serialize(os.Stdout)
	// Output:
	// [{"bornYear": 1979, "name": "Lee"},
	// {"bornYear": 1980, "name": "Jake"},
	// {"bornYear": 1990, "name": "Kristen"}]
}
