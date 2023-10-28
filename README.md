# rex

Experimental relational NoSQL database. It is my playground for ideas and API will change over time. There is a lot more to do before it can be even considered interesting.

## Example

``` go
shows := rex.NewRelation().Insert(map[string]any{"show": "Adventure Time"})
characters := rex.NewRelation().Insert(
	map[string]any{"name": "Finn"},
	map[string]any{"name": "Marceline"})
adventure := rex.NaturalJoin(shows, characters)
adventure.Each(func(tm map[string]any, isPossible bool) bool {
	fmt.Println(isPossible, tm)
	return true
})
// Output:
// true map[name:Finn show:Adventure Time]
// true map[name:Marceline show:Adventure Time]
```