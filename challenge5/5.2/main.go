package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rows [][]string

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
			t = strings.Replace(t, "    ", "[0]", -1)
			t = strings.Replace(t, " ", "", -1)
			t = strings.Replace(t, "[", "", -1)
			parts := strings.Split(t, "]")

			var l []string
			for i, v := range parts[:len(parts)-1] {
				v := v
				if len(rows) <= i {
					l = make([]string, 0)
					rows = append(rows, l)
				} else {
					l = rows[i]
				}
				if v != "0" {
					l = append(l, v)
					rows[i] = l
				}
			}
			continue
		}
		Move(t)
	}

	for _, i := range rows {
		fmt.Print(i[0])
	}

}

func AddToFront(items []string, slice []string) []string {
	s := append(items, slice...)
	return s
}

func Move(cmd string) {
	cmd = strings.Replace(cmd, "move ", "", -1)
	cmd = strings.Replace(cmd, " from ", ",", -1)
	cmd = strings.Replace(cmd, " to ", ",", -1)

	parts := strings.Split(cmd, ",")

	items, _ := strconv.Atoi(parts[0])
	sourceIdx, _ := strconv.Atoi(parts[1])
	targetIdx, _ := strconv.Atoi(parts[2])

	s := rows[sourceIdx-1]
	t := rows[targetIdx-1]

	cS := make([]string, len(s))
	copy(cS, s)

	cT := make([]string, len(t))
	copy(cT, t)

	cp := s[:items]

	nS := s[items:]
	newSource := make([]string, len(nS))
	copy(newSource, nS)

	cp = append(cp, cT...)

	rows[sourceIdx-1] = newSource
	rows[targetIdx-1] = cp

	fmt.Printf("%v -> %v -> %v -> %v -> %v -> %v\n", s, t, cS, cT, cp, rows[targetIdx-1])
}
