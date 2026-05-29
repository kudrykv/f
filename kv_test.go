package f_test

import (
	"reflect"
	"testing"

	"github.com/kudrykv/f"
)

func TestKV(t *testing.T) {
	t.Parallel()

	type S struct {
		Name        string
		Description string
	}

	t.Run("kv", func(t *testing.T) {
		t.Parallel()

		list := []S{
			{Name: "A", Description: "Apple"},
			{Name: "B", Description: "Banana"},
			{Name: "C", Description: "Cherry"},
		}

		expected := map[string]S{
			"A": {Name: "A", Description: "Apple"},
			"B": {Name: "B", Description: "Banana"},
			"C": {Name: "C", Description: "Cherry"},
		}

		mapper := func(s S) (string, S) { return s.Name, s }

		actual := f.KV(list, mapper)

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	})
}
