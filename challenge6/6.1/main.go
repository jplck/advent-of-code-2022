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

		if len(readMsg) >= MSG_LEN && FindOccurences(readMsg) == 0 {
			break out
		}
	}
	fmt.Println(markerCount)
}

func FindOccurences(message string) int {
	readMsgLen := len(message)
	msgStartIdx := readMsgLen - MSG_LEN
	occurenceCount := 0

	for i := msgStartIdx; i < readMsgLen; i++ {
		stringToEvaluate := message[msgStartIdx:]
		char := message[i : i+1]
		occurenceCount += strings.Count(stringToEvaluate, char) - 1
	}
	return occurenceCount
}
