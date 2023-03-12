package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestTable(t *testing.T) {
	setupUsers := func() *rex.Table {
		users := rex.Table{}
		users.InsertOne(`{"name": "Emma", "bornYear": 1995}`)
		users.InsertOne(`{"name": "Jake", "bornYear": 2001}`)
		users.InsertOne(`{"name": "Mia", "bornYear": 2002}`)
		return &users
	}
	t.Run("select all", func(t *testing.T) {
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
	t.Run("insert another column", func(t *testing.T) {
		users := setupUsers()
		users.InsertOne(`{"bornMonth": "April"}`)
		ac := dump(users.Select())
		ex := dump([][]any{
			{rex.Empty{}, 1995, "Emma"},
			{rex.Empty{}, 2001, "Jake"},
			{rex.Empty{}, 2002, "Mia"},
			{"April", rex.Empty{}, rex.Empty{}},
		})
		if ac != ex {
			t.Errorf("expected\n%vgot\n%v", ex, ac)
		}
	})
	t.Run("select all projection", func(t *testing.T) {
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
	t.Run("select by born year", func(t *testing.T) {
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
