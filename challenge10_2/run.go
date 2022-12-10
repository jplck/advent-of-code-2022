package challenge10_2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jplck/advent-of-code-2022/utils"
)

type Register []int

func Sort(values map[int]string) []string {
	result := make([]string, len(values))
	for k, v := range values {
		result[k] = v
	}
	return result
}

func Run(inputFile string) {

	registerX := 1
	reg := Register{}

	reader := utils.GetInputReader(inputFile, utils.Lines)

	for reader.Scan() {
		readValue := reader.Text()
		commandParts := strings.Split(readValue, " ")

		switch commandParts[0] {
		case "noop":
			reg = append(reg, 0)
		case "addx":
			value, err := strconv.Atoi(commandParts[1])
			utils.Must(err)
			e := []int{
				0,
				value,
			}
			reg = append(reg, e...)

		default:
			panic("command not valid")
		}
	}
	crt := CRT{}
	crt.Init()

	rowIdx := 0
	colIdx := 0
	for idx, regValue := range reg {
		if (idx)%40 == 0 {
			colIdx = 0
			rowIdx++
		}
		crt.Add(rowIdx, colIdx, registerX)
		registerX += regValue
		colIdx++
	}

	fmt.Println("Result 10.2")
	fmt.Println("--------------------------")
	for _, v := range crt.Rows {
		res := Sort(v)
		for _, r := range res {
			fmt.Print(r)
		}
		fmt.Print("\n")
	}
}

type CRT struct {
	Rows map[int]map[int]string
}

func (c *CRT) Init() {
	c.Rows = make(map[int]map[int]string)
}

func (c *CRT) Add(rowIdx int, posX int, spritePos int) {
	if _, ok := c.Rows[rowIdx]; !ok {
		c.Rows[rowIdx] = make(map[int]string)
	}

	val := " "

	if spritePos != 0 {
		for _, spriteIdx := range utils.Range(spritePos-1, spritePos+1) {
			if spriteIdx == posX {
				val = "#"
				break
			}
		}
	}

	c.Rows[rowIdx][posX] = val
}
