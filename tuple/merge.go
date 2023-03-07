package tuple

func Merge[V any](u, v map[string]V) map[string]V {
	w := make(map[string]V)
	for k, v := range u {
		w[k] = v
	}
	for k, v := range v {
		w[k] = v
	}
	return w
}
