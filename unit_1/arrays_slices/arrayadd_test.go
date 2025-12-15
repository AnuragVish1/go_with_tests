package arraysslices

import (
	"testing"
)

func TestArrayAdd(t *testing.T) {
	t.Run("Addition of the elements in the array", func(t *testing.T) {
		numbers := []int{1, 1, 1, 1, 1}
		got := arrayAddition(numbers)
		want := 5

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

}
