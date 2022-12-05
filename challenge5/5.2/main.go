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

		if t != "" && stackdef {
			t = strings.Replace(t, "    ", "[0]", -1)
			t = strings.Replace(t, " ", "", -1)
			t = strings.Replace(t, "[", "", -1)
			parts := strings.Split(t, "]")

			for i, v := range parts {
				var l []string
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
		} else if t == "" {
			stackdef = false
			continue
		}
		Move(t)
	}
	//fmt.Print(rows)
	/*for _, i := range rows {
		fmt.Print(i[0])
	}*/

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

	v := rows[sourceIdx-1][0:items]
	l := AddToFront(v, rows[targetIdx-1])
	rows[targetIdx-1] = l
	t := rows[sourceIdx-1][len(v):]
	rows[sourceIdx-1] = t
	fmt.Println(rows)
}
