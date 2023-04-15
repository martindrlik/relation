package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestNaturalJoin(t *testing.T) {
	employees := rex.NewRelation().
		InsertMany(
			employee("Harry", 3415, "Finance"),
			employee("Sally", 2241, "Sales"),
			employee("George", 3401, "Finance"),
			employee("Harriet", 2202, "Sales"),
			employee("Mary", 1257, "Human Resources"))
	departments := rex.NewRelation().
		InsertMany(
			department("Finance", "George"),
			department("Sales", "Harriet"),
			department("Production", "Charles"))
	expected := rex.NewRelation().
		InsertMany(
			join(employee("Harry", 3415, "Finance"), department("Finance", "George")),
			join(employee("Sally", 2241, "Sales"), department("Sales", "Harriet")),
			join(employee("George", 3401, "Finance"), department("Finance", "George")),
			join(employee("Harriet", 2202, "Sales"), department("Sales", "Harriet")))
	if !expected.Equals(employees.NaturalJoin(departments)) {
		t.Error("expected equal after natural join")
	}
}

func employee(name string, id int, department string) []func() (string, any) {
	return rex.One(
		func() (string, any) { return "name", name },
		func() (string, any) { return "empId", id },
		func() (string, any) { return "deptName", department })
}

func department(name, manager string) []func() (string, any) {
	return rex.One(
		func() (string, any) { return "deptName", name },
		func() (string, any) { return "manager", manager })
}

func join[T any](a ...[]T) []T {
	r := []T{}
	for _, a := range a {
		r = append(r, a...)
	}
	return r
}
