package store_test

import (
	"testing"

	"github.com/martindrlik/store"
)

func TestTable(t *testing.T) {
	t.Run("insert one", func(t *testing.T) {
		table := store.Table{}
		nerr(table.InsertOne(`{"name": "Foo", "value": 1990}`))
		tuple, ok := table.At(0)
		actual := dump(tuple...)
		expect := dump("Foo", float64(1990))
		if !ok || actual != expect {
			t.Errorf("expected tuple %v and ok to be true, got %v and %v", expect, actual, ok)
		}
	})
	t.Run("insert two", func(t *testing.T) {
		table := store.Table{}
		nerr(table.InsertOne(`{"name": "Foo"}`))
		nerr(table.InsertOne(`{"value": 1990}`))
		tuple, ok := table.At(0)
		actual := dump(tuple...)
		expect := dump("Foo", store.NoValue{})
		if !ok || actual != expect {
			t.Errorf("expected first tuple %v and ok to be true, got %v and %v", expect, actual, ok)
		}

		tuple, ok = table.At(1)
		actual = dump(tuple...)
		expect = dump(store.NoValue{}, float64(1990))
		if !ok || actual != expect {
			t.Errorf("expected second tuple %v and ok to be true, got %v and %v", expect, actual, ok)
		}
	})
	t.Run("insert two with same column", func(t *testing.T) {
		table := store.Table{}
		nerr(table.InsertOne(`{"name": "Foo"}`))
		nerr(table.InsertOne(`{"name": "Bar", "value": 1990}`))
		tuple, ok := table.At(0)
		actual := dump(tuple...)
		expect := dump("Foo", store.NoValue{})
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
