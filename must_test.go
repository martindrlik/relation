package rex_test

func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
