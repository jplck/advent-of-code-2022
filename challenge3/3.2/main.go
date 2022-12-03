package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	sum := 0
	group := make([]string, 3)
	idx := 0
	for fileScanner.Scan() {

		t := fileScanner.Text()
		group[idx] = t

		idx++

		if idx == 3 {
			idx = 0
			var item string
			fmt.Println(group[2])
			for _, c1 := range group[0] {
				for _, c2 := range group[1] {
					if c2 == c1 {
						for _, c3 := range group[2] {
							if c2 == c3 {
								item = string(c3)
							}
						}
					}
				}
			}
			fmt.Println(item)
			prio := getPriority(string(item))
			sum += prio
		}
	}
	fmt.Println(sum)
}

func getPriority(value string) int {
	var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	for i, v := range alphabet {
		if string(v) == value {
			return i + 1
		}
	}
	return 0
}
