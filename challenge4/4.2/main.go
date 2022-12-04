package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	sum := 0
	for fileScanner.Scan() {

		t := fileScanner.Text()
		parts := strings.Split(t, ",")

		e1l, e1r := ParseRange(parts[0])
		e2l, e2r := ParseRange(parts[1])

		r1 := Range(e1l, e1r)
		r2 := Range(e2l, e2r)

	out:
		for _, v1 := range r1 {
			for _, v2 := range r2 {
				if v1 == v2 {
					sum++
					break out
				}
			}
		}

	}
	fmt.Println(sum)
}

func ParseRange(rangeStr string) (left, right int) {
	e := strings.Split(rangeStr, "-")
	left, err := strconv.Atoi(e[0])
	must(err)
	right, err = strconv.Atoi(e[1])
	must(err)
	return
}

func Range(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}