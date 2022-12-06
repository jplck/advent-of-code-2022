package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	max := 0
	sum := 0
	for fileScanner.Scan() {
		t := fileScanner.Text()

		if t == "" {
			if sum > max {
				max = sum
			}
			sum = 0
			continue
		}

		v, err := strconv.Atoi(t)

		if err != nil {
			panic(err)
		}

		sum += v
	}

	fmt.Println(max)
}
