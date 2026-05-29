package f

func KV[E ~[]S, S any, K comparable, V any](list E, mapper func(S) (K, V)) map[K]V {
	if len(list) == 0 {
		return nil
	}

	result := make(map[K]V, len(list))

	for _, item := range list {
		k, v := mapper(item)
		result[k] = v
	}

	return result
}
