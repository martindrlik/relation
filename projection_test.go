package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestProjection(t *testing.T) {
	adventureTime := rex.NewRelation().Insert(
		map[string]any{"name": "Finn", "last": "Mertens"},
	)
	projected := adventureTime.Projection("name")
	if !projected.Contains(map[string]any{"name": "Finn"}) {
		t.Errorf("expected %v to contain %v", projected, map[string]any{"name": "Finn"})
	}
}
