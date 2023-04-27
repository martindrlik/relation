# rex command

## Example

```
% cat birthday.json 
[{"name":"Jake","born":{"year":1980}},
 {"name":"Kristen","born":{"year":1990}}]
% cat names.json 
[{"name":"Jake"},{"name":"Kristen"}]
% rex -naturaljoin birthday.json names.json 
[{"born": {"year": 1980}, "name": "Jake"},
{"born": {"year": 1990}, "name": "Kristen"}]
%
```