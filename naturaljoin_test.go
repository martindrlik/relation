package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestNaturalJoin(t *testing.T) {
	employee := rex.NewRelation().
		InsertOne(name("Harry"), empId(3415), deptName("Finance")).
		InsertOne(name("Sally"), empId(2241), deptName("Sales")).
		InsertOne(name("George"), empId(3401), deptName("Finance")).
		InsertOne(name("Harriet"), empId(2202), deptName("Sales")).
		InsertOne(name("Mary"), empId(1257), deptName("Human Resources"))
	dept := rex.NewRelation().
		InsertOne(deptName("Finance"), manager("George")).
		InsertOne(deptName("Sales"), manager("Harriet")).
		InsertOne(deptName("Production"), manager("Charles"))
	expected := rex.NewRelation().
		InsertOne(name("Harry"), empId(3415), deptName("Finance"), manager("George")).
		InsertOne(name("Sally"), empId(2241), deptName("Sales"), manager("Harriet")).
		InsertOne(name("George"), empId(3401), deptName("Finance"), manager("George")).
		InsertOne(name("Harriet"), empId(2202), deptName("Sales"), manager("Harriet"))
	if !expected.Equals(employee.NaturalJoin(dept)) {
		t.Error("expected equal after natural join")
	}
}

func empId(id int) func() (string, any) {
	return func() (string, any) {
		return "empId", id
	}
}

func deptName(dept string) func() (string, any) {
	return func() (string, any) {
		return "deptName", dept
	}
}

func manager(name string) func() (string, any) {
	return func() (string, any) {
		return "manager", name
	}
}
