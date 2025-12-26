package sync

import (
	"sync"
	"testing"
)

func NewCounter() *Counter {
	return &Counter{}
}

func TestCounter(t *testing.T) {
	t.Run("Testing counter", func(t *testing.T) {
		counter := NewCounter()

		counter.Inc()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		want := 4

		validateCounter(t, counter, want)

	})

	t.Run("It runs on a concurrent machine", func(t *testing.T) {
		want := 1000

		counter := NewCounter()

		var wg sync.WaitGroup

		wg.Add(want)

		for range want {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		validateCounter(t, counter, want)
	})
}

func validateCounter(t *testing.T, counter *Counter, want int) {
	t.Helper()
	got := counter.Val()
	if got != want {
		t.Errorf("Got %v want %v", got, want)
	}
}
