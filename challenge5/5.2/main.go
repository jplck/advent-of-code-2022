package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var rows []*ElfStack

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

func (s *ElfStack) PopTop(items int) {
	tmp := make([]string, len(s.Items)-items)
	copy(tmp, s.Items[items:])
	s.Items = tmp
}

func (s *ElfStack) Read(items int) []string {
	return s.Items[:items]
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	stackdef := true
	for fileScanner.Scan() {
		t := fileScanner.Text()

		if t == "" || t == " " {
			stackdef = false
			fmt.Println(rows)
			continue
		} else if t != "" && stackdef {
			reg := regexp.MustCompile(`\ {4}|\[\w*]`)
			parts := reg.FindAllString(t, -1)

			if len(rows) == 0 {
				rows = make([]*ElfStack, len(parts))
			}

			for i, v := range parts {
				if len(strings.TrimSpace(v)) == 0 {
					continue
				}

				l := rows[i]
				if l == nil {
					l = &ElfStack{}
					rows[i] = l
				}

				l.Push([]string{v}, false)

			}
			continue
		}
		Move(t)
	}

	for _, i := range rows {
		fmt.Print(i.Read(1)[0])
	}

}

func Move(cmd string) {
	itemCnt, sourceIdx, targetIdx := ParseCmd(cmd)

	s := rows[sourceIdx-1]
	t := rows[targetIdx-1]

	itemsToCopy := s.Read(itemCnt)
	t.Push(itemsToCopy, true)
	s.PopTop(itemCnt)
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
