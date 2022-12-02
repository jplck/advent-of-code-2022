package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const ROCK = "A"
const PAPER = "B"
const SCISSOR = "C"
const RESP_ROCK = "X"
const RESP_PAPER = "Y"
const RESP_SCISSOR = "Z"

const LOOSE = RESP_ROCK
const DRAW = RESP_PAPER
const WIN = RESP_SCISSOR

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	sum := 0

	for fileScanner.Scan() {
		t := fileScanner.Text()
		parts := strings.Split(t, " ")
		a := parts[0]
		b := parts[1]

		if b == LOOSE { //loose
			switch a {
			case ROCK:
				b = RESP_SCISSOR
			case PAPER:
				b = RESP_ROCK
			case SCISSOR:
				b = RESP_PAPER
			}
		}

		if b == DRAW { //draw
			switch a {
			case ROCK:
				b = RESP_ROCK
			case PAPER:
				b = RESP_PAPER
			case SCISSOR:
				b = RESP_SCISSOR
			}
		}

		if b == WIN { //win
			switch a {
			case ROCK:
				b = RESP_PAPER
			case PAPER:
				b = RESP_SCISSOR
			case SCISSOR:
				b = RESP_ROCK
			}
		}

		sum += calcResult(a, b)
	}

	fmt.Println(sum)
}

func calcResult(a, b string) int {
	result := 0

	if a == ROCK {
		switch b {
		case RESP_ROCK:
			result = 3 + 1
		case RESP_PAPER:
			result = 6 + 2
		case RESP_SCISSOR:
			result = 0 + 3
		}
	}

	if a == PAPER {
		switch b {
		case RESP_ROCK:
			result = 0 + 1
		case RESP_PAPER:
			result = 3 + 2
		case RESP_SCISSOR:
			result = 6 + 3
		}
	}

	if a == SCISSOR {
		switch b {
		case RESP_ROCK:
			result = 6 + 1
		case RESP_PAPER:
			result = 0 + 2
		case RESP_SCISSOR:
			result = 3 + 3
		}
	}
	return result
}