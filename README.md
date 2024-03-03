# rex

Experimental relational NoSQL database. It is my playground for ideas and API will change over time. There is a lot more to do before it can be even considered interesting.

## Example

``` go
func Example() {

	t := require.NoError(table.New("name", "age"))
	t.Append(tuple.T{"name": "John", "age": 42})

	v := require.NoError(table.New("name", "age"))
	v.Append(tuple.T{"name": "John", "age": 42})
	v.Append(tuple.T{"name": "Jake"})

	w := require.NoError(t.Union(v))
	fmt.Print(box.Table(w))

	// Output:
	// ┏━━━━━━┯━━━━━┓
	// ┃ name │ age ┃
	// ┠──────┼─────┨
	// ┃ John │ 42  ┃
	// ┃ Jake │ *   ┃
	// ┗━━━━━━┷━━━━━┛

}
```