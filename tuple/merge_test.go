package tuple_test

import (
	"maps"
	"testing"

	"github.com/martindrlik/rex/tuple"
)

func TestMerge(t *testing.T) {
	u := map[string]any{"foo": 1, "bar": 2.0}
	v := map[string]any{"foo": 1, "baz": "3"}
	actual := tuple.Merge(u, v)
	expect := map[string]any{"foo": 1, "bar": 2.0, "baz": "3"}
	if !maps.Equal(actual, expect) {
		t.Errorf("expected %v got %v", expect, actual)
	}

	t.Run("overwrite", func(t *testing.T) {
		const fooval = 2
		v := map[string]any{"foo": fooval, "baz": "3"}
		actual := tuple.Merge(u, v)
		expect := map[string]any{"foo": fooval, "bar": 2.0, "baz": "3"}
		if !maps.Equal(actual, expect) {
			t.Errorf("expected %v got %v", expect, actual)
		}
	})
}
