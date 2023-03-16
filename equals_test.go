package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

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
	u := rex.R{}
	v := rex.R{}
	must(u.Insert(rex.String(`{"name": "Jake", "address": {"city": "New York", "street": "Broadway"}}`)))
	must(v.Insert(rex.String(`{"name": "Luke", "address": {"city": "New York", "street": "Broadway"}}`)))
	if u.Equals(v) {
		t.Errorf("u should not be equal to v")
	}
}
