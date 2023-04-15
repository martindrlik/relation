package rex

type tuple map[string]any

func (t tuple) shallowCopy() tuple {
	v := map[string]any{}
	for k, w := range t {
		v[k] = w
	}
	return v
}
