package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestInsertOne(t *testing.T) {
	t.Run("adding columns", func(t *testing.T) {
		users := rex.Table{}
		users.InsertOne(`{"name": "Jake"}`)
		users.InsertOne(`{"bornYear": 2001}`)
		ac := dump(users.Select())
		ex := dump([][]any{
			{rex.Empty{}, "Jake"},
			{2001, rex.Empty{}},
		})
		if ac != ex {
			t.Errorf("expected\n%vgot\n%v", ex, ac)
		}
	})
}
