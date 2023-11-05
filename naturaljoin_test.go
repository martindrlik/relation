package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestNaturalJoin(t *testing.T) {
	left := in(
		map[string]any{"x": 1, "name": "Finn"},
		map[string]any{"x": 2, "name": "Jake"})
	right := in(
		map[string]any{"x": 1, "age": 17},
		map[string]any{"x": 2, "age": 28})
	joined := rex.NaturalJoin(left, right)
	finnAndAge := map[string]any{"x": 1, "name": "Finn", "age": 17}
	jakeAndAge := map[string]any{"x": 2, "name": "Jake", "age": 28}
	if !joined.Contains(finnAndAge) {
		t.Errorf("expected %v to contain %v", joined, finnAndAge)
	}
	if !joined.Contains(jakeAndAge) {
		t.Errorf("expected %v to contain %v", joined, jakeAndAge)
	}
}
