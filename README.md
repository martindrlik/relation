# rex

Experimental relational NoSQL database. It is my playground for ideas and API will change over time. There is a lot more to do before it can be even considered interesting.

## Examples

```go
func TestProject(t *testing.T) {
	r := rex.NewRelation()
	s := rex.NewRelation()
	bornYear := bornYear(1980)
	r.InsertOne(bornYear)
	s.InsertOne(bornYear, name("Jake"))
	if !r.Equals(s.Project(attr(bornYear))) {
		t.Error("expected equal after projection")
	}
}

func TestRestrict(t *testing.T) {
	kristen := name("Kristen")
	r := rex.NewRelation().
		InsertOne(kristen, bornYear(1990))
	s := rex.NewRelation().
		InsertOne(name("Jake"), bornYear(1980)).
		InsertOne(name("Lee"), bornYear(1979)).
		InsertOne(kristen, bornYear(1990))
	if !r.Equals(s.Restrict(func(tuple map[string]any) bool {
		return tuple[attr(kristen)] == value(kristen)
	})) {
		t.Error("expected equal after restriction")
	}
}

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

func TestUnion(t *testing.T) {
	r := rex.NewRelation().
		InsertOne(name("Harry")).
		InsertOne(name("Sally"))
	s := rex.NewRelation().
		InsertOne(name("George")).
		InsertOne(name("Harriet")).
		InsertOne(name("Mary"))
	expected := rex.NewRelation().
		InsertOne(name("Harry")).
		InsertOne(name("Sally")).
		InsertOne(name("George")).
		InsertOne(name("Harriet")).
		InsertOne(name("Mary"))
	if !expected.Equals(r.Union(s)) {
		t.Error("expected equal after union")
	}
}
```