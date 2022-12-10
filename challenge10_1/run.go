package challenge10_1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jplck/advent-of-code-2022/utils"
)

type Register []int

func Run(inputFile string) {

	registerX := 1
	reg := Register{}

	reader := utils.GetInputReader(inputFile, utils.Lines)
	signalStrength := 0

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

	for idx, regValue := range reg {
		if idx == 20 || (20+idx)%40 == 0 {
			signalStrength += idx * registerX
		}
		registerX += regValue
	}

	fmt.Printf("Result 10.1: %v\n", signalStrength)
}
