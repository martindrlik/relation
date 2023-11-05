# rex

Experimental relational NoSQL database. It is my playground for ideas and API will change over time. There is a lot more to do before it can be even considered interesting.

## Example

``` go
func Example() {
	shows := rex.NewRelation().Insert(map[string]any{"show": "Adventure Time"})
	characters := rex.NewRelation().Insert(
		map[string]any{"name": "Finn"},
		map[string]any{"name": "Marceline"})
	adventure := rex.NaturalJoin(shows, characters)
	adventure.Each(func(t map[string]any) error {
		fmt.Println(t)
		return nil
	})
	// Output:
	// map[name:Finn show:Adventure Time]
	// map[name:Marceline show:Adventure Time]
}
```