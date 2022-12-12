package challenge12

import (
	"fmt"

	"github.com/jplck/advent-of-code-2022/utils"
)

type Position struct {
	X int
	Y int
}

type TreeNode struct {
	Left   *TreeNode
	Right  *TreeNode
	Uppper *TreeNode
	Bottom *TreeNode
}

func Run(inputFile string) {
	reader := utils.GetInputReader(inputFile, utils.Lines)
	rows := make(map[int]map[int]int, 0)
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
		}
		if end != -1 {
			endPos = Position{
				X: end,
				Y: len(rows),
			}
		}

		rows[len(rows)] = row

	}
	s := StoreVisit(nil, startPos)
	Move(startPos, endPos, rows, s, 0)
}

func Move(startPos Position, endPos Position, rows map[int]map[int]int, route map[string]int, steps int) {
	if startPos.X == endPos.X && startPos.Y == endPos.Y {
		fmt.Println(route)
		fmt.Println(steps)
	}

	currentElevation := rows[startPos.Y][startPos.X]

	if isOption, _, newPosition := GetOption(rows, startPos.X+1, startPos.Y, route, currentElevation); isOption {
		r := StoreVisit(route, newPosition)
		Move(newPosition, endPos, rows, r, steps+1)
	}
	if isOption, _, newPosition := GetOption(rows, startPos.X-1, startPos.Y, route, currentElevation); isOption {
		r := StoreVisit(route, newPosition)
		Move(newPosition, endPos, rows, r, steps+1)
	}
	if isOption, _, newPosition := GetOption(rows, startPos.X, startPos.Y-1, route, currentElevation); isOption {
		r := StoreVisit(route, newPosition)
		Move(newPosition, endPos, rows, r, steps+1)
	}
	if isOption, _, newPosition := GetOption(rows, startPos.X, startPos.Y+1, route, currentElevation); isOption {
		r := StoreVisit(route, newPosition)
		Move(newPosition, endPos, rows, r, steps+1)
	}
}

func AlreadyVisited(position Position, route map[string]int) bool {
	if _, ok := route[fmt.Sprintf("%v/%v", position.Y, position.X)]; ok {
		return true
	}
	return false
}

func StoreVisit(route map[string]int, position Position) map[string]int {
	route = CopyMap(route)
	idx := fmt.Sprintf("%v/%v", position.Y, position.X)
	//fmt.Println(idx)
	route[idx] = len(route)
	return route
}

func CopyMap(m map[string]int) map[string]int {
	n := make(map[string]int)
	for idx, v := range m {
		n[idx] = v
	}
	return n
}

func GetOption(rows map[int]map[int]int, posX, posY int, route map[string]int, currentElevation int) (isOption bool, result int, newPosition Position) {
	if row, ok := rows[posY]; ok {
		if colValue, ok := row[posX]; ok && !AlreadyVisited(Position{X: posX, Y: posY}, route) {
			hasTheRightElevation := colValue-currentElevation <= 1
			return hasTheRightElevation, colValue, Position{X: posX, Y: posY}
		}
	}
	return false, -1, Position{}
}

func ParseRow(row string) (rowResult map[int]int, startPos int, endPos int) {
	startPos = -1
	endPos = -1
	rowResult = make(map[int]int, len(row))
	for idx, rowValue := range row {
		if string(rowValue) == "S" {
			startPos = idx
			rowValue = 'a'
		} else if string(rowValue) == "E" {
			endPos = idx
			rowValue = 'z'
		}
		charNumber := ConvertToNumber(rowValue)
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
