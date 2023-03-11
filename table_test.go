package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestTable(t *testing.T) {
	t.Run("insert one", func(t *testing.T) {
		table := rex.Table{}
		actual := dump(table.InsertOne(`{"name": "Foo", "value": 1990}`).At(0)...)
		expect := dump("Foo", float64(1990))
		if actual != expect {
			t.Errorf("expected tuple %v, got %v", expect, actual)
		}
	})
	t.Run("insert two", func(t *testing.T) {
		table := rex.Table{}
		table.InsertOne(`{"name": "Foo"}`).InsertOne(`{"value": 1990}`)
		actual := dump(table.At(0)...)
		expect := dump("Foo", rex.Empty{})
		if actual != expect {
			t.Errorf("expected first tuple %v, got %v", expect, actual)
		}
		actual = dump(table.At(1)...)
		expect = dump(rex.Empty{}, float64(1990))
		if actual != expect {
			t.Errorf("expected second tuple %v, got %v", expect, actual)
		}
	})
	t.Run("insert two with same column", func(t *testing.T) {
		table := rex.Table{}
		table.InsertOne(`{"name": "Foo"}`).InsertOne(`{"name": "Bar", "value": 1990}`)
		actual := dump(table.At(0)...)
		expect := dump("Foo", rex.Empty{})
		if actual != expect {
			t.Errorf("expected first tuple %v, got %v", expect, actual)
		}
		actual = dump(table.At(1)...)
		expect = dump("Bar", float64(1990))
		if actual != expect {
			t.Errorf("expected second tuple %v, got %v", expect, actual)
		}
	})
}
