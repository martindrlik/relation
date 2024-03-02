package relation

import "errors"

var (
	ErrMissingSchema  = errors.New("missing schema")
	ErrSchemaMismatch = errors.New("schema mismatch")
)
