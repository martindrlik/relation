package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func BenchmarkInsertTupleTo1000(b *testing.B) {
	r := rex.NewRelation()
	num := func(i int) func() (string, any) { return func() (string, any) { return "n", i } }
	for i := 0; i < 1000; i++ {
		r.InsertOne(num(i))
	}
	item := map[string]any{"n": 1000}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.InsertTuple(item)
	}
}
