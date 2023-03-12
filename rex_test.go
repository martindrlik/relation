package rex_test

import (
	"encoding/json"
	"strings"
)

func dump(rows [][]any) string {
	sb := &strings.Builder{}
	enc := json.NewEncoder(sb)
	err := enc.Encode(rows)
	if err != nil {
		panic(err)
	}
	return sb.String()
}
