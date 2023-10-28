package rex_test

import (
	"testing"

	"golang.org/x/exp/maps"
)

var (
	adventureTime    = show("Adventure Time")
	adventureRelease = born(2010, 4, 5)
	finn             = glue(adventureRelease, finnName)
	finnName         = name("Finn")
	jake             = glue(adventureRelease, jakeName)
	jakeName         = name("Jake")
	marceline        = glue(adventureRelease, marcelineName)
	marcelineName    = name("Marceline")

	futurama        = show("Futurama")
	futuramaRelease = born(1999, 3, 28)
	leelaName       = name("Leela")
	leela           = glue(futuramaRelease, leelaName)
)

func TestRelation(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		actual := take1(in(finn)).m
		if !maps.Equal(actual, finn) {
			t.Errorf("%v is not equal to %v", actual, finn)
		}
	})
	t.Run("two", func(t *testing.T) {
		first, second := take2(in(finn, marceline))
		if !maps.Equal(first.m, finn) {
			t.Errorf("first %v is not equal to %v", first, finn)
		}
		if !maps.Equal(second.m, marceline) {
			t.Errorf("second %v is not equal to %v", second, marceline)
		}
	})
	t.Run("duplicate", func(t *testing.T) {
		count := 0
		in(finn, finn).Each(func(m map[string]any, b bool) bool {
			count++
			return true
		})
		if count != 1 {
			t.Errorf("expected one tuple got %v", count)
		}
	})
}
