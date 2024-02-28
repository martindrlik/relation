package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestRelation(t *testing.T) {
	r1 := must((&rex.R{}).Add(rex.T{"name": "John", "age": 42}))
	r2 := must((&rex.R{}).Add(rex.T{"name": "John", "age": 42}))
	r3 := must((&rex.R{}).Add(rex.T{"name": "Jake", "age": 34}))
	r4 := must((&rex.R{}).Add(rex.T{"name": "John", "age": 42, "city": "London"}))
	r5 := must(must((&rex.R{}).Add(rex.T{"name": "John", "age": 42})).Add(rex.T{"name": "Jake", "age": 34}))
	t.Run("Equal", func(t *testing.T) {
		if !r1.Equal(r2) {
			t.Error("r1 and r2 should be equal")
		}
	})
	t.Run("NotEqual", func(t *testing.T) {
		if r1.Equal(r3) {
			t.Error("r1 and r3 should not be equal")
		}
		if r1.Equal(r4) {
			t.Error("r1 and r4 should not be equal (different schema)")
		}
		if r1.Equal(r5) {
			t.Error("r1 and r5 should not be equal (r5 has more tuples)")
		}
	})
}
