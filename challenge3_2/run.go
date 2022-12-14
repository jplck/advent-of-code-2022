package challenge3_2

import (
	"fmt"

	"github.com/jplck/advent-of-code-2022/utils"
)

func Run(inputFile string) {
	reader := utils.GetInputReader(inputFile, utils.Lines)

	sum := 0
	group := make([]string, 3)
	idx := 0
	for reader.Scan() {

		t := reader.Text()
		group[idx] = t

		idx++

		if idx == 3 {
			idx = 0
			var item string
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
