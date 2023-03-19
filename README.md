# rex

Experimental relational NoSQL database. It is my playground for ideas and API will change over time. There is a lot more to do before it can be even considered interesting.

## Examples

```golang
func ExampleDump() {
	users := rex.R{}
	must(users.Insert(rex.String(`{"name": "Jake"}`)))
	must(users.Insert(rex.String(`{"age": 35}`)))
	must(users.Insert(rex.String(`{"occupation": "developer"}`)))
	must(users.Insert(rex.String(`{"age": 35}`))) // duplicate is not inserted
	fmt.Println(rex.Dump(users))
	// Output:
	// age | name | occupation
	// 35  | ✕    | ✕
	// ✕   | Jake | ✕
	// ✕   | ✕    | developer
}

func ExampleDumpNested() {
	users := rex.R{}
	must(users.Insert(rex.String(`{"name": "Jake", "address": {"city": "New York", "street": "Broadway"}}}`)))
	fmt.Println(rex.Dump(users, rex.Pad("city", 8)))
	// Output:
	// address | name
	// *R1     | Jake
	// -- R1:
	// city     | street
	// New York | Broadway
}

func ExampleEquals() {
	r := rex.R{}
	s := rex.R{}
	must(r.Insert(rex.String(`{"name": "Jake", "address": {"city": "New York", "street": "Broadway"}}`)))
	must(s.Insert(rex.String(`{"name": "Jake", "address": {"street": "Broadway", "city": "New York"}}`)))
	fmt.Println(r.Equals(s))
	// Output: true
}

func ExampleExcept() {
	r := rex.R{}
	s := rex.R{}
	must(r.Insert(rex.String(`{"name": "Jake", "address": {"city": "New York", "street": "Broadway"}}`)))
	must(r.Insert(rex.String(`{"name": "Mia", "address": {"city": "New York", "street": "Broadway"}}`)))
	must(s.Insert(rex.String(`{"name": "Jake", "address": {"street": "Broadway", "city": "New York"}}`)))
	t := r.Except(s)
	fmt.Println(rex.Dump(t))
	// Output:
	// address | name
	// *R1     | Mia
	// -- R1:
	// city | street
	// New York | Broadway
}

func ExampleUnion() {
	r := rex.R{}
	s := rex.R{}
	must(r.Insert(rex.String(`{"name": "Jake", "age": 24}`)))
	must(r.Insert(rex.String(`{"city": "Olomouc"}`)))
	must(s.Insert(rex.String(`{"name": "Aya", "age": 30}`)))
	must(s.Insert(rex.String(`{"city": "Prague"}`)))
	t := r.Union(s)
	fmt.Println(rex.Dump(t, rex.Pad("city", 7)))
	// Output:
	// age | city    | name
	// 24  | ✕       | Jake
	// 30  | ✕       | Aya
	// ✕   | Olomouc | ✕
	// ✕   | Prague  | ✕
}
```
