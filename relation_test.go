package rex

import (
	"testing"
)

func TestRelation(t *testing.T) {
	t.Run("IsEmpty", func(t *testing.T) {
		r := R{}
		if !r.IsEmpty() {
			t.Errorf("expected %v to be empty", r)
		}
	})
	t.Run("Add", func(t *testing.T) {
		r := R{}
		r.Add(T{"a": 10})
		if r.IsEmpty() {
			t.Errorf("expected %v to not be empty", r)
		}
	})
	t.Run("Remove", func(t *testing.T) {
		r := R{}
		r.Add(T{"a": 10})
		r.Remove(T{"a": 10})
		if !r.IsEmpty() {
			t.Errorf("expected %v to be empty", r)
		}
	})
	t.Run("Len", func(t *testing.T) {
		r := R{}
		r.Add(T{"a": 10})
		if r.Len() != 1 {
			t.Errorf("expected %v to have length 1", r)
		}
		r.Add(T{"a": 11})
		if r.Len() != 2 {
			t.Errorf("expected %v to have length 2", r)
		}

		r.Add(T{"a": 11})
		if r.Len() != 2 {
			t.Errorf("expected %v to have length 2 as no new tuple added", r)
		}
	})
	t.Run("IsEqual", func(t *testing.T) {
		r := R{}
		r.Add(T{"a": 10, "b": 11})
		s := R{}
		s.Add(T{"a": 10, "b": 11})
		if !r.IsEqual(&s) {
			t.Errorf("expected %v to be equal to %v", r, s)
		}

		a := R{}
		a.Add(T{"a": 10})
		if r.IsEqual(&a) {
			t.Errorf("expected %v to not be equal to %v", r, a)
		}
	})
	t.Run("HasAttributes", func(t *testing.T) {
		r := R{}
		r.Add(T{"a": 10, "b": 11})
		if !r.HasAttributes("a", "b") {
			t.Errorf("expected %v to have attributes a, b", r)
		}
		if r.HasAttributes("a", "b", "c") {
			t.Errorf("expected %v to not have attributes a, b, c", r)
		}
	})
	t.Run("Project", func(t *testing.T) {
		r := R{}
		r.Add(T{"a": 10, "b": 11})
		a := R{}
		a.Add(T{"a": 10})
		if !r.Project("a").IsEqual(&a) {
			t.Errorf("expected %v to be equal to %v", r.Project("a"), a)
		}

		b := R{}
		b.Add(T{"b": 11})
		if !b.Project("a").IsEmpty() {
			t.Errorf("expected (b) projecting on (a) to be empty got %v", b.Project("a"))
		}
	})
	t.Run("Union", func(t *testing.T) {
		a1, a2 := R{}, R{}
		a1.Add(T{"a": 1})
		a2.Add(T{"a": 2})
		a12 := R{}
		a12.Add(T{"a": 1})
		a12.Add(T{"a": 2})
		if !a1.Union(&a2).IsEqual(&a12) {
			t.Errorf("expected %v to be equal to %v", a1.Union(&a2), a12)
		}

		// not union compatible
		a, b := R{}, R{}
		a.Add(T{"a": 1})
		b.Add(T{"b": 2})
		if !a.Union(&b).IsEmpty() {
			t.Errorf("expected %v to be empty (a) and (b) are not union compatible", a.Union(&b))
		}
	})
	t.Run("Difference", func(t *testing.T) {
		a := R{}
		a.Add(T{"a": 1})
		if !a.Difference(&a).IsEmpty() {
			t.Errorf("expected %v to be empty", a.Difference(&a))
		}

		a1, a2 := R{}, R{}
		a1.Add(T{"a": 1})
		a2.Add(T{"a": 2})
		if !a1.Difference(&a2).IsEqual(&a1) {
			t.Errorf("expected %v to be equal to %v", a1.Difference(&a2), a1)
		}

		// not union compatible
		a, b := R{}, R{}
		a.Add(T{"a": 1})
		b.Add(T{"b": 2})
		if !a.Difference(&b).IsEmpty() {
			t.Errorf("expected %v to be empty (a) and (b) are not union compatible", a.Difference(&b))
		}
	})
	t.Run("NaturalJoin", func(t *testing.T) {
		ab, bc := R{}, R{}
		ab.Add(T{"a": 10, "b": 11})
		bc.Add(T{"b": 11, "c": 12})
		abc := R{}
		abc.Add(T{"a": 10, "b": 11, "c": 12})
		if !ab.NaturalJoin(&bc).IsEqual(&abc) {
			t.Errorf("expected %v to be equal to %v", ab.NaturalJoin(&bc), abc)
		}

		ax, ay := R{}, R{}
		ax.Add(T{"a": "x"})
		ay.Add(T{"a": "y"})
		if !ax.NaturalJoin(&ay).IsEmpty() {
			t.Errorf("expected %v to be empty", ax.NaturalJoin(&ay))
		}

		empty := R{}
		if !ax.NaturalJoin(&empty).IsEmpty() {
			t.Errorf("expected %v natural join empty to be empty got %v", ax, ax.NaturalJoin(&empty))
		}
	})
	t.Run("CartesianProduct", func(t *testing.T) {
		a, b := R{}, R{}
		a.Add(T{"a": 1})
		a.Add(T{"a": 2})
		b.Add(T{"b": 10})
		b.Add(T{"b": 20})
		cp := R{}
		cp.Add(T{"a": 1, "b": 10})
		cp.Add(T{"a": 1, "b": 20})
		cp.Add(T{"a": 2, "b": 10})
		cp.Add(T{"a": 2, "b": 20})
		if !a.NaturalJoin(&b).IsEqual(&cp) {
			t.Errorf("expected %v to be equal to %v", a.NaturalJoin(&b), cp)
		}
	})
}
