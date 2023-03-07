package schema

// Projection returns a new schema with the keys of p that are in s.
func (s Schema) Projection(p ...string) Schema {
	w := make(Schema)
	for _, a := range p {
		if t, ok := s[a]; ok {
			w[a] = t
		}
	}
	return w
}
