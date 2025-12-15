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

func sumAll(numbersSlices ...[]int) []int {

	var finalSlice []int

	for _, number := range numbersSlices {
		finalSlice = append(finalSlice, arrayAddition(number))
	}
	return finalSlice

}

func sumAllTails(numberSlices ...[]int) []int {
	var finalSlice []int

	for _, number := range numberSlices {
		if len(number) == 0 {
			finalSlice = append(finalSlice, 0)
			continue
		}
		tail := number[1:]
		finalSlice = append(finalSlice, arrayAddition(tail))
	}
	return finalSlice
}
