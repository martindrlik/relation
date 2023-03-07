package persist

import (
	"encoding/json"
	"io"
	"slices"

	"github.com/martindrlik/rex/table"
)

func WriteJson(w io.Writer, t *table.Table) error {
	enc := json.NewEncoder(w)
	return enc.Encode(slices.Collect(t.List()))
}
