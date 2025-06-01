package sdk

func Map[T any, F any](values []T, fn func(idx int, value T) F) []F {
	t := make([]F, 0)

	for i, item := range values {
		t = append(t, fn(i, item))
	}

	return t
}
