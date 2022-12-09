package utils

import (
	"bufio"
	"os"
)

func Min(numbers []int) int {
	min := numbers[0]
	for _, size := range numbers {
		if size < min {
			min = size
		}
	}
	return min
}

func Max(numbers []int) int {
	max := numbers[0]
	for _, size := range numbers {
		if size > max {
			max = size
		}
	}
	return max
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func Invert(slice []int) []int {
	rev := make([]int, 0)
	for i := len(slice) - 1; i >= 0; i-- {
		rev = append(rev, slice[i])
	}
	return rev
}

const (
	Lines = 0
	Chars = 1
	Words = 2
)

func GetInputReader(dir string, readType int) *bufio.Scanner {
	fileHandle, err := os.Open(dir)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(fileHandle)

	switch readType {
	case Lines:
		fileScanner.Split(bufio.ScanLines)
	case Chars:
		fileScanner.Split(bufio.ScanRunes)
	case Words:
		fileScanner.Split(bufio.ScanWords)
	default:
		fileScanner.Split(bufio.ScanLines)
	}
	return fileScanner
}

func Range(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
