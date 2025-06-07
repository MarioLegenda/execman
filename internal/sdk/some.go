package sdk

func Some[T any](values []T, fn func(idx int, value T) bool) bool {
	for i, item := range values {
		if fn(i, item) {
			return true
		}
	}

	return false
}
