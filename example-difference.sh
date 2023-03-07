#!/bin/zsh

table_one='[{"year": 2049}, {"year": 2050}]'
table_two='[{"year": 2050}]'
result=`go run . difference -ia $table_one -ia $table_two -of json`

echo "Table 1: $table_one"
echo "Table 2: $table_two"
echo "Result:  $result"
