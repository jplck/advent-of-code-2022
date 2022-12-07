package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const FOLDER_SIZE_THRESHOLD = 100000
const TOTAL_SPACE_AVAILABLE = 70000000
const SPACE_REQUIRED_FOR_UPDATE = 30000000

const PREFIX_DIR = "dir "
const PREFIX_LS = "$ ls"
const PREFIX_CD = "$ cd "

func main() {
	fileHandle, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(fileHandle)
	fileScanner.Split(bufio.ScanLines)

	root := Dir{
		Name: "root",
	}
	root.AddDir(&Dir{
		Name: "/",
	})

	currentDir := &root
	for fileScanner.Scan() {
		readInput := fileScanner.Text()

		prefix, rowValue := parseInputRow(readInput)

		switch prefix {
		case PREFIX_LS:
			continue
		case PREFIX_CD:
			if rowValue == ".." {
				currentDir = currentDir.Parent
				continue
			}
			currentDir = currentDir.Cd(rowValue)
		case PREFIX_DIR:
			newDir := Dir{
				Name: rowValue,
			}
			currentDir.AddDir(&newDir)
		default:
			fileName, fileSize := parseFile(readInput)
			newFile := File{
				Name: fileName,
				Size: fileSize,
			}
			currentDir.AddFile(&newFile)
		}
	}
	//challenge 7.1
	//sum := SumDirs(root.Dirs["/"].Dirs)
	//fmt.Println(sum)

	//Challenge 7.2
	sizes := FindFolderSizesAvailableForDeletion(root)
	fmt.Println(Min(sizes))
}

func FindFolderSizesAvailableForDeletion(root Dir) []int {
	totalSpace := TOTAL_SPACE_AVAILABLE
	usedSpace := root.Size()
	freeSpace := totalSpace - usedSpace
	additionalSpaceRequiredForUpdate := SPACE_REQUIRED_FOR_UPDATE - freeSpace

	availableDirSizes := FindDirToDelete(
		additionalSpaceRequiredForUpdate,
		root.Dirs,
	)

	return availableDirSizes
}

func FindDirToDelete(requiredSize int, dirs map[string]*Dir) []int {
	availableDirSizes := make([]int, 0)
	for _, dir := range dirs {
		if dir.Size() >= requiredSize {
			availableDirSizes = append(availableDirSizes, dir.Size())
		}
		availableDirSizes = append(
			availableDirSizes,
			FindDirToDelete(requiredSize, dir.Dirs)...,
		)
	}
	return availableDirSizes
}

func SumDirSizesBelowThreshold(dirs map[string]*Dir) int {
	sum := 0
	for _, dir := range dirs {
		dirSize := dir.Size()
		if dirSize <= FOLDER_SIZE_THRESHOLD {
			sum += dir.Size()
		}
		sum += SumDirSizesBelowThreshold(dir.Dirs)
	}
	return sum
}

func parseInputRow(row string) (prefix string, value string) {

	types := []string{
		PREFIX_CD,
		PREFIX_DIR,
		PREFIX_LS,
	}

	for _, rowType := range types {
		if strings.HasPrefix(row, rowType) {
			prefix = rowType
			value = strings.Replace(row, prefix, "", 1)
			break
		}
	}
	return
}

func parseFile(row string) (fileName string, fileSize int) {
	fileComponents := strings.Split(row, " ")
	fileName = fileComponents[1]
	fileSize, err := strconv.Atoi(fileComponents[0])
	must(err)
	return
}
