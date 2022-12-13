package challenge13

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/jplck/advent-of-code-2022/utils"
)

type ParsedResult []interface{}

func Run(inputFile string) {
	reader := utils.GetInputReader(inputFile, utils.Lines)
	pack := make([]ParsedResult, 2)
	cnt := 0
	sum := 0
	packIdx := 1
	for reader.Scan() {
		readValue := reader.Text()

		if readValue == "" {
			cnt = 0
			if Calculate(pack[0], pack[1]) {
				sum += packIdx
				fmt.Println(packIdx)
			}
			packIdx++
			continue
		}

		var parsedResult ParsedResult
		err := json.Unmarshal([]byte(readValue), &parsedResult)
		utils.Must(err)

		pack[cnt] = parsedResult
		cnt++
	}
	fmt.Println(sum)
}

func Calculate(left ParsedResult, right ParsedResult) bool {
	for leftIdx, leftValue := range left {

		if len(right)-1 >= leftIdx {

			rightValue := right[leftIdx]

			leftType := reflect.TypeOf(leftValue).String()
			rightType := reflect.TypeOf(rightValue).String()

			if leftType != rightType {
				if leftType == "float64" {
					leftValue = []interface{}{
						leftValue,
					}
				} else if rightType == "float64" {
					rightValue = []interface{}{
						rightValue,
					}
				}
				return Calculate(leftValue.([]interface{}), rightValue.([]interface{}))
			}

			var trueLeftValue float64
			var trueRightValue float64

			if leftType == "float64" && rightType == "float64" {
				trueLeftValue = leftValue.(float64)
				trueRightValue = rightValue.(float64)
			}
			if leftType == "[]interface {}" && rightType == "[]interface {}" {
				return Calculate(leftValue.([]interface{}), rightValue.([]interface{}))
			}

			if trueLeftValue < trueRightValue {
				return true
			} else if trueLeftValue > trueRightValue {
				return false
			}

		} else {
			return false
		}

	}
	return true
}
