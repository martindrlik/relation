package main

import "fmt"

func main() {
	user := Table{}
	fail(user.Append(`{"name": "Martin", "age": 39}`))
	fail(user.Append(`{"name": "Martin", "age": "x"}`))
	row := must(user.Row(1))
	for _, v := range row {
		fmt.Println(v)
	}
}
