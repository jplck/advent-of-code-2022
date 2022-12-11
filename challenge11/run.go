package challenge11

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
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
}

func (m *Monkey) Inspect() {
	for _, item := range m.Items {
		m.InspectionCounter++
		newWorryValue := ParseOperation(m.Operation, item) / WORRY_DIVISOR

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
			items := FindAllNumbers(readValue)
			currentMonkey.Items = append(currentMonkey.Items, items...)
		} else if strings.HasPrefix(readValue, "  Test: divisible by ") {
			currentMonkey.TestDivisor = FindAllNumbers(readValue)[0]
		} else if strings.HasPrefix(readValue, "    If true: throw to ") {
			currentMonkey.TrueTargetIdx = FindAllNumbers(readValue)[0]
		} else if strings.HasPrefix(readValue, "    If false: throw to ") {
			currentMonkey.FalseTargetIdx = FindAllNumbers(readValue)[0]
		} else if strings.HasPrefix(readValue, "  Operation: new = ") {
			currentMonkey.Operation = strings.Split(readValue, "= ")[1]
		}
	}

	LinkMonkeys(monkeys)

	for round := 0; round < ROUNDS; round++ {
		for _, m := range monkeys {
			m.Inspect()
		}
	}

	inspectionCounts := make([]int, 0)
	for _, m := range monkeys {
		inspectionCounts = append(inspectionCounts, m.InspectionCounter)
	}
	fmt.Println(inspectionCounts)
	arrSlice := inspectionCounts[:] // created a slice of the array
	sort.Sort(sort.Reverse(sort.IntSlice(arrSlice)))

	fmt.Printf("RESULT 11.2: %v\n", arrSlice[0]*arrSlice[1])
}

func LinkMonkeys(monkeys []*Monkey) {
	for _, m := range monkeys {
		m.TrueTarget = monkeys[m.TrueTargetIdx]
		m.FalseTarget = monkeys[m.FalseTargetIdx]
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

func FindAllNumbers(searchInStr string) []int {
	reg := regexp.MustCompile(`[0-9]+`)
	items := reg.FindAllString(searchInStr, -1)
	result := make([]int, 0)
	for _, v := range items {
		num, err := strconv.Atoi(v)
		utils.Must(err)
		result = append(result, num)
	}
	return result
}
