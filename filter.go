package f

func Filter[S ~[]E, E any](list S, f func(E) bool) S {
	var filtered S

	for i := range list {
		if !f(list[i]) {
			continue
		}

		filtered = append(filtered, list[i])
	}

	return filtered
}
