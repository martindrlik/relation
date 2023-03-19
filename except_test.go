package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestExcept(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		u, v := rex.R{}, rex.R{}
		must(u.Insert(rex.String(`{"name": "Jake"}`)))
		must(v.Insert(rex.String(`{"name": "Jake"}`)))
		w := u.Except(v)
		if n := w.Len(); n != 0 {
			t.Errorf("expected result to have 0 tuples, got %v", n)
		}
	})
	t.Run("simple 2", func(t *testing.T) {
		u, v := rex.R{}, rex.R{}
		must(u.Insert(rex.String(`{"name": "Jake"}`)))
		must(u.Insert(rex.String(`{"name": "Luke"}`)))
		must(v.Insert(rex.String(`{"name": "Jake"}`)))
		w := u.Except(v)
		if n := w.Len(); n != 1 {
			t.Errorf("expected result to have 1 tuples, got %v", n)
		}
	})
	t.Run("nested", func(t *testing.T) {
		u, v := rex.R{}, rex.R{}
		must(u.Insert(rex.String(`{"name": "Jake", "personal": {"age": 35}}`)))
		must(v.Insert(rex.String(`{"name": "Jake", "personal": {"age": 35}}`)))
		w := u.Except(v)
		if n := w.Len(); n != 0 {
			t.Errorf("expected result to have 0 tuples, got %v", n)
		}
	})
	t.Run("nested 2", func(t *testing.T) {
		u, v := rex.R{}, rex.R{}
		must(u.Insert(rex.String(`{"name": "Jake", "personal": {"age": 35}}`)))
		must(u.Insert(rex.String(`{"name": "Jake", "personal": {"age": 24}}`)))
		must(v.Insert(rex.String(`{"name": "Jake", "personal": {"age": 35}}`)))
		w := u.Except(v)
		if n := w.Len(); n != 1 {
			t.Errorf("expected result to have 1 tuples, got %v", n)
		}
	})
	t.Run("nested 3", func(t *testing.T) {
		u, v := rex.R{}, rex.R{}
		must(u.Insert(rex.String(`{"x": {"name": "Jake"}}`)))
		must(v.Insert(rex.String(`{"y": {"name": "Jake"}}`)))
		w := u.Except(v)
		if n := w.Len(); n != 1 {
			t.Errorf("expected result to have 1 tuples, got %v", n)
		}
	})
	t.Run("nested 4", func(t *testing.T) {
		u, v := rex.R{}, rex.R{}
		must(u.Insert(rex.String(`{"x": "Foo"}`)))
		must(v.Insert(rex.String(`{"x": 1}`)))
		w := u.Except(v)
		if n := w.Len(); n != 1 {
			t.Errorf("expected result to have 1 tuples, got %v", n)
		}
	})
}
