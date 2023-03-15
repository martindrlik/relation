# rex

Experimental relational NoSQL database. It is my playground for ideas and API will change over time. There is a lot more to do before it can be even considered interesting.

``` golang
users := rex.R{}
users.InsertOne(strings.NewReader(`{"name": "Jake"}`))
users.InsertOne(strings.NewReader(`{"age": 35}`))
users.InsertOne(strings.NewReader(`{"occupation": "developer"}`))
users.InsertOne(strings.NewReader(`{"age": 35}`)) // duplicate is not inserted
fmt.Println(rex.Dump(users, 3, 4, 10))
// Output:
// age | name | occupation
// ✕   | Jake | ✕
// 35  | ✕    | ✕
// ✕   | ✕    | developer
```
