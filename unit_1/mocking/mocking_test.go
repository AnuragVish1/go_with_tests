package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountDown(t *testing.T) {

	t.Run("CountDown in descending order", func(t *testing.T) {
		buffer := bytes.Buffer{}
		spyCaller := SpyCaller{}
		CountDown(&buffer, &spyCaller)
		got := buffer.String()
		want := "3\n2\n1\nGo"

		if got != want {
			t.Errorf("Got %s want %s", got, want)
		}

		if spyCaller.Calls != 3 {
			t.Errorf("Sleep only ran for %v times but should have ran for 3 times", spyCaller.Calls)
		}
	})

	t.Run("If printing in correct order", func(t *testing.T) {
		fakeCountDownOp := &FakeCountDownOperation{}

		CountDown(fakeCountDownOp, fakeCountDownOp)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(fakeCountDownOp.Calls, want) {
			t.Errorf("Got %v want %v", fakeCountDownOp.Calls, want)
		}
	})
}

func TestConfigurableSleep(t *testing.T) {
	t.Run("Testing if the sleep is configurable", func(t *testing.T) {
		sleep := 5 * time.Second
		fakeTime := &FakeTime{}
		sleeper := ConfigurableSleeper{sleep, fakeTime.SetDurationSlept}
		sleeper.Sleep()

		if fakeTime.durationSlept != sleep {
			t.Errorf("Got %q want %q", fakeTime.durationSlept, sleep)
		}
	})
}
