package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestRestrict(t *testing.T) {
	kristen := name("Kristen")
	r := rex.NewRelation().
		InsertOne(kristen, bornYear(1990))
	s := rex.NewRelation().
		InsertOne(name("Jake"), bornYear(1980)).
		InsertOne(name("Lee"), bornYear(1979)).
		InsertOne(kristen, bornYear(1990))
	if !r.Equals(s.Restrict(func(tuple map[string]any) bool {
		return tuple[attr(kristen)] == value(kristen)
	})) {
		t.Error("expected equal after restriction")
	}
}
