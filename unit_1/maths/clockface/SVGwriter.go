package clockface

import (
	"fmt"
	"io"
	"time"
)

const (
	secondHandLength = 90
	minuteHandLength = 70
	hourHandLength   = 40
	clockCenterX     = 150
	clockCenterY     = 150
)

const (
	svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

	bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

	svgEnd = `</svg>`
)

func SvgWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	SecondHand(w, t)
	MinuteHand(w, t)
	HourHand(w, t)
	io.WriteString(w, svgEnd)
}

func HourHand(w io.Writer, t time.Time) {
	p := makePoint(HourHandPoint(t), hourHandLength)

	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%0.3f" y2="%0.3f"
      style="fill:none;stroke:#171717;stroke-width:2px;stroke-linecap:round;"/>`, p.X, p.Y)
}

func MinuteHand(w io.Writer, t time.Time) {
	p := makePoint(MinuteHandPoint(t), minuteHandLength)

	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%0.3f" y2="%0.3f"
      style="fill:none;stroke:#171717;stroke-width:2px;stroke-linecap:round;"/>`, p.X, p.Y)
}

func SecondHand(w io.Writer, tm time.Time) {
	p := makePoint(SecondHandPoint(tm), secondHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%0.3f" y2="%0.3f"
      style="fill:none;stroke:#171717;stroke-width:0.8px;stroke-linecap:round;"/>`, p.X, p.Y)
}

func makePoint(p Point, length float64) Point {
	p = Point{X: p.X * length, Y: p.Y * length}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}
	return p
}
