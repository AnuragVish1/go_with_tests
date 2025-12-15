package integers

import (
	"fmt"
	"testing"
)

func TestAddition(t *testing.T) {
	t.Run("Addition of integers", func(t *testing.T) {
		got := addition(2, 2)
		want := 4

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func Exampleaddition() {
	sum := addition(9, 6)
	fmt.Println(sum)
	// Output: 15
}
