package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestDifference(t *testing.T) {
	adventureTime := in(finn, jake, marceline)
	vampires := in(marceline)
	difference := rex.Difference(adventureTime, vampires)
	if difference.Contains(marceline) {
		t.Errorf("expected %v - %v not to contain %v", adventureTime, vampires, marceline)
	}
}
