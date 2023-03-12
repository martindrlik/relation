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
			{"Emma", 1995},
			{"Jake", 2001},
			{"Mia", 2002},
		})
		if ac != ex {
			t.Errorf("expected\n%vgot\n%v", ex, ac)
		}
	})
	t.Run("select by bornYear", func(t *testing.T) {
		users := setupUsers()
		ac := dump(users.Select(rex.Where(`{"bornYear": 2001}`)))
		ex := dump([][]any{
			{"Jake", 2001},
		})
		if ac != ex {
			t.Errorf("expected\n%vgot\n%v", ex, ac)
		}
	})
	//	t.Run("insert two", func(t *testing.T) {
	//		ta := rex.Table{}
	//		ta.InsertOne(`{"name": "Foo"}`).InsertOne(`{"value": 1990}`)
	//		ac := dump(ta.SelectRange(0, "name", "value")...)
	//		ex := dump(rex.Field{"name", "Foo"}, rex.Field{"value", rex.Empty{}})
	//		if ac != ex {
	//			t.Errorf("expected first tuple %v, got %v", ex, ac)
	//		}
	//		ac = dump(ta.SelectRange(1, "name", "value"))
	//		ex = dump(rex.Field{"name", rex.Empty{}}, rex.Field{"value", float64(1990)})
	//		if ac != ex {
	//			t.Errorf("expected second tuple %v, got %v", ex, ac)
	//		}
	//	})
	//
	//	t.Run("insert two with same column", func(t *testing.T) {
	//		ta := rex.Table{}
	//		ta.InsertOne(`{"name": "Foo"}`).InsertOne(`{"name": "Bar", "value": 1990}`)
	//		ac := dump(ta.SelectRange(0, "name", "value")...)
	//		ex := dump(rex.Field{"name", "Foo"}, rex.Field{"value", rex.Empty{}})
	//		if ac != ex {
	//			t.Errorf("expected first tuple %v, got %v", ex, ac)
	//		}
	//		ac = dump(ta.SelectRange(1, "name", "value")...)
	//		ex = dump(rex.Field{"name", "Bar"}, rex.Field{"value", float64(1990)})
	//		if ac != ex {
	//			t.Errorf("expected second tuple %v, got %v", ex, ac)
	//		}
	//	})
}

func TestSelectRange(t *testing.T) {
	users := rex.Table{}
	users.InsertOne(`{"name": "Emma", "bornYear": 1995}`)
	users.InsertOne(`{"name": "Jake", "bornYear": 2001}`)
	users.InsertOne(`{"name": "Mia", "bornYear": 2002}`)
	x := users.SelectRange(0, 1, "name", "bornYear")
	ac := dump(x)
	ex := dump([][]any{
		{"Emma", 1995},
		{"Jake", 2001},
	})
	if ac != ex {
		t.Errorf("expected\n%v\ngot\n%v", ex, ac)
	}
}
