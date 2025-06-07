package sdk

func Every[T any](values []T, fn func(idx int, value T) bool) bool {
	for i, item := range values {
		if !fn(i, item) {
			return false
		}
	}

	return true
}
