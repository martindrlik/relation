package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestSelect(t *testing.T) {
	setupUsers := func() *rex.Table {
		users := rex.Table{}
		users.InsertOne(`{"name": "Emma", "bornYear": 1995}`)
		users.InsertOne(`{"name": "Jake", "bornYear": 2001}`)
		users.InsertOne(`{"name": "Mia", "bornYear": 2002}`)
		return &users
	}
	t.Run("all", func(t *testing.T) {
		users := setupUsers()
		ac := dump(users.Select())
		ex := dump([][]any{
			{1995, "Emma"},
			{2001, "Jake"},
			{2002, "Mia"},
		})
		if ac != ex {
			t.Errorf("expected\n%vgot\n%v", ex, ac)
		}
	})
	t.Run("project", func(t *testing.T) {
		users := setupUsers()
		ac := dump(users.Select(rex.Project("name", "bornYear")))
		ex := dump([][]any{
			{"Emma", 1995},
			{"Jake", 2001},
			{"Mia", 2002},
		})
		if ac != ex {
			t.Errorf("expected\n%vgot\n%v", ex, ac)
		}
	})
	t.Run("by born year", func(t *testing.T) {
		users := setupUsers()
		ac := dump(users.Select(rex.Where(`{"bornYear": 2001}`)))
		ex := dump([][]any{
			{2001, "Jake"},
		})
		if ac != ex {
			t.Errorf("expected\n%vgot\n%v", ex, ac)
		}
	})
}
