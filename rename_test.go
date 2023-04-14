package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestRename(t *testing.T) {
	newYorkCity := city("New York")
	newYorkName := name("New York")
	r := rex.NewRelation().InsertOne(newYorkCity)
	s := rex.NewRelation().InsertOne(newYorkName)
	if !r.Equals(s.Rename(map[string]string{
		attr(newYorkName): attr(newYorkCity),
	})) {
		t.Error("expected equal after rename")
	}
}

func TestIllegalRename(t *testing.T) {
	newYorkCity := city("New York")
	r := rex.NewRelation().InsertOne(newYorkCity)
	defer func() {
		s := recover().(string)
		expected := "illegal rename city to city"
		if s != expected {
			t.Errorf("expected %q got %q", expected, s)
		}
	}()
	r.Rename(map[string]string{
		attr(newYorkCity): attr(newYorkCity),
	})
}
