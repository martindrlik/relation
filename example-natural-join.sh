#!/bin/zsh

cast='[{"movie_id": "br2049", "actor_id": "rg", "character_id": "br2049k"}]'
movies='[{"movie_id": "br2049", "movie_name": "Blade Runner 2049", "movie_release_year": 2017}]'
actors='[{"actor_id": "rg", "actor_name": "Ryan Gosling"}]'
characters='[{"character_id": "br2049k", "character_name": "K"}]'

result=`go run . natural-join -ia $cast -ia $movies -of json`
result=`go run . natural-join -ia $result -ia $actors -of json`
result=`go run . natural-join -ia $result -ia $characters -of json movie_name actor_name character_name`

echo "Cast:       $cast"
echo "Movies:     $movies"
echo "Actors:     $actors"
echo "Characters: $characters"
echo "Project:    {\"movie_name\", \"actor_name\", \"character_name\"}"
echo "Result:     $result"
