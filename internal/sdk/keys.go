package sdk

func Keys[K comparable, F any](data map[K]F) []K {
	keys := make([]K, 0)
	for key, _ := range data {
		keys = append(keys, key)
	}

	return keys
}
