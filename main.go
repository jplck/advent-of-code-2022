package main

import (
	"os"
	"strconv"

	"github.com/jplck/advent-of-code-2022/challenge10_1"
	"github.com/jplck/advent-of-code-2022/challenge10_2"
	"github.com/jplck/advent-of-code-2022/challenge11"
	"github.com/jplck/advent-of-code-2022/challenge1_1"
	"github.com/jplck/advent-of-code-2022/challenge1_2"
	"github.com/jplck/advent-of-code-2022/challenge2_1"
	"github.com/jplck/advent-of-code-2022/challenge2_2"
	"github.com/jplck/advent-of-code-2022/challenge3_1"
	"github.com/jplck/advent-of-code-2022/challenge3_2"
	"github.com/jplck/advent-of-code-2022/challenge4_1"
	"github.com/jplck/advent-of-code-2022/challenge4_2"
	"github.com/jplck/advent-of-code-2022/challenge5_1"
	"github.com/jplck/advent-of-code-2022/challenge5_2"
	"github.com/jplck/advent-of-code-2022/challenge6_12"
	"github.com/jplck/advent-of-code-2022/challenge7_12"
	"github.com/jplck/advent-of-code-2022/challenge8_12"
	"github.com/jplck/advent-of-code-2022/challenge9_1"
	"github.com/jplck/advent-of-code-2022/utils"
)

func main() {

	switch os.Args[1] {
	case "1":
		challenge1_1.Run("./puzzles/input1")
		challenge1_2.Run("./puzzles/input1")
	case "2":
		challenge2_1.Run("./puzzles/input2")
		challenge2_2.Run("./puzzles/input2")
	case "3":
		challenge3_1.Run("./puzzles/input3")
		challenge3_2.Run("./puzzles/input3")
	case "4":
		challenge4_1.Run("./puzzles/input4")
		challenge4_2.Run("./puzzles/input4")
	case "5":
		challenge5_1.Run("./puzzles/input5")
		challenge5_2.Run("./puzzles/input5")
	case "6":
		challenge6_12.Run("./puzzles/input6")
	case "7":
		challenge7_12.Run("./puzzles/input7")
	case "8":
		challenge8_12.Run("./puzzles/input8")
	case "9":
		challenge9_1.Run("./puzzles/input9_1")
	case "10":
		challenge10_1.Run("./puzzles/input10")
		challenge10_2.Run("./puzzles/input10")
	case "11":
		rounds, err := strconv.Atoi(os.Args[2])
		utils.Must(err)
		challenge11.Run("./puzzles/input11", rounds)
	}

}
