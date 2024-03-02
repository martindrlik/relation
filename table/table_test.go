package table_test

import (
	"testing"
)

func TestTable(t *testing.T) {
	// t.Run("Equal", func(t *testing.T) {
	// 	t1 := rex.NewTable("name", "age").Append(tuple.Tuple{"name": "John", "age": 42})
	// 	t2 := rex.NewTable("name", "age").Append(tuple.Tuple{"name": "John", "age": 42})
	// 	if !t1.Equal(t2) {
	// 		t.Error("t1 and t2 should be equal")
	// 	}
	// })

	// t.Run("NotEqual", func(t *testing.T) {
	// 	t1 := rex.NewTable("name", "age").Append(tuple.Tuple{"name": "John", "age": 42})
	// 	t2 := rex.NewTable("name", "age").Append(tuple.Tuple{"name": "Jake", "age": 34})
	// 	t3 := rex.NewTable("name", "age", "city").Append(tuple.Tuple{"name": "John", "age": 42, "city": "London"})
	// 	t4 := rex.NewTable("name", "age").Append(tuple.Tuple{"name": "John", "age": 42}).Append(tuple.Tuple{"name": "Jake", "age": 34})
	// 	if t1.Equal(t2) {
	// 		t.Error("t1 and t2 should not be equal")
	// 	}
	// 	if t1.Equal(t3) {
	// 		t.Error("t1 and t3 should not be equal (different schema)")
	// 	}
	// 	if t1.Equal(t4) {
	// 		t.Error("t1 and t4 should not be equal (t4 has more tuples)")
	// 	}
	// })

	// t.Run("SchemaInOrder", func(t *testing.T) {
	// 	t1 := rex.NewTable("name", "age")
	// 	if !slices.Equal(t1.Schema(), []string{"name", "age"}) {
	// 		t.Errorf("t1 schema is expected to be %v got %v", []string{"name", "age"}, t1.SchemaSet())
	// 	}
	// })

	// t.Run("Union", func(t *testing.T) {
	// 	u1 := rex.NewTable("name", "age").Append(tuple.Tuple{"name": "John", "age": 42})
	// 	u2 := rex.NewTable("name", "age").Append(tuple.Tuple{"name": "Jake"})
	// 	expect := rex.NewTable("name", "age").Append(tuple.Tuple{"name": "John", "age": 42}).Append(tuple.Tuple{"name": "Jake"})
	// 	actual := require.Must(rex.Union(u1, u2))
	// 	if !actual.Equal(expect) {
	// 		t.Error("u1 union u2 should be equal to u12")
	// 	}
	// })
}
