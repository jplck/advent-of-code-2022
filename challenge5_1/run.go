package challenge5_1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jplck/advent-of-code-2022/utils"
)

func Run(inputFile string) {
	reader := utils.GetInputReader(inputFile, utils.Lines)

	stackdef := true
	var rows [][]string
	for reader.Scan() {

		t := reader.Text()

		if t != "" && stackdef {
			t = strings.Replace(t, "    ", "[0]", -1)
			t = strings.Replace(t, " ", "", -1)
			t = strings.Replace(t, "[", "", -1)
			parts := strings.Split(t, "]")

			for i, v := range parts {
				var l []string
				if len(rows) <= i {
					l = make([]string, 0)
					rows = append(rows, l)
				} else {
					l = rows[i]
				}
				if v != "0" {
					l = append(l, v)
					rows[i] = l
				}
			}
			continue
		} else if t == "" {
			stackdef = false
			continue
		}
		Move(rows, t)
	}

	for _, i := range rows {
		fmt.Print(i[0])
	}

}

func AddToFront(item string, slice []string) []string {
	s := append([]string{item}, slice...)
	return s
}

func Move(matrix [][]string, cmd string) {
	cmd = strings.Replace(cmd, "move ", "", -1)
	cmd = strings.Replace(cmd, " from ", ",", -1)
	cmd = strings.Replace(cmd, " to ", ",", -1)

	parts := strings.Split(cmd, ",")

	items, _ := strconv.Atoi(parts[0])
	sourceIdx, _ := strconv.Atoi(parts[1])
	targetIdx, _ := strconv.Atoi(parts[2])

	for range Range(1, items) {
		v := matrix[sourceIdx-1][0]
		l := AddToFront(v, matrix[targetIdx-1])
		matrix[targetIdx-1] = l
		matrix[sourceIdx-1] = matrix[sourceIdx-1][1:]
	}

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
