package rex_test

import (
	"fmt"

	"github.com/martindrlik/rex"
)

func ExampleNaturalJoin() {
	q := rex.Table{}
	r := rex.Table{}
	q.InsertOne(`{"x": 2, "y": 3}`)
	r.InsertOne(`{"x": 2, "y": 3, "z": 5}`)
	t := q.NaturalJoin(&r).Project("x", "y", "z").At(0)
	fmt.Printf("%.0f + %.0f = %.0f", t[0], t[1], t[2])
	// Output: 2 + 3 = 5
}
