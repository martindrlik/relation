package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestProject(t *testing.T) {
	users := rex.Table{}
	users.InsertOne(`{"name": "Martin", "age": 39}`)
	p := users.Project("age")
	actual := dump(p.At(0)...)
	expect := dump(float64(39))
	if actual != expect {
		t.Errorf("expected tuple %v, got %v", expect, actual)
	}
}
