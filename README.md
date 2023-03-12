# rex

Experimental relational NoSQL database. It is my playground for ideas and API will change over time. There is a lot more to do before it can be even considered interesting.

``` golang
var users, score rex.Table
users.InsertOne(`{"username": "Jake"}`)
score.InsertOne(`{"username": "Jake", "score": 100}`)
dump(rex.NaturalJoin(&users, &score).
	Select(
		rex.Where(`{"score": 100}`), // restriction
		rex.Project("username")))    // projection
// Output:
// Jake
```
