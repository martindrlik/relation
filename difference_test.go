package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestDifference(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		adventure := rex.Difference(in(finn, jake, marceline, leela), in(leela))
		if adventure.Contains(leela) {
			t.Errorf("%v is not character of Adventure Time", leela)
		}
	})
	t.Run("no common attribute", func(t *testing.T) {
		adventure := rex.Difference(in(finn, jake, marceline), in(glue(futuramaRelease)))
		for _, m := range []map[string]any{finn, jake, marceline} {
			if !adventure.Contains(m) {
				t.Errorf("adventure should contain %v", m)
			}
		}
	})
}
