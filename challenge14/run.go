package challenge14

import (
	"fmt"
	"strings"

	"github.com/jplck/advent-of-code-2022/utils"
)

type CaveMap map[string]string

func Run(inputFile string) {
	reader := utils.GetInputReader(inputFile, utils.Lines)

	caveMap := make(CaveMap)

	for reader.Scan() {
		readValue := reader.Text()
		commands := strings.Split(readValue, " -> ")
		CreateStones(commands, caveMap)
	}
	fmt.Println(caveMap)
}

func CreateStones(commands []string, caveMap CaveMap) {
	var prevCmdIndices []int
	for _, placementCmd := range commands {
		indices := utils.FindAllNumbers(placementCmd)

		if prevCmdIndices != nil {
			xPrev := prevCmdIndices[0]
			yPrev := prevCmdIndices[1]
			x := indices[0]
			y := indices[1]
			xRange := utils.Range(xPrev, x)
			yRange := utils.Range(yPrev, y)

			if len(xRange) == 1 {
				for _, yPos := range yRange {
					caveMap[fmt.Sprintf("%v/%v", xRange[0], yPos)] = "#"
				}
			} else if len(yRange) == 1 {
				for _, xPos := range xRange {
					caveMap[fmt.Sprintf("%v/%v", xPos, yRange[0])] = "#"
				}
			}
		}

		prevCmdIndices = indices
	}
}
