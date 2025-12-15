package iterations

import (
	"fmt"
	"testing"
)

func TestIternation(t *testing.T) {
	t.Run("Iterating characters n number of times", func(t *testing.T) {
		got := iteration("2", 10)
		want := "2222222222"

		if got != want {
			t.Errorf("result %q expected %q", got, want)
		}
	})
}

func TestCompare(t *testing.T) {
	t.Run("Comparing two string wheather they are equal or not", func(t *testing.T) {
		got := compareStrings("hello", "hello")
		want := "yes"

		if got != want {
			t.Errorf("result %q expected %q", got, want)
		}
	})
}

func BenchmarkIteration(b *testing.B) {
	for b.Loop() {
		iteration("h", 5)
	}
}

func Exampleiteration() {
	fmt.Println(iteration("a", 5))
	// Output: aaaaa
}
