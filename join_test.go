package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestNaturalJoin(t *testing.T) {
	x := rex.Table{}
	y := rex.Table{}
	x.InsertOne(`{"a": 5, "b": 6}`).NaturalJoin(y.InsertOne(`{"a": 5, "b": 6, "c": 11}`))
}
