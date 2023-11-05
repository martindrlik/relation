package rex_test

import (
	"testing"
)

func TestContains(t *testing.T) {
	finnAndJake := in(finn, jake)
	if !finnAndJake.Contains(finn) {
		t.Errorf("expected %v to contain %v", finnAndJake, finn)
	}
	if finnAndJake.Contains(marceline) {
		t.Errorf("expected %v not to contain %v", finnAndJake, marceline)
	}
}
