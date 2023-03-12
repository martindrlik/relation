package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestDelete(t *testing.T) {
	setupUsers := func() *rex.Table {
		users := rex.Table{}
		users.InsertOne(`{"name": "Emma", "bornYear": 1995}`)
		users.InsertOne(`{"name": "Jake", "bornYear": 2001}`)
		users.InsertOne(`{"name": "Mia", "bornYear": 2002}`)
		return &users
	}
	t.Run("all", func(t *testing.T) {
		users := setupUsers()
		if ac, ex := users.Delete(), 3; ac != ex {
			t.Errorf("expected to remove %v, removed %v", ex, ac)
		}
		ac := dump(users.Select())
		ex := dump([][]any{})
		if ac != ex {
			t.Errorf("expected\n%vgot\n%v", ex, ac)
		}
	})
	t.Run("first one", func(t *testing.T) {
		users := setupUsers()
		if ac, ex := users.Delete(rex.Where(`{"bornYear": 1995}`)), 1; ac != ex {
			t.Errorf("expected to remove %v, removed %v", ex, ac)
		}
		ac := dump(users.Select())
		ex := dump([][]any{
			{2002, "Mia"},
			{2001, "Jake"},
		})
		if ac != ex {
			t.Errorf("expected\n%vgot\n%v", ex, ac)
		}
	})
	t.Run("last one", func(t *testing.T) {
		users := setupUsers()
		if ac, ex := users.Delete(rex.Where(`{"name": "Mia"}`)), 1; ac != ex {
			t.Errorf("expected to remove %v, removed %v", ex, ac)
		}
		ac := dump(users.Select())
		ex := dump([][]any{
			{1995, "Emma"},
			{2001, "Jake"},
		})
		if ac != ex {
			t.Errorf("expected\n%vgot\n%v", ex, ac)
		}
	})
}
