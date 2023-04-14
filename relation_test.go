package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestEquals(t *testing.T) {
	t.Run("not equals", func(t *testing.T) {
		r := rex.NewRelation()
		s := rex.NewRelation()
		r.InsertOne(city("New York"))
		s.InsertOne(city("Olomouc"))
		if r.Equals(s) {
			t.Error("expected to be not equal")
		}
	})
	t.Run("not equals nested", func(t *testing.T) {
		r := rex.NewRelation().InsertOne(
			city("New York"),
			streets(rex.NewRelation().
				InsertOne(street("Broadway")).
				InsertOne(street("Park Avenue"))))
		s := rex.NewRelation().InsertOne(
			city("New York"),
			streets(rex.NewRelation().
				InsertOne(street("Broadway")).
				InsertOne(street("St. Mark’s Place"))))
		if r.Equals(s) {
			t.Error("expected to be not equal")
		}
	})
	t.Run("equals", func(t *testing.T) {
		r := rex.NewRelation()
		s := rex.NewRelation()
		r.InsertOne(city("New York"))
		s.InsertOne(city("New York"))
		if !r.Equals(s) {
			t.Error("expected to be equal")
		}
	})
	t.Run("equals nested", func(t *testing.T) {
		r := rex.NewRelation().InsertOne(
			city("New York"),
			streets(rex.NewRelation().
				InsertOne(street("Broadway")).
				InsertOne(street("Park Avenue"))))
		s := rex.NewRelation().InsertOne(
			city("New York"),
			streets(rex.NewRelation().
				InsertOne(street("Broadway")).
				InsertOne(street("Park Avenue"))))
		if !r.Equals(s) {
			t.Error("expected to be equal")
		}
	})
}

func city(city any) func() (string, any) {
	return func() (string, any) {
		return "city", city
	}
}

func street(street any) func() (string, any) {
	return func() (string, any) {
		return "street", street
	}
}

func streets(streets any) func() (string, any) {
	return func() (string, any) {
		return "streets", streets
	}
}

func bornYear(year int) func() (string, any) {
	return func() (string, any) {
		return "bornYear", year
	}
}

func name(name string) func() (string, any) {
	return func() (string, any) {
		return "name", name
	}
}

func attr(fn func() (string, any)) string {
	a, _ := fn()
	return a
}

func value(fn func() (string, any)) any {
	_, v := fn()
	return v
}
