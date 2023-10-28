package rex_test

import (
	"testing"

	"golang.org/x/exp/maps"
)

func TestProjection(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		actual := take1(in(finn).Projection(attr(finnName))).m
		if expect := glue(finnName); !maps.Equal(actual, expect) {
			t.Errorf("%v is not equal to %v", actual, expect)
		}
	})
	t.Run("no attribute to project", func(t *testing.T) {
		count := 0
		in(glue(finnName)).Projection(attr(adventureRelease)).Each(func(m map[string]any, b bool) bool {
			count++
			return true
		})
		if count != 0 {
			t.Errorf("expected count to be zero got %v", count)
		}
	})
}
