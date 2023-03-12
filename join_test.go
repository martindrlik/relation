package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestNaturalJoin(t *testing.T) {
	pub := rex.Table{}
	pvt := rex.Table{}
	pub.InsertOne(`{"username": "foo", "score": 2}`)
	pvt.InsertOne(`{"username": "foo", "email": "foo@example.com"}`)
	result := rex.NaturalJoin(&pub, &pvt)
	view := result.Select(rex.Project("score", "email"))
	ac := dump(view)
	ex := dump([][]any{{2, "foo@example.com"}})
	if ac != ex {
		t.Errorf("expected\n%vgot\n%v", ex, ac)
	}
}
