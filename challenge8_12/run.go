package challenge8_12

import (
	"fmt"
	"strconv"

	"github.com/jplck/advent-of-code-2022/utils"
)

type Grid [][]int

func Run(inputFile string) {
	grid := CreateGrid(inputFile)

	visibleTrees := len(grid)*2 + len(grid[0])*2 - 4
	scenicScore := 0

	for posX, valuesInRow := range grid {
		if posX == 0 || posX == len(grid)-1 {
			continue
		}
		for posY := range valuesInRow {
			if posY == 0 || posY == len(valuesInRow)-1 {
				continue
			}

			//Challenge 8.1
			if CheckIfVisible(posX, posY, grid) {
				visibleTrees++
			}

			//Challenge 8.2
			score := CheckScenicScore(posX, posY, grid)
			if score > scenicScore {
				scenicScore = score
			}
		}
	}

	fmt.Printf("VISIBLE TREES: %v\n", visibleTrees)
	fmt.Printf("Highest Scenic Score: %v\n", scenicScore)
}

func CreateGrid(fromFile string) Grid {
	grid := Grid{}
	row := 0
	reader := utils.GetInputReader(fromFile, utils.Lines)
	for reader.Scan() {
		readInput := reader.Text()
		column := convertRow(readInput)
		grid = append(grid, column)
		row++
	}
	return grid
}

func GetGridRays(treeX, treeY int, grid Grid) (left []int, right []int, top []int, bottom []int) {
	left = grid[treeY][:treeX]
	right = grid[treeY][treeX+1:]

	column := GetColumn(treeX, grid)

	top = column[:treeY]
	bottom = column[treeY+1:]
	return
}

func CheckIfVisible(treeX int, treeY int, grid Grid) bool {
	tree := grid[treeY][treeX]
	left, right, top, bottom := GetGridRays(treeX, treeY, grid)
	return utils.Max(left) < tree || utils.Max(right) < tree || utils.Max(top) < tree || utils.Max(bottom) < tree
}

func CheckScenicScore(treeX int, treeY int, grid Grid) int {
	tree := grid[treeY][treeX]
	left, right, top, bottom := GetGridRays(treeX, treeY, grid)
	return GetTreeDistance(tree, left, true) *
		GetTreeDistance(tree, right, false) *
		GetTreeDistance(tree, top, true) *
		GetTreeDistance(tree, bottom, false)
}

func GetTreeDistance(height int, trees []int, shouldReverseInput bool) int {
	if shouldReverseInput {
		trees = utils.Invert(trees)
	}
	for idx, value := range trees {
		if value >= height {
			return idx + 1
		}
	}
	return len(trees)
}

func GetColumn(column int, grid Grid) []int {
	col := make([]int, 0)
	for _, row := range grid {
		col = append(col, row[column])
	}
	return col
}

func convertRow(row string) []int {
	bytes := []byte(row)
	result := make([]int, len(row))
	for idx, value := range bytes {
		conversion, err := strconv.Atoi(string(value))
		utils.Must(err)
		result[idx] = conversion
	}
	return result
}
