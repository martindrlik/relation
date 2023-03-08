package store_test

func nerr(err error) {
	if err != nil {
		panic(err)
	}
}

func must[T any](t T, err error) T {
	nerr(err)
	return t
}
