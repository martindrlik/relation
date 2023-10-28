package rex_test

import (
	"testing"

	"github.com/martindrlik/rex"
)

func TestNaturalJoin(t *testing.T) {
	t.Run("cartesian product", func(t *testing.T) {
		adventure := rex.NaturalJoin(
			in(glue(finnName), glue(jakeName)),
			in(glue(adventureRelease)))
		for _, m := range []map[string]any{finn, jake} {
			if !adventure.Contains(m) {
				t.Errorf("expected %v to contain %v", adventure, m)
			}
		}
		finn, jake := take2(adventure)
		if !finn.isPossible {
			t.Errorf("expected %v to be possible meaning not directly inserted", finn)
		}
		if !jake.isPossible {
			t.Errorf("expected %v to be possible meaning not directly inserted", jake)
		}
	})
	t.Run("department", func(t *testing.T) {
		employee := func(name, department string) map[string]any {
			return map[string]any{"name": name, "department": department}
		}
		department := func(name, manager string) map[string]any {
			return map[string]any{"department": name, "manager": manager}
		}
		var (
			harry   = employee("Harry", "Finance")
			sally   = employee("Sally", "Sales")
			george  = employee("George", "Finance")
			harriet = employee("Harriet", "Sales")

			finance    = department("Finance", "George")
			sales      = department("Sales", "Harriet")
			production = department("Production", "Charles")
		)
		adventure := rex.NaturalJoin(
			in(harry, sally, george, harriet),
			in(finance, sales, production))
		for _, m := range []map[string]any{
			gluem(harry, finance),
			gluem(sally, sales),
			gluem(george, finance),
			gluem(harriet, sales),
		} {
			if !adventure.Contains(m) {
				t.Errorf("expected %v to contain %v", adventure, m)
			}
		}
	})
}
