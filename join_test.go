package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestNaturalJoin(t *testing.T) {
	public := rex.Table{}
	private := rex.Table{}
	public.InsertOne(`{"username": "foo", "score": 2}`)
	private.InsertOne(`{"username": "foo", "email": "foo@example.com"}`)
	rows := public.NaturalJoin(&private).Select(rex.Project("score", "email"))
	ac := dump(rows)
	ex := dump([][]any{{2, "foo@example.com"}})
	if ac != ex {
		t.Errorf("expected\n%vgot\n%v", ex, ac)
	}
}
