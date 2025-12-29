package main

import (
	"os"
	"time"
	"unit1/maths/clockface"
)

func main() {
	time := time.Now()
	clockface.SvgWriter(os.Stdout, time)
}
