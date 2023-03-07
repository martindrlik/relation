# rex

Experimental relational NoSQL database. Mainly touching upon the incomplete data problem. It is my playground for ideas and API will change over time. There is a lot more to do before it can be even considered interesting.

## Program usage
```
Usage:
	rex <command> <input> <options> [attribute ...]
Commands:
	difference
	natural-join
	union
Input:
	-fa <file>   [-ta <file>   ...]: name of file that contains array of tuples
	-ia <inline> [-ia <inline> ...]: inline array of tuples
Options:
	-of <format>: output format: table or json
Note:
	JSON is used as an input format
```

## Examples
```
#!/bin/zsh

movies='[{"mno": 1, "title": "Dune: Part Two", "year": 2024}]'
timothee_paul='{"mno": 1, "actor": "Timothée Chalamet", "character": "Paul Atreides"}'
zendaya_chani='{"mno": 1, "actor": "Zendaya", "character": "Chani"}'
rebecca_='{"mno": 1, "actor": "Rebecca Ferguson"}'
cast="[$timothee_paul, $zendaya_chani, $rebecca_]"

echo "Movies:"
go run . union -ia $movies

echo "Cast:"
go run . union -ia $cast

echo "movies.NaturalJoin(cast).Projection: title actor character"
go run . natural-join -ia $movies -ia $cast title actor character
```