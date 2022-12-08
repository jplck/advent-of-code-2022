package challenge4_1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jplck/advent-of-code-2022/utils"
)

func Run(inputFile string) {
	reader := utils.GetInputReader(inputFile, utils.Lines)

	sum := 0
	for reader.Scan() {

		t := reader.Text()
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
			sum++
		}
	}
	fmt.Println(sum)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
