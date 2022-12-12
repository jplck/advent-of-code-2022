package challenge12

import (
	"fmt"

	"github.com/jplck/advent-of-code-2022/utils"
)

type Position struct {
	X int
	Y int
}

func Run(inputFile string) {
	reader := utils.GetInputReader(inputFile, utils.Lines)
	rows := make(map[int][]int, 0)
	startPos := Position{
		X: 0,
		Y: 0,
	}
	endPos := Position{
		X: 0,
		Y: 0,
	}
	for reader.Scan() {
		readValue := reader.Text()
		row, start, end := ParseRow(readValue)

		if start != -1 {
			startPos = Position{
				X: start,
				Y: len(rows),
			}
		} else if end != -1 {
			endPos = Position{
				X: end,
				Y: len(rows),
			}
		}

		rows[len(rows)] = row
	}

	Move(startPos, endPos, rows)
}

func Move(startPos Position, endPos Position, rows map[int][]int) {
	top := GetOption(rows, startPos.Y, startPos.X)
	fmt.Println(top)
}

func GetOption(rows map[int][]int, posX, posY int) int {
	if value, ok := rows[posY]; ok {
		return value[posX]
	}
	return -1
}

func ParseRow(row string) (rowResult []int, startPos int, endPos int) {
	startPos = -1
	endPos = -1
	rowResult = make([]int, len(row))
	for idx, rowValue := range row {
		charNumber := ConvertToNumber(rowValue)
		if string(rowValue) == "S" {
			startPos = idx
		} else if string(rowValue) == "E" {
			endPos = idx
		}
		rowResult[idx] = charNumber
	}
	return
}

func ConvertToNumber(rowValue rune) int {
	return int(rowValue)
}

func IsSmaller() bool {
	return true
}
