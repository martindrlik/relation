package store_test

import (
	"testing"

	"github.com/martindrlik/store"
)

func TestColumn(t *testing.T) {
	t.Run("insert one", func(t *testing.T) {
		column := store.Column{}
		column.Insert("foo")
		if n := column.Len(); n != 1 {
			t.Errorf("expected length to be 1 got %d", n)
		}
		if v, ok := column.At(0); !ok || v != "foo" {
			t.Errorf("expected \"foo\" got %v %v", v, ok)
		}
	})
	t.Run("delete last of two", func(t *testing.T) {
		column := store.Column{}
		column.Insert("foo")
		column.Insert("bar")
		column.Delete(1)
		if n := column.Len(); n != 1 {
			t.Errorf("expected length to be 1 got %d", n)
		}
		if v, ok := column.At(0); !ok || v != "foo" {
			t.Errorf("expected last \"foo\" move to first but got %v %v", v, ok)
		}
	})
	t.Run("delete first of two", func(t *testing.T) {
		column := store.Column{}
		column.Insert("foo")
		column.Insert("bar")
		column.Delete(0)
		if n := column.Len(); n != 1 {
			t.Errorf("expected length to be 1 got %d", n)
		}
		if v, ok := column.At(0); !ok || v != "bar" {
			t.Errorf("expected last \"bar\" move to first but got %v %v", v, ok)
		}
	})
	t.Run("delete last one", func(t *testing.T) {
		column := store.Column{}
		column.Insert("foo")
		column.Delete(0)
		if n := column.Len(); n != 0 {
			t.Errorf("expected length to be 0 got %d", n)
		}
	})
	t.Run("out of range delete nothing removed", func(t *testing.T) {
		column := store.Column{}
		column.Insert("foo")
		column.Delete(-1)
		column.Delete(1)
		if v, ok := column.At(0); !ok || v != "foo" {
			t.Errorf("expected \"foo\" got %v %v", v, ok)
		}
	})
}
