package challenge1_1

import (
	"fmt"
	"strconv"

	"github.com/jplck/advent-of-code-2022/utils"
)

func Run(inputFile string) {
	reader := utils.GetInputReader(inputFile, utils.Lines)

	max := 0
	sum := 0
	for reader.Scan() {
		t := reader.Text()

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
