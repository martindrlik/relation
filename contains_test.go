package rex_test

import (
	"testing"
)

func TestContains(t *testing.T) {
	t.Run("contains", func(t *testing.T) {
		if !in(finn).Contains(finn) {
			t.Errorf("expected %v to contain %v", in(finn), finn)
		}
		if !in(finn, marceline).Contains(marceline) {
			t.Errorf("expected %v to contain %v", in(finn, marceline), marceline)
		}
	})
	t.Run("not contain", func(t *testing.T) {
		adventure := in(glue(finnName), glue(jakeName), glue(marcelineName))
		if adventure.Contains(leela) {
			t.Errorf("expected %v not to contain %v", adventure, leela)
		}
	})
}
