package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestProject(t *testing.T) {
	users := rex.Table{}
	users.InsertOne(`{"name": "Martin", "age": 39}`)
	p := users.Project("age")
	tuple, ok := p.At(0)
	actual := dump(tuple...)
	expect := dump(float64(39))
	if !ok || actual != expect {
		t.Errorf("expected tuple %v and ok to be true, got %v and %v", expect, actual, ok)
	}
}
