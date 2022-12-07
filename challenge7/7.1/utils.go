package main

func Min(numbers []int) int {
	min := numbers[0]
	for _, size := range numbers {
		if size < min {
			min = size
		}
	}
	return min
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
