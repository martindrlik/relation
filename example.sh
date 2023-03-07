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
