package rex_test

import (
	"strings"
	"testing"

	"github.com/martindrlik/rex"
)

func TestInsertOneJson(t *testing.T) {
	r := rex.NewRelation().InsertOneJson(strings.NewReader(`{"name": "Jake", "bornYear": 1980}`))
	expected := rex.NewRelation().InsertOne(name("Jake"), bornYear(1980))
	if !expected.Equals(r) {
		t.Error("expected equal after insert one json")
	}
}

func TestInsertManyJson(t *testing.T) {
	r := rex.NewRelation().InsertManyJson(strings.NewReader(`[
		{"name": "Jake", "bornYear": 1980},
		{"name": "Kristen", "bornYear": 1990}]`))
	expected := rex.NewRelation().
		InsertOne(name("Jake"), bornYear(1980)).
		InsertOne(name("Kristen"), bornYear(1990))
	if !expected.Equals(r) {
		t.Error("expected equal after insert many json")
	}
}

func TestInsertOneJsonNested(t *testing.T) {
	r := rex.NewRelation().InsertOneJson(strings.NewReader(`{"name": "Jake", "born": {"year": 1980}}`))
	y := rex.NewRelation().InsertOneJson(strings.NewReader(`{"year": 1980}`))
	expected := rex.NewRelation().InsertOne(name("Jake"), func() (string, any) { return "born", y })
	if !expected.Equals(r) {
		t.Error("expected equal after insert one json nested")
	}
}
