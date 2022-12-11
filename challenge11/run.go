package challenge11

import (
	"fmt"
	"strings"

	"github.com/Knetic/govaluate"
	"github.com/jplck/advent-of-code-2022/utils"
)

const ROUNDS = 20
const WORRY_DIVISOR = 3

type Monkey struct {
	Items             []int
	TestDivisor       int
	TrueTargetIdx     int
	FalseTargetIdx    int
	Operation         string
	TrueTarget        *Monkey
	FalseTarget       *Monkey
	InspectionCounter int
	Idx               int
}

func (m *Monkey) Inspect() {
	for _, item := range m.Items {
		m.InspectionCounter++

		newWorryValue := int(ParseOperation(m.Operation, float64(item)) / WORRY_DIVISOR)

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

func Run(inputFile string) {
	monkeys := ParseMonkeys(inputFile)
	RunInspection(monkeys)

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
				Idx:               len(monkeys),
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

func RunInspection(monkeys []*Monkey) {
	for round := 0; round < ROUNDS; round++ {
		for _, m := range monkeys {
			m.TrueTarget = monkeys[m.TrueTargetIdx]
			m.FalseTarget = monkeys[m.FalseTargetIdx]
			m.Inspect()
		}
	}
}

func ParseOperation(operation string, old float64) float64 {
	expression, err := govaluate.NewEvaluableExpression(operation)
	utils.Must(err)
	parameters := make(map[string]interface{}, 8)
	parameters["old"] = old
	result, err := expression.Evaluate(parameters)
	utils.Must(err)
	return result.(float64)
}
