package require

// Must panics if err is not nil otherwise returns t.
func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
