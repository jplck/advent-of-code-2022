package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const MSG_LEN = 4

func main() {
	fileHandle, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(fileHandle)
	fileScanner.Split(bufio.ScanRunes)

	readMsg := ""
	markerCount := 0

out:
	for fileScanner.Scan() {

		markerCount++
		readInputChar := fileScanner.Text()

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
