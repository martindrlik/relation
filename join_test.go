package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestNaturalJoin(t *testing.T) {
	x := rex.Table{}
	y := rex.Table{}
	x.InsertOne(`{"a": 5, "b": 6}`)
	y.InsertOne(`{"a": 5, "b": 6, "c": 11}`)
	actual := dump(x.NaturalJoin(&y).At(0)...)
	expect := dump(float64(5), float64(6), float64(11))
	if actual != expect {
		t.Errorf("expected tuple %v, got %v", expect, actual)
	}
}
