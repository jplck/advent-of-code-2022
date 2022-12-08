package challenge5_2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jplck/advent-of-code-2022/utils"
)

type MultiStack interface {
	PushToTop(values []interface{})
	PushToBottom(values []interface{})
	PopTop(items int)
	Read(items int) []string
}

type ElfStack struct {
	Items []string
}

func (s *ElfStack) Push(values []string, top bool) {

	existing := make([]string, len(s.Items))
	copy(existing, s.Items)
	toAdd := make([]string, len(values))
	copy(toAdd, values)

	if top {
		toAdd = append(toAdd, existing...)
	} else {
		toAdd = append(existing, toAdd...)
	}

	s.Items = toAdd
}

func (s *ElfStack) Pop(nrOfItems int) {
	tmp := make([]string, len(s.Items)-nrOfItems)
	copy(tmp, s.Items[nrOfItems:])
	s.Items = tmp
}

func (s *ElfStack) Peak(nrOfItems int) []string {
	return s.Items[:nrOfItems]
}

func Run(inputFile string) {
	var stackColumns []*ElfStack

	reader := utils.GetInputReader(inputFile, utils.Lines)

	lineIsStackDefinition := true
	for reader.Scan() {
		t := reader.Text()

		if strings.TrimSpace(t) == "" {
			lineIsStackDefinition = false
			continue
		}

		if lineIsStackDefinition {
			reg := regexp.MustCompile(`\ {4}|\[\w*]`)
			stackDefInRow := reg.FindAllString(t, -1)

			if len(stackColumns) == 0 {
				stackColumns = make([]*ElfStack, len(stackDefInRow))
			}

			for partIdx, partValue := range stackDefInRow {
				if len(strings.TrimSpace(partValue)) == 0 {
					continue
				}

				l := stackColumns[partIdx]
				if l == nil {
					l = &ElfStack{}
					stackColumns[partIdx] = l
				}

				l.Push([]string{partValue}, false)

			}
			continue
		}

		Move(t, stackColumns)
	}

	for _, stackColumn := range stackColumns {
		fmt.Print(stackColumn.Peak(1)[0])
	}

}

func Move(cmd string, stackColumns []*ElfStack) {
	itemCnt, sourceIdx, targetIdx := ParseCmd(cmd)

	s := stackColumns[sourceIdx-1]
	t := stackColumns[targetIdx-1]

	itemsToCopy := s.Peak(itemCnt)
	t.Push(itemsToCopy, true)
	s.Pop(itemCnt)
}

func ParseCmd(cmd string) (items, sourceIdx, targetIdx int) {
	reg := regexp.MustCompile(`\d+`)

	r := reg.FindAllString(cmd, -1)

	items, err := strconv.Atoi(r[0])
	Must(err)
	sourceIdx, err = strconv.Atoi(r[1])
	Must(err)
	targetIdx, err = strconv.Atoi(r[2])
	Must(err)
	return
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
