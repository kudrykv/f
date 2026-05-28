package f_test

import (
	"reflect"
	"testing"

	"github.com/kudrykv/f"
)

func TestGroupBy(t *testing.T) {
	t.Parallel()

	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	matcher := func(i int) string {
		if i%2 == 0 {
			return "even"
		}

		return "odd"
	}

	expected := map[string][]int{
		"even": {2, 4, 6, 8},
		"odd":  {1, 3, 5, 7, 9},
	}

	result := f.GroupBy(list, matcher)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("GroupBy(%v) = %v, want %v", list, result, expected)
	}
}
