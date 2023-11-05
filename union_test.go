package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
	"golang.org/x/exp/maps"
)

func TestUnion(t *testing.T) {
	actual1st, actual2th := take2(rex.Union(in(finn), in(marceline)))
	expect1st, expect2th := take2(in(finn, marceline))
	if !maps.Equal(actual1st, expect1st) {
		t.Errorf("first %v is not equal to %v", actual1st, expect1st)
	}
	if !maps.Equal(actual2th, expect2th) {
		t.Errorf("second %v is not equal to %v", actual2th, expect2th)
	}
}
