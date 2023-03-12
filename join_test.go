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
	t.Run("simple", func(t *testing.T) {
		result := rex.NaturalJoin(&pub, &pvt)
		view := result.Select(rex.Project("score", "email"))
		ac := dump(view)
		ex := dump([][]any{{2, "foo@example.com"}})
		if ac != ex {
			t.Errorf("expected\n%vgot\n%v", ex, ac)
		}
	})
	t.Run("multiple tables", func(t *testing.T) {
		credits := rex.Table{}
		credits.InsertOne(`{"username": "foo", "credits": 100}`)
		result := rex.NaturalJoin(&pub, &pvt, &credits)
		view := result.Select(rex.Project("score", "email", "credits"))
		ac := dump(view)
		ex := dump([][]any{{2, "foo@example.com", 100}})
		if ac != ex {
			t.Errorf("expected\n%vgot\n%v", ex, ac)
		}
	})
	t.Run("empty intersection", func(t *testing.T) {
		pts := rex.Table{}
		pts.InsertOne(`{"points": 100}`)
		actual := rex.NaturalJoin(&pub, &pts)
		if actual != nil {
			t.Errorf("expected natural join to be nil, got %v", dump(actual.Select()))
		}
	})
}
