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
				if !kindaEqualFloat64(got, test.angle) {
					t.Errorf("Got %v want %v", got, test.angle)
				}
			})
		}
	})
}

func TestMinutesToRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
	}

	for _, test := range cases {
		t.Run("Checking radians value for minutes", func(t *testing.T) {
			got := MinutesToRadians(test.time)
			if !kindaEqualFloat64(got, test.angle) {
				t.Errorf("Got %v want %v", got, test.angle)
			}
		})
	}
}

func TestHourToRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(21, 0, 0), math.Pi * 1.5},
		{simpleTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
	}

	for _, test := range cases {
		t.Run("Checking if the hour to radians is correct", func(t *testing.T) {
			got := HourToRadians(test.time)
			if !kindaEqualFloat64(got, test.angle) {
				t.Errorf("Got %v want %v", got, test.angle)
			}
		})
	}
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
			got := SecondHandPoint(test.time)

			if !kindaEqualPoints(got, test.point) {
				t.Errorf("Got %v want %v", got, test.point)
			}
		})
	}
}

func TestMinuteHandPoints(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 30, 0), Point{0, -1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
	}

	for _, test := range cases {

		t.Run("Testing the Minute hand points", func(t *testing.T) {
			got := MinuteHandPoint(test.time)

			if !kindaEqualPoints(got, test.point) {
				t.Errorf("Got %v want %v", got, test.point)
			}
		})
	}
}

func TestHourHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(6, 0, 0), Point{0, -1}},
		{simpleTime(21, 0, 0), Point{-1, 0}},
	}

	for _, test := range cases {
		t.Run("Checking for the point for the hour hand", func(t *testing.T) {
			got := HourHandPoint(test.time)
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
