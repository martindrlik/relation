package rex

// NaturalJoin returns set of all combinations of tuples in left and right relation that are equal on their common attribute names.
func NaturalJoin(left, right *Relation) *Relation {
	r := NewRelation()

	left.Each(func(lt map[string]any) error {
		right.Each(func(rt map[string]any) error {
			t, ok := Tuple(lt).NaturalJoin(rt)
			if ok {
				r.Insert(t)
			}
			return nil
		})
		return nil
	})
	return r
}

func (t Tuple) NaturalJoin(u Tuple) (Tuple, bool) {
	ta := t.attrs()
	ua := u.attrs()
	common := ta.intersection(ua)
	shouldAdd := func() bool {
		if len(common) > 0 {
			for a := range common {
				if !equal(t[a], u[a]) {
					return false
				}
			}
		}
		return true
	}()
	if shouldAdd {
		return t.combine(u), true
	}
	return nil, false
}
