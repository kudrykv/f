package f

import "fmt"

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

func MapErr[In, Out any](list []In, mapper func(In) (Out, error)) ([]Out, error) {
	if len(list) == 0 {
		return nil, nil
	}

	var (
		err    error
		result = make([]Out, len(list))
	)

	for i := range list {
		if result[i], err = mapper(list[i]); err != nil {
			return nil, fmt.Errorf("map %d: %w", i, err)
		}
	}

	return result, err
}
