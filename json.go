package rex

import (
	"encoding/json"
	"io"
)

func mustDecode[T any](r io.Reader) (t T) {
	dec := json.NewDecoder(r)
	err := dec.Decode(&t)
	if err != nil {
		panic(err)
	}
	return t
}
