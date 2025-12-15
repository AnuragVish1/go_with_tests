package arraysslices

import (
	"reflect"
	"slices"
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

	t.Run("Addition of multiple slices", func(t *testing.T) {
		got := sumAll([]int{1, 1}, []int{1, 1})
		want := []int{2, 2}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}

func TestArrayTails(t *testing.T) {

	checksum := func(got []int, want []int, t *testing.T) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Addition of tails only", func(t *testing.T) {
		got := sumAllTails([]int{0, 1}, []int{2, 3})
		want := []int{1, 3}
		checksum(got, want, t)
	})

	t.Run("Checking if empty slice is taken into consideration", func(t *testing.T) {
		got := sumAllTails([]int{}, []int{1, 2})
		want := []int{0, 2}
		checksum(got, want, t)
	})
}
