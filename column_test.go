package rex

// import (
// 	"testing"

// 	"github.com/martindrlik/rex"
// )

// func TestColumn(t *testing.T) {
// 	t.Run("insert one", func(t *testing.T) {
// 		co := rex.Column{}
// 		co.Insert("foo")
// 		if ln := co.Len(); ln != 1 {
// 			t.Errorf("expected length to be 1 got %d", ln)
// 		}
// 		if f, ok := co.At(0); !ok || f.Value != "foo" {
// 			t.Errorf("expected \"foo\" got %v %v", f.Value, ok)
// 		}
// 	})
// 	t.Run("at out of range", func(t *testing.T) {
// 		co := rex.Column{}
// 		co.Insert("foo")
// 		if _, ok := co.At(-1); ok {
// 			t.Error("expected no value at -1")
// 		}
// 		if _, ok := co.At(1); ok {
// 			t.Error("expected no value at 1")
// 		}
// 	})
// 	t.Run("remove last one of two", func(t *testing.T) {
// 		co := rex.Column{}
// 		co.Insert("foo")
// 		co.Insert("bar")
// 		co.RemoveAt(1)
// 		if ln := co.Len(); ln != 1 {
// 			t.Errorf("expected length to be 1 got %d", ln)
// 		}
// 		if f, ok := co.At(0); !ok || f.Value != "foo" {
// 			t.Errorf("expected last \"foo\" move to first but got %v %v", f.Value, ok)
// 		}
// 	})
// 	t.Run("remove first one of two", func(t *testing.T) {
// 		co := rex.Column{}
// 		co.Insert("foo")
// 		co.Insert("bar")
// 		co.RemoveAt(0)
// 		if ln := co.Len(); ln != 1 {
// 			t.Errorf("expected length to be 1 got %d", ln)
// 		}
// 		if f, ok := co.At(0); !ok || f.Value != "bar" {
// 			t.Errorf("expected last \"bar\" move to first but got %v %v", f.Value, ok)
// 		}
// 	})
// 	t.Run("remove last", func(t *testing.T) {
// 		co := rex.Column{}
// 		co.Insert("foo")
// 		co.RemoveAt(0)
// 		if ln := co.Len(); ln != 0 {
// 			t.Errorf("expected length to be 0 got %d", ln)
// 		}
// 	})
// 	t.Run("remove out of range", func(t *testing.T) {
// 		co := rex.Column{}
// 		co.Insert("foo")
// 		co.RemoveAt(-1)
// 		co.RemoveAt(1)
// 		if f, ok := co.At(0); !ok || f.Value != "foo" {
// 			t.Errorf("expected \"foo\" got %v %v", f.Value, ok)
// 		}
// 	})
// }
