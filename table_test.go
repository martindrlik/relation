package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestTable(t *testing.T) {
	t.Run("insert one", func(t *testing.T) {
		table := rex.Table{}
		tuple, ok := table.InsertOne(`{"name": "Foo", "value": 1990}`).At(0)
		actual := dump(tuple...)
		expect := dump("Foo", float64(1990))
		if !ok || actual != expect {
			t.Errorf("expected tuple %v and ok to be true, got %v and %v", expect, actual, ok)
		}
	})
	t.Run("insert two", func(t *testing.T) {
		table := rex.Table{}
		table.InsertOne(`{"name": "Foo"}`).InsertOne(`{"value": 1990}`)
		tuple, ok := table.At(0)
		actual := dump(tuple...)
		expect := dump("Foo", rex.Empty{})
		if !ok || actual != expect {
			t.Errorf("expected first tuple %v and ok to be true, got %v and %v", expect, actual, ok)
		}
		tuple, ok = table.At(1)
		actual = dump(tuple...)
		expect = dump(rex.Empty{}, float64(1990))
		if !ok || actual != expect {
			t.Errorf("expected second tuple %v and ok to be true, got %v and %v", expect, actual, ok)
		}
	})
	t.Run("insert two with same column", func(t *testing.T) {
		table := rex.Table{}
		table.InsertOne(`{"name": "Foo"}`).InsertOne(`{"name": "Bar", "value": 1990}`)
		tuple, ok := table.At(0)
		actual := dump(tuple...)
		expect := dump("Foo", rex.Empty{})
		if !ok || actual != expect {
			t.Errorf("expected first tuple %v and ok to be true, got %v and %v", expect, actual, ok)
		}
		tuple, ok = table.At(1)
		actual = dump(tuple...)
		expect = dump("Bar", float64(1990))
		if !ok || actual != expect {
			t.Errorf("expected second tuple %v and ok to be true, got %v and %v", expect, actual, ok)
		}
	})
}
