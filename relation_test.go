package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestInsert(t *testing.T) {
	t.Run("duplicate", func(t *testing.T) {
		type tc struct {
			name string
			json []string
		}
		for _, tc := range []tc{
			{"simple", []string{`{"name": "Jake"}`, `{"name": "Jake"}`}},
			{"nested", []string{`{"address": {"city": "New York"}}`, `{"address": {"city": "New York"}}`}},
			{"nested unordered", []string{
				`{"address": {"city": "New York", "street": "Broadway"}}`,
				`{"address": {"street": "Broadway", "city": "New York"}}`}},
		} {
			r := rex.R{}
			for _, s := range tc.json {
				must(r.Insert(rex.String(s)))
			}
			if n := r.Len(); n != 1 {
				t.Errorf("[%s] expected one tuple got %v", tc.name, n)
			}
		}
	})
}

func TestEquals(t *testing.T) {
	u := rex.R{}
	v := rex.R{}
	must(u.Insert(rex.String(`{"address": {"city": "New York", "street": "Broadway", "x": {}}}`)))
	must(v.Insert(rex.String(`{"address": {"city": "New York", "street": "Broadway", "x": {}}}`)))
	if !u.Equals(v) {
		t.Errorf("u should be equal to v")
	}
}

func TestNotEquals(t *testing.T) {
	t.Run("value", func(t *testing.T) {
		u := rex.R{}
		v := rex.R{}
		must(u.Insert(rex.String(`{"name": "Jake"}`)))
		must(v.Insert(rex.String(`{"name": "Luke"}`)))
		if u.Equals(v) {
			t.Errorf("u should not be equal to v")
		}
	})
	t.Run("attributes", func(t *testing.T) {
		u := rex.R{}
		v := rex.R{}
		must(u.Insert(rex.String(`{"name": "Jake"}`)))
		must(v.Insert(rex.String(`{"username": "Jake"}`)))
		if u.Equals(v) {
			t.Errorf("u should not be equal to v")
		}
	})
	t.Run("attributes 2", func(t *testing.T) {
		u := rex.R{}
		v := rex.R{}
		must(u.Insert(rex.String(`{"name": "Jake"}`)))
		must(v.Insert(rex.String(`{"username": "Jake", "age": 35}`)))
		if u.Equals(v) {
			t.Errorf("u should not be equal to v")
		}
	})
	t.Run("nested", func(t *testing.T) {
		u := rex.R{}
		v := rex.R{}
		must(u.Insert(rex.String(`{"address": {"city": "New York", "street": "Broadway"}}`)))
		must(v.Insert(rex.String(`{"address": {"city": "New York", "street": "Park Avenue"}}`)))
		if u.Equals(v) {
			t.Errorf("u should not be equal to v")
		}
	})
}
