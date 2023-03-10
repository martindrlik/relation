package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestColumn(t *testing.T) {
	t.Run("insert one", func(t *testing.T) {
		col := rex.Column{}
		col.Insert("foo")
		if ln := col.Len(); ln != 1 {
			t.Errorf("expected length to be 1 got %d", ln)
		}
		if v, ok := col.At(0); !ok || v != "foo" {
			t.Errorf("expected \"foo\" got %v %v", v, ok)
		}
	})
	t.Run("remove last one of two", func(t *testing.T) {
		col := rex.Column{}
		col.Insert("foo")
		col.Insert("bar")
		col.RemoveAt(1)
		if ln := col.Len(); ln != 1 {
			t.Errorf("expected length to be 1 got %d", ln)
		}
		if v, ok := col.At(0); !ok || v != "foo" {
			t.Errorf("expected last \"foo\" move to first but got %v %v", v, ok)
		}
	})
	t.Run("remove first one of two", func(t *testing.T) {
		col := rex.Column{}
		col.Insert("foo")
		col.Insert("bar")
		col.RemoveAt(0)
		if ln := col.Len(); ln != 1 {
			t.Errorf("expected length to be 1 got %d", ln)
		}
		if v, ok := col.At(0); !ok || v != "bar" {
			t.Errorf("expected last \"bar\" move to first but got %v %v", v, ok)
		}
	})
	t.Run("remove last", func(t *testing.T) {
		col := rex.Column{}
		col.Insert("foo")
		col.RemoveAt(0)
		if ln := col.Len(); ln != 0 {
			t.Errorf("expected length to be 0 got %d", ln)
		}
	})
	t.Run("remove out of range", func(t *testing.T) {
		col := rex.Column{}
		col.Insert("foo")
		col.RemoveAt(-1)
		col.RemoveAt(1)
		if v, ok := col.At(0); !ok || v != "foo" {
			t.Errorf("expected \"foo\" got %v %v", v, ok)
		}
	})
}
