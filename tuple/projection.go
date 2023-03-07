package tuple

func (t Tuple) Projection(p ...string) Tuple {
	w := make(Tuple, len(p))
	for _, a := range p {
		w[a] = t[a]
	}
	return w
}
