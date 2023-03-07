package table_test

import (
	"fmt"

	"github.com/martindrlik/rex/schema"
	"github.com/martindrlik/rex/table"
	"github.com/martindrlik/rex/tuple"
)

func ExampleTable() {
	foo := map[string]any{"foo": 1}
	bar := map[string]any{"bar": 2.0}
	foobar := tuple.Merge(foo, bar)

	q, _ := table.New(schema.FromTuple(foobar))
	q.Add(foo)
	q.Add(bar)
	q.Add(foobar)

	for t := range q.List() {
		fmt.Println(t)
	}

	// Output:
	// map[foo:1]
	// map[bar:2]
	// map[bar:2 foo:1]
}

func ExampleTable_Add() {
	foo := map[string]any{"foo": 1}
	bar := map[string]any{"bar": 2.0}
	q, _ := table.New(schema.FromTuple(foo))
	fmt.Println(q.Add(bar) == schema.ErrMismatch)

	r, _ := table.New(schema.FromTuple(tuple.Merge(foo, bar)))
	baz := map[string]any{"baz": "3"}
	fmt.Println(r.Add(baz) == schema.ErrMismatch)
	fmt.Println(r.Add(bar) == nil)

	// Output:
	// true
	// true
	// true
}

func ExampleTable_NaturalJoin() {
	foobar := map[string]any{"foo": 1, "bar": 2.0}
	barbaz := map[string]any{"bar": 2.0, "baz": "3"}

	q, _ := table.New(schema.FromTuple(foobar))
	q.Add(foobar)
	r, _ := table.New(schema.FromTuple(barbaz))
	r.Add(barbaz)

	s := q.NaturalJoin(r)

	for t := range s.List() {
		fmt.Println(t)
	}

	// Output:
	// map[bar:2 baz:3 foo:1]
}

func ExampleTable_Projection() {
	foo := map[string]any{"foo": 1}
	foo2 := map[string]any{"foo": 2}
	bar := map[string]any{"bar": 3.0}
	foobar := tuple.Merge(foo, bar)
	foo2bar := tuple.Merge(foo2, bar)

	q, _ := table.New(schema.FromTuple(foobar))
	_ = q.Add(foo)
	_ = q.Add(foo2)
	_ = q.Add(bar)
	_ = q.Add(foobar)
	_ = q.Add(foo2bar)

	s, _ := q.Projection("foo")

	for t := range s.List() {
		fmt.Println(t)
	}

	// Output:
	// map[foo:1]
	// map[foo:2]
}
