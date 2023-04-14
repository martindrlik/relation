package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestProject(t *testing.T) {
	r := rex.NewRelation()
	s := rex.NewRelation()
	bornYear := bornYear(1980)
	r.InsertOne(bornYear)
	s.InsertOne(bornYear, name("Jake"))
	if !r.Equals(s.Project(attr(bornYear))) {
		t.Error("expected equal after projection")
	}
}
