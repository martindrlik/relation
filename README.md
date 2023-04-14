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
```