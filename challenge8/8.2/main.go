package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Grid [][]int

func main() {
	grid := createGrid()
	scenicScore := 0

	for row, columnArray := range grid {
		if row == 0 || row == len(grid)-1 {
			continue
		}
		for colPosition := range columnArray {
			if colPosition == 0 || colPosition == len(columnArray)-1 {
				continue
			}

			score := checkIfVisible(row, colPosition, grid)
			if score > scenicScore {
				scenicScore = score
			}

		}
	}

	fmt.Printf("Highest Scenic Score: %v\n", scenicScore)
}

func checkIfVisible(treePosRow int, treePosCol int, grid Grid) int {
	tree := grid[treePosRow][treePosCol]

	treeLeft := GetTreeDistance(tree, grid[treePosRow][:treePosCol], true)
	treeRight := GetTreeDistance(tree, grid[treePosRow][treePosCol+1:], false)

	column := getColumn(treePosCol, grid)

	treeTop := GetTreeDistance(tree, column[:treePosRow], true)
	treeBottom := GetTreeDistance(tree, column[treePosRow+1:], false)

	return treeLeft * treeRight * treeTop * treeBottom
}

func reverse(trees []int) []int {
	rev := make([]int, 0)
	for i := len(trees) - 1; i >= 0; i-- {
		rev = append(rev, trees[i])
	}
	return rev
}

func GetTreeDistance(height int, trees []int, shouldReverseInput bool) int {
	if shouldReverseInput {
		trees = reverse(trees)
	}
	for idx, value := range trees {
		if value >= height {
			return idx + 1
		}
	}
	return len(trees)
}

func getColumn(column int, grid Grid) []int {
	col := make([]int, 0)
	for _, row := range grid {
		col = append(col, row[column])
	}
	return col
}

func createGrid() Grid {
	fileHandle, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(fileHandle)
	fileScanner.Split(bufio.ScanLines)

	grid := Grid{}
	row := 0

	for fileScanner.Scan() {
		readInput := fileScanner.Text()
		column := convertRow(readInput)
		grid = append(grid, column)
		row++
	}
	return grid
}

func convertRow(row string) []int {
	bytes := []byte(row)
	result := make([]int, len(row))
	for idx, value := range bytes {
		conversion, err := strconv.Atoi(string(value))
		must(err)
		result[idx] = conversion
	}
	return result
}
