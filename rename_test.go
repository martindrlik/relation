package rex_test

import (
	"testing"

	"golang.org/x/exp/maps"
)

func TestRename(t *testing.T) {
	actual := take1(in(map[string]any{"firstName": "Finn", "tvShow": "Adventure Time"}).Rename(map[string]string{"firstName": "name"})).m
	expect := map[string]any{"name": "Finn", "tvShow": "Adventure Time"}
	if !maps.Equal(actual, expect) {
		t.Errorf("%v is not equal to %v", actual, expect)
	}
}
