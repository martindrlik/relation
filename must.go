package store

func nerr(err error) {
	if err != nil {
		panic(err)
	}
}

func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
