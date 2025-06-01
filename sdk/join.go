package sdk

func Join[T any](values []T, fn func(idx int, value T) (string, error)) (string, error) {
	t := ""
	for i, item := range values {
		str, err := fn(i, item)
		if err != nil {
			return "", err
		}

		t += str
	}

	return t, nil
}
