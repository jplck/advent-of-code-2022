package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	sum := 0
	highest := make([]int, 223)
	cnt := 0
	for fileScanner.Scan() {
		t := fileScanner.Text()

		if t == "" {
			highest[cnt] = sum
			cnt++

			sum = 0
			continue
		}

		v, err := strconv.Atoi(t)

		if err != nil {
			panic(err)
		}

		sum += v
	}
	highest[len(highest)-1] = sum
	sort.Ints(highest)

	fmt.Println(cnt)
	fmt.Println(Sum(highest[len(highest)-3:]))
}

func Sum(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}
