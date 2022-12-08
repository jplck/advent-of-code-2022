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
	visibleTrees := len(grid)*2 + len(grid[0])*2 - 4

	for row, columnArray := range grid {
		if row == 0 || row == len(grid)-1 {
			continue
		}
		for colPosition := range columnArray {
			if colPosition == 0 || colPosition == len(columnArray)-1 {
				continue
			}
			if checkIfVisible(row, colPosition, grid) {
				visibleTrees++
			}
		}
	}

	fmt.Printf("VISIBLE TREES: %v\n", visibleTrees)
}

func checkIfVisible(treePosRow int, treePosCol int, grid Grid) bool {
	tree := grid[treePosRow][treePosCol]

	treeLeft := Max(grid[treePosRow][:treePosCol])
	treeRight := Max(grid[treePosRow][treePosCol+1:])

	column := getColumn(treePosCol, grid)

	treeTop := Max(column[:treePosRow])
	treeBottom := Max(column[treePosRow+1:])

	return treeLeft < tree || treeRight < tree || treeTop < tree || treeBottom < tree
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
