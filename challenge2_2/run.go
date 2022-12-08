package challenge2_2

import (
	"fmt"
	"strings"

	"github.com/jplck/advent-of-code-2022/utils"
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

func Run(inputFile string) {
	reader := utils.GetInputReader(inputFile, utils.Lines)

	sum := 0

	for reader.Scan() {
		t := reader.Text()
		parts := strings.Split(t, " ")
		a := parts[0]
		b := parts[1]

		if b == LOOSE {
			if a == ROCK {
				b = RESP_SCISSOR
			} else if a == PAPER {
				b = RESP_ROCK
			} else if a == SCISSOR {
				b = RESP_PAPER
			}
		} else if b == DRAW {
			if a == ROCK {
				b = RESP_ROCK
			} else if a == PAPER {
				b = RESP_PAPER
			} else if a == SCISSOR {
				b = RESP_SCISSOR
			}
		} else if b == WIN {
			if a == ROCK {
				b = RESP_PAPER
			} else if a == PAPER {
				b = RESP_SCISSOR
			} else if a == SCISSOR {
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
