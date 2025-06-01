package sdk

func Includes[T comparable](values []T, search T) bool {
	for _, val := range values {
		if val == search {
			return true
		}
	}

	return false
}
