package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

var (
	write = "write"
	sleep = "sleep"
)

const (
	countDownLimit = 3
	lastMessage    = "Go"
)

type Sleeper interface {
	Sleep()
}

type SpyCaller struct {
	Calls int
}

func (s *SpyCaller) Sleep() {
	s.Calls++
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type FakeTime struct {
	durationSlept time.Duration
}

func (f *FakeTime) SetDurationSlept(duration time.Duration) {
	f.durationSlept = duration
}

type FakeCountDownOperation struct {
	Calls []string
}

func (f *FakeCountDownOperation) Sleep() {
	f.Calls = append(f.Calls, sleep)
}

func (f *FakeCountDownOperation) Write(p []byte) (n int, err error) {
	f.Calls = append(f.Calls, write)
	return
}
func main() {
	sleeper := ConfigurableSleeper{duration: 1 * time.Second, sleep: time.Sleep}
	CountDown(os.Stdout, &sleeper)
}

func CountDown(writer io.Writer, sleeper Sleeper) {

	for i := range countDownFrom(3) {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprintf(writer, lastMessage)
}

func countDownFrom(i int) func(func(int) bool) {
	return func(yield func(int) bool) {
		for i := i; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}
