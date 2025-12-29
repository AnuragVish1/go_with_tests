package clockface_test

import (
	"bytes"
	"encoding/xml"
	"slices"
	"testing"
	"time"
	"unit1/maths/clockface"
)

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}
type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`

	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}
type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

// func TestSecondHandPoints(t *testing.T) {
// 	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)
// 	want := clockface.Point{X: 150, Y: 150 - 90}
// 	got := clockface.SecondHand(tm)

// 	if got != want {
// 		t.Errorf("Got %v want %v", got, want)
// 	}
// }

// func TestSecondHandAtHalfMin(t *testing.T) {
// 	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)
// 	want := clockface.Point{X: 150, Y: 150 + 90}

// 	got := clockface.SecondHand(tm)
// 	if got != want {
// 		t.Errorf("Got %v want %v", got, want)
// 	}
// }

func TestSvgWriterMidNight(t *testing.T) {

	t.Run("Check if the svg is correct", func(t *testing.T) {

		tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

		buff := bytes.Buffer{}

		clockface.SvgWriter(&buff, tm)

		svg := SVG{}

		xml.Unmarshal(buff.Bytes(), &svg)

		want := Line{150, 150, 150, 60}

		if slices.Contains(svg.Line, want) {
			return
		}

		t.Errorf("Wanted %v but got %v", want, buff.String())

	})
}

func TestSvgWriterSecondHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{simpleTime(0, 0, 0), Line{150, 150, 150, 60}},
		{simpleTime(0, 0, 30), Line{150, 150, 150, 240}},
	}

	for _, test := range cases {
		t.Run("Checking if value is correct", func(t *testing.T) {
			buff := bytes.Buffer{}
			clockface.SvgWriter(&buff, test.time)
			svg := SVG{}
			xml.Unmarshal(buff.Bytes(), &svg)

			if !slices.Contains(svg.Line, test.line) {
				t.Errorf("Got %v want %v", svg.Line, test.line)
			}

		})
	}
}

func TestSvgWriterMinuteHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{simpleTime(0, 0, 0), Line{150, 150, 150, 80}},
	}

	for _, test := range cases {
		t.Run("Checking if the value for minute is correct", func(t *testing.T) {
			buff := bytes.Buffer{}
			clockface.SvgWriter(&buff, test.time)
			svg := SVG{}
			xml.Unmarshal(buff.Bytes(), &svg)
			if !slices.Contains(svg.Line, test.line) {
				t.Errorf("Got %v want %v", svg.Line, test.line)
			}
		})
	}

}

func TestSvgWriterHourHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{simpleTime(6, 0, 0), Line{150, 150, 150, 190}},
	}

	for _, test := range cases {
		t.Run("Checking for hour hand", func(t *testing.T) {
			buff := bytes.Buffer{}
			clockface.SvgWriter(&buff, test.time)
			svg := SVG{}
			xml.Unmarshal(buff.Bytes(), &svg)
			if !slices.Contains(svg.Line, test.line) {
				t.Errorf("Got %v want %v", svg.Line, test.line)
			}
		})
	}
}

func simpleTime(i1, i2, i3 int) time.Time {
	return time.Date(312, time.December, 28, i1, i2, i3, 0, time.UTC)
}
