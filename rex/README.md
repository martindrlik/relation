# rex

Simple CLI for rex package.

## Example

First start rex server.

``` zsh
% rex server
```

Then in another terminal you can use rex CLI to insert and list tuples.

``` shell
% rex insert -target x -json '{"name": "Finn the Human"}'
% rex insert -target x -json '{"name": "Jake the Dog"}'
% rex insert -target x -json '{"name": "Marceline the Vampire Queen"}'
% rex insert -target y -json '{"show": "Adventure Time"}'
% rex natural-join -target z -left x -right y
% rex list -target z
map[name:Finn the Human show:Adventure Time]*
map[name:Jake the Dog show:Adventure Time]*
map[name:Marceline the Vampire Queen show:Adventure Time]*
```