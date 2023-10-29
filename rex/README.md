# rex

Simple CLI for rex package.

## Example

First start rex server.

``` zsh
% rex --server
```

Then in another terminal you can use rex CLI to insert and list tuples.

``` shell
% rex insert --json '{"name": "Finn the Human"}'
% rex insert --json '{"name": "Jake the Dog"}'
% rex insert --json '{"name": "Marceline the Vampire Queen"}'
% rex list
map[name:Finn the Human]
map[name:Jake the Dog]
map[name:Marceline the Vampire Queen]
```