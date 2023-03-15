# rex

Experimental relational NoSQL database. It is my playground for ideas and API will change over time. There is a lot more to do before it can be even considered interesting.

## Examples

```golang
r := rex.R{}
s := rex.R{}
r.InsertOne(strings.NewReader(`{"name": "Jake", "age": 24}`))
r.InsertOne(strings.NewReader(`{"city": "Olomouc"}`))
s.InsertOne(strings.NewReader(`{"name": "Aya", "age": 30}`))
s.InsertOne(strings.NewReader(`{"city": "Prague"}`))
t := r.Union(s)
fmt.Println(rex.Dump(t, 3, 7, 4))
// Output:
// age | city    | name
// 24  | ✕       | Jake
// 30  | ✕       | Aya
// ✕   | Olomouc | ✕
// ✕   | Prague  | ✕
```

``` golang
users := rex.R{}
users.InsertOne(strings.NewReader(`{"name": "Jake"}`))
users.InsertOne(strings.NewReader(`{"age": 35}`))
users.InsertOne(strings.NewReader(`{"occupation": "developer"}`))
users.InsertOne(strings.NewReader(`{"age": 35}`)) // duplicate is not inserted
fmt.Println(rex.Dump(users, 3, 4, 10))
// Output:
// age | name | occupation
// 35  | ✕    | ✕
// ✕   | Jake | ✕
// ✕   | ✕    | developer
```
