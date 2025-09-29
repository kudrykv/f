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

func FilterInPlace[S ~[]E, E any](list S, keep func(E) bool) S {
	index := 0

	for i := range list {
		if keep(list[i]) {
			list[index] = list[i]
			index++
		}
	}

	return list[:index]
}
