package f

func Map[In, Out any](list []In, mapper func(In) Out) []Out {
	if len(list) == 0 {
		return nil
	}

	result := make([]Out, len(list))

	for i := range list {
		result[i] = mapper(list[i])
	}

	return result
}
