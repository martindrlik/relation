package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestRS(t *testing.T) {
	t.Run("R", func(t *testing.T) {
		rs := rex.New()
		rs.Add(rex.T{"a": 10})
		rs.Add(rex.T{"b": 11})
		rs.Add(rex.T{"a": "ar", "b": "bar"})
		r := rs.R()
		s := rex.R{rex.T{"a": "ar", "b": "bar"}}
		if !r.IsEqual(&s) {
			t.Errorf("expected %v to be equal to %v", r, s)
		}
	})
}
