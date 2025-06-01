package sdk

func Delete[F any](values map[any]F, fn func(key any, value F) bool) map[any]F {
	deleted := make(map[any]F)
	for key, value := range values {
		toDelete := fn(key, value)

		if toDelete {
			delete(values, key)
			deleted[key] = value
		}
	}

	return deleted
}
