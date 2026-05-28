package f

func GroupBy[E ~[]S, S any, K comparable](list E, f func(S) K) map[K]E {
	grouped := make(map[K]E)

	for _, item := range list {
		key := f(item)
		grouped[key] = append(grouped[key], item)
	}

	return grouped
}
