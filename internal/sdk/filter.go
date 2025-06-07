package sdk

func Filter[T any](values []T, fn func(idx int, value T) bool) []T {
	t := make([]T, 0)

	for i, item := range values {
		if fn(i, item) {
			t = append(t, item)
		}
	}

	return t
}
