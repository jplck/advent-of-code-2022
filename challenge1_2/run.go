package challenge1_2

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/jplck/advent-of-code-2022/utils"
)

func Run(inputFile string) {
	reader := utils.GetInputReader(inputFile, utils.Lines)

	sum := 0
	highest := make([]int, 0)
	cnt := 0
	for reader.Scan() {
		t := reader.Text()

		if t == "" {
			highest = append(highest, sum)
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

	fmt.Println(Sum(highest[len(highest)-3:]))
}

func Sum(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}
