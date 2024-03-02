package relation

func must[T any](t T, err error) T {
	// when expecting no error and getting one will panic
	if err != nil {
		panic(err)
	}
	return t
}
