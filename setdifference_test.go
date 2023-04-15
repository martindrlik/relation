package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestSetDifference(t *testing.T) {
	r := rex.NewRelation().
		InsertOne(name("Harry")).
		InsertOne(name("Sally")).
		InsertOne(name("George"))
	s := rex.NewRelation().
		InsertOne(name("George"))
	expected := rex.NewRelation().
		InsertOne(name("Harry")).
		InsertOne(name("Sally"))
	if !expected.Equals(r.SetDifference(s)) {
		t.Error("expected equal after set difference")
	}
}
