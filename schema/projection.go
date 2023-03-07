package schema

// Projection returns a new schema with the keys of p that are in s.
func (s Schema) Projection(p ...string) (Schema, bool) {
	w := make(Schema)
	for _, a := range p {
		if t, ok := s[a]; ok {
			w[a] = t
		} else {
			return nil, false
		}
	}
	return w, true
}
