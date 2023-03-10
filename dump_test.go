package rex_test

import (
	"fmt"
	"strings"
)

func dump(t ...any) string {
	var sb strings.Builder
	for _, t := range t {
		if sb.Len() > 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(fmt.Sprintf("%T(%v)", t, t))
	}
	return sb.String()
}
