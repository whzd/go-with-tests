package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	var sums []int
	for _, num := range numbers {
		sums = append(sums, Sum(num))
	}
	return sums
}
