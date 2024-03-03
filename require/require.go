package require

// NoError panics if err is not nil otherwise returns t.
func NoError[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func NilError(err error) {
	if err != nil {
		panic(err)
	}
}
