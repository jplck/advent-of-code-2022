package challenge11

import (
	"fmt"
	"strings"

	"github.com/Knetic/govaluate"
	"github.com/jplck/advent-of-code-2022/utils"
)

const WORRY_DIVISOR = 1

type Monkey struct {
	Items             []int
	TestDivisor       int
	TrueTargetIdx     int
	FalseTargetIdx    int
	Operation         string
	TrueTarget        *Monkey
	FalseTarget       *Monkey
	InspectionCounter int
}

func GetCommonFactor(monkeys []*Monkey) int {
	common := 1
	for _, monkey := range monkeys {
		common *= monkey.TestDivisor
	}
	return common
}

func (m *Monkey) Inspect(superMod int) {
	for _, item := range m.Items {
		m.InspectionCounter++

		newWorryValue := ParseOperation(m.Operation, item) / WORRY_DIVISOR

		if WORRY_DIVISOR != 3 {
			/*
				Remark for me: Challenge was to find a way to reduce the sice of item,
				without changing the outcome of the following logic. As I understood now, with the help
				of a reddit thread, you can use module arithmetics to find a common multiplier/divisor.
				was a bit above my head :)
			*/
			newWorryValue = newWorryValue % superMod
		}

		if newWorryValue%m.TestDivisor == 0 {
			m.TrueTarget.Throw(newWorryValue)
			continue
		}
		m.FalseTarget.Throw(newWorryValue)
	}
	m.Items = make([]int, 0)
}

func (m *Monkey) Throw(item int) {
	m.Items = append(m.Items, item)
}

func Run(inputFile string, rounds int) {
	monkeys := ParseMonkeys(inputFile)
	RunInspection(monkeys, rounds)

	inspectionCounts := make([]int, 0)
	for _, m := range monkeys {
		inspectionCounts = append(inspectionCounts, m.InspectionCounter)
	}
	fmt.Println(inspectionCounts)

	sortedResult := utils.SortArrayOfInts(inspectionCounts)

	fmt.Printf("RESULT 11.2: %v\n", sortedResult[0]*sortedResult[1])
}

func ParseMonkeys(inputFile string) []*Monkey {

	reader := utils.GetInputReader(inputFile, utils.Lines)

	monkeys := make([]*Monkey, 0)
	var currentMonkey *Monkey

	for reader.Scan() {
		readValue := reader.Text()
		if strings.HasPrefix(readValue, "Monkey") {
			currentMonkey = &Monkey{
				InspectionCounter: 0,
			}
			monkeys = append(monkeys, currentMonkey)
		} else if strings.HasPrefix(readValue, "  Starting items:") {
			items := utils.FindAllNumbers(readValue)
			currentMonkey.Items = append(currentMonkey.Items, items...)
		} else if strings.HasPrefix(readValue, "  Test: divisible by ") {
			currentMonkey.TestDivisor = utils.FindAllNumbers(readValue)[0]
		} else if strings.HasPrefix(readValue, "    If true: throw to ") {
			currentMonkey.TrueTargetIdx = utils.FindAllNumbers(readValue)[0]
		} else if strings.HasPrefix(readValue, "    If false: throw to ") {
			currentMonkey.FalseTargetIdx = utils.FindAllNumbers(readValue)[0]
		} else if strings.HasPrefix(readValue, "  Operation: new = ") {
			currentMonkey.Operation = strings.Split(readValue, "= ")[1]
		}
	}
	return monkeys
}

func RunInspection(monkeys []*Monkey, rounds int) {
	superMod := GetCommonFactor(monkeys)
	for round := 0; round < rounds; round++ {
		for _, m := range monkeys {
			m.TrueTarget = monkeys[m.TrueTargetIdx]
			m.FalseTarget = monkeys[m.FalseTargetIdx]
			m.Inspect(superMod)
		}
	}
}

func ParseOperation(operation string, old int) int {
	expression, err := govaluate.NewEvaluableExpression(operation)
	utils.Must(err)
	parameters := make(map[string]interface{}, 8)
	parameters["old"] = old
	result, err := expression.Evaluate(parameters)
	utils.Must(err)
	return int(result.(float64))
}
