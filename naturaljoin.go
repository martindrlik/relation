package rex

// NaturalJoin returns set of all combinations of tuples in left and right relation that are equal on their common attribute names.
func NaturalJoin(left, right *Relation) *Relation {
	r := NewRelation()
	for lk, lr := range *left {
		la := lk.split()
		li := la.indexMap()
		for rk, rr := range *right {
			ra := rk.split()
			ri := ra.indexMap()
			same := map[string]struct{}{}
			for a := range li {
				if _, ok := ri[a]; ok {
					same[a] = struct{}{}
				}
			}
			shouldAdd := func(lt, rt tuple) bool {
				for a := range same {
					if lt[li[a]] != rt[ri[a]] {
						return false
					}
				}
				return true
			}
			for _, lt := range lr {
				for _, rt := range rr {
					if !shouldAdd(lt.tuple, rt.tuple) {
						continue
					}
					k, tx := tupleMap(joinMaps(lt.toMap(la), rt.toMap(ra))).ktx()
					tx.isPossible = lt.isPossible || rt.isPossible || len(same) == 0
					(*r)[k] = append((*r)[k], tx)
				}
			}

		}
	}
	return r
}

func joinMaps[K comparable, V any](a ...map[K]V) map[K]V {
	m := map[K]V{}
	for _, a := range a {
		for k, v := range a {
			m[k] = v
		}
	}
	return m
}
