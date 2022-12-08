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
