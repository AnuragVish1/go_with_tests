package arraysslices

func main() {

}

func arrayAddition(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
