package rex_test

import (
	"encoding/json"
	"strings"
)

func dump(fs [][]any) string {
	sb := &strings.Builder{}
	enc := json.NewEncoder(sb)
	err := enc.Encode(fs)
	if err != nil {
		panic(err)
	}
	return sb.String()
}
