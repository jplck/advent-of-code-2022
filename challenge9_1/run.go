package challenge9_1

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jplck/advent-of-code-2022/utils"
)

const (
	Left   = "L"
	Right  = "R"
	Top    = "U"
	Bottom = "D"
	Covers = "C"
)

type Position struct {
	X int
	Y int
}

type Grid map[string]string

func Run(inputFile string) {

	reader := utils.GetInputReader(inputFile, utils.Lines)

	grid := make(map[string]string)

	var head = Position{
		X: 0,
		Y: 0,
	}

	var tail = Position{
		X: 0,
		Y: 0,
	}

	for reader.Scan() {

		readValue := reader.Text()
		moveCommandParts := strings.Split(readValue, " ")

		steps, err := strconv.Atoi(moveCommandParts[1])
		utils.Must(err)

		cmd := moveCommandParts[0]

		Move(cmd, steps, grid, &head, &tail)
	}
	fmt.Printf("Result 9.1: %v\n", len(grid))
}

func Move(cmd string, steps int, grid map[string]string, head *Position, tail *Position) {
	r := utils.Range(0, steps-1)
	for range r {

		oldX := head.X
		oldY := head.Y

		switch cmd {
		case Left:
			head.X--
		case Right:
			head.X++
		case Top:
			head.Y++
		case Bottom:
			head.Y--
		default:
			panic("Unable to read move cmd")
		}

		distX := math.Abs(float64(head.X - tail.X))
		distY := math.Abs(float64(head.Y - tail.Y))

		if distX == 2 || distY == 2 {
			tail.X = oldX
			tail.Y = oldY

			grid[fmt.Sprintf("%v/%v", tail.X, tail.Y)] = "#"

		}

	}
}
