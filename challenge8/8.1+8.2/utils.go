package main

func Max(numbers []int) int {
	max := numbers[0]
	for _, size := range numbers {
		if size > max {
			max = size
		}
	}
	return max
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func invert(slice []int) []int {
	rev := make([]int, 0)
	for i := len(slice) - 1; i >= 0; i-- {
		rev = append(rev, slice[i])
	}
	return rev
}
