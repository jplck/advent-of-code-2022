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
		e1 := strings.Split(parts[0], "-")
		e2 := strings.Split(parts[1], "-")

		e1l, err := strconv.Atoi(e1[0])
		must(err)
		e1r, err := strconv.Atoi(e1[1])
		must(err)

		e2l, err := strconv.Atoi(e2[0])
		must(err)
		e2r, err := strconv.Atoi(e2[1])
		must(err)

		if e1l <= e2l && e1r >= e2r || e1l >= e2l && e1r <= e2r {
			fmt.Printf("%d-%d,%d-%d\n", e1l, e1r, e2l, e2r)
			sum++
		}

		//fmt.Printf("%d-%d,%d-%d\n", e1l, e1r, e2l, e2r)

	}
	fmt.Println(sum)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
