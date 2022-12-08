package challenge6_12

import (
	"fmt"
	"strings"

	"github.com/jplck/advent-of-code-2022/utils"
)

const MSG_LEN = 4

func Run(inputFile string) {
	reader := utils.GetInputReader(inputFile, utils.Chars)

	readMsg := ""
	markerCount := 0

out:
	for reader.Scan() {

		markerCount++
		readInputChar := reader.Text()

		readMsg = fmt.Sprintf("%v%v", readMsg, readInputChar)

		if len(readMsg) >= MSG_LEN && !HasDuplicates(readMsg) {
			break out
		}
	}
	fmt.Println(markerCount)
}

func HasDuplicates(message string) bool {
	readMsgLen := len(message)
	msgStartIdx := readMsgLen - MSG_LEN

	for posInMsg := msgStartIdx; posInMsg < readMsgLen; posInMsg++ {
		stringToEvaluate := message[msgStartIdx:]
		char := message[posInMsg : posInMsg+1]
		if strings.Count(stringToEvaluate, char)-1 > 0 {
			return true
		}
	}
	return false
}
