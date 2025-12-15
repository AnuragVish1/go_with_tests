package iterations

import "strings"

func main() {

}

func iteration(character string, charRange int) string {
	var repeatString strings.Builder
	for range charRange {
		repeatString.WriteString(character)
	}
	return repeatString.String()
}

func compareStrings(a string, b string) string {
	if strings.Contains(a, b) {
		return "yes"
	}
	return "no"
}
