package clockface

import (
	"math"
	"testing"
	"time"
)

func TestSecondsToRadians(t *testing.T) {
	t.Run("Converting seconds to radians", func(t *testing.T) {

		cases := []struct {
			time  time.Time
			angle float64
		}{
			{simpleTime(0, 0, 30), math.Pi},
			{simpleTime(0, 0, 0), 0},
			{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
			{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
		}

		for _, test := range cases {
			t.Run(testName(test.time), func(t *testing.T) {
				got := SecondsToRadians(test.time)
				want := test.angle

				if got != want {
					t.Errorf("Got %v want %v", got, want)
				}
			})
		}
	})
}

func simpleTime(i1, i2, i3 int) time.Time {
	return time.Date(312, time.December, 28, i1, i2, i3, 0, time.UTC)
}

func testName(time time.Time) string {
	return time.Format("20:15:13")
}

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, test := range cases {
		t.Run("Please run", func(t *testing.T) {
			radians := SecondsToRadians(test.time)
			got := SecondHandPoint(radians)

			if !kindaEqualPoints(got, test.point) {
				t.Errorf("Got %v want %v", got, test.point)
			}
		})
	}
}

func kindaEqualFloat64(a, b float64) bool {
	const threshold = 1e-7
	return math.Abs(a-b) < threshold
}

func kindaEqualPoints(a, b Point) bool {
	return kindaEqualFloat64(a.X, b.X) && kindaEqualFloat64(a.Y, b.Y)
}
