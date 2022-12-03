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

	for fileScanner.Scan() {
		t := fileScanner.Text()

		len := len(t) / 2

		compartment1 := t[:len]
		compartment2 := t[len:]

		var item string
		for _, c1 := range compartment1 {
			for _, c2 := range compartment2 {
				if c2 == c1 {
					item = string(c1)
					break
				}
			}
		}

		prio := getPriority(string(item))
		sum += prio

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
