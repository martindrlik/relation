package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestUnion(t *testing.T) {
	r := rex.NewRelation().
		InsertOne(name("Harry")).
		InsertOne(name("Sally"))
	s := rex.NewRelation().
		InsertOne(name("George")).
		InsertOne(name("Harriet")).
		InsertOne(name("Mary"))
	expected := rex.NewRelation().
		InsertOne(name("Harry")).
		InsertOne(name("Sally")).
		InsertOne(name("George")).
		InsertOne(name("Harriet")).
		InsertOne(name("Mary"))
	if !expected.Equals(r.Union(s)) {
		t.Error("expected equal after union")
	}
}
