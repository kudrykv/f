package f_test

import (
	"slices"
	"strconv"
	"testing"

	"github.com/kudrykv/f"
)

func TestMap(t *testing.T) {
	t.Parallel()

	t.Run("non-empty list", func(t *testing.T) {
		t.Parallel()

		list := []int{1, 2, 3}
		expected := []string{"2", "4", "6"}
		mapper := func(value int) string { return strconv.Itoa(value * 2) }

		actual := f.Map(list, mapper)

		if !slices.Equal(expected, actual) {
			t.Fatalf("Map(%v) = %v, want %v", list, actual, expected)
		}
	})

	t.Run("empty list", func(t *testing.T) {
		t.Parallel()

		var list []int

		mapper := func(value int) int { return value * 2 }

		result := f.Map(list, mapper)

		if len(result) != 0 {
			t.Fatalf("Map(%v) = %v, want %v", list, result, []string(nil))
		}

		if result != nil {
			t.Fatalf("Map(%v) = %v, want %v", list, result, []string(nil))
		}
	})
}
