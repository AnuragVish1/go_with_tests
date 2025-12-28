package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func SecondHand(tm time.Time) Point {
	return Point{150, 60}
}

func SecondsToRadians(time time.Time) float64 {
	return math.Pi / (30 / float64(time.Second()))
}

func SecondHandPoint(radians float64) Point {
	x := math.Sin(radians)
	y := math.Cos(radians)

	return Point{x, y}
}
