package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hourInHalfClock    = 6
	hourInClock        = 2 * hourInHalfClock
)

func SecondsToRadians(time time.Time) float64 {
	return math.Pi / (secondsInHalfClock / float64(time.Second()))
}

func MinutesToRadians(time time.Time) float64 {
	return (SecondsToRadians(time) / secondsInClock) + math.Pi/(minutesInHalfClock/float64(time.Minute()))
}

func HourToRadians(time time.Time) float64 {
	12
	return (MinutesToRadians(time) / hourInClock) + (math.Pi / (hourInHalfClock / float64(time.Hour()%hourInClock)))
}

func SecondHandPoint(tm time.Time) Point {
	return angleToPoint(SecondsToRadians(tm))
}

func MinuteHandPoint(tm time.Time) Point {
	return angleToPoint(MinutesToRadians(tm))
}

func HourHandPoint(tm time.Time) Point {
	return angleToPoint(HourToRadians(tm))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
