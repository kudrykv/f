package f_test

import (
	"f"
	"slices"
	"testing"
)

func TestFilter(t *testing.T) {
	t.Parallel()

	t.Run("filter", func(t *testing.T) {
		t.Parallel()

		list := []int{1, 2, 3, 4, 5}
		keep := func(i int) bool { return i%2 == 0 }
		expected := []int{2, 4}

		actual := f.Filter(list, keep)

		if !slices.Equal(expected, actual) {
			t.Fatalf("Filter(%v) = %v, want %v", list, actual, expected)
		}
	})

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		var emptyList []int

		keep := func(i int) bool { return i%2 == 0 }

		actual := f.Filter(emptyList, keep)

		if !slices.Equal(emptyList, actual) {
			t.Fatalf("Filter(%v) = %v, want %v", emptyList, actual, emptyList)
		}

		if actual != nil {
			t.Fatalf("Filter(%v) = %v, want %v", emptyList, actual, []int(nil))
		}
	})
}
