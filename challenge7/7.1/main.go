package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
		if readInput == "$ ls" {
			continue
		}

		if doCd, dirName := ChangeDir(readInput); doCd {
			if dirName == ".." {
				currentDir = currentDir.Parent
				continue
			}
			currentDir = currentDir.Cd(dirName)
			continue
		}

		if isDir, dirName := IsDir(readInput); isDir {
			newDir := Dir{
				Name: dirName,
			}
			currentDir.AddDir(&newDir)
			continue
		}

		fileName, fileSize := ParseFile(readInput)
		newFile := File{
			Name: fileName,
			Size: fileSize,
		}
		currentDir.AddFile(&newFile)

	}
	//challenge 7.1
	//sum := SumDirs(root.Dirs["/"].Dirs)
	//fmt.Println(sum)

	//Challenge 7.2
	sizes := FreeUpSpace(root)
	fmt.Println(Min(sizes))
}

func FreeUpSpace(root Dir) []int {
	totalSpace := 70000000
	usedSpace := root.Size()
	freeSpace := totalSpace - usedSpace
	additionalSpaceRequiredForUpdate := 30000000 - freeSpace

	availableDirSizes := FindDirToDelete(additionalSpaceRequiredForUpdate, root.Dirs)
	return availableDirSizes
}

func FindDirToDelete(requiredSize int, dirs map[string]*Dir) []int {
	availableDirSizes := make([]int, 0)
	for _, dir := range dirs {
		if dir.Size() >= requiredSize {
			availableDirSizes = append(availableDirSizes, dir.Size())
		}
		availableDirSizes = append(availableDirSizes, FindDirToDelete(requiredSize, dir.Dirs)...)
	}
	return availableDirSizes
}

func SumDirs(dirs map[string]*Dir) int {
	sum := 0
	for _, dir := range dirs {
		size := dir.Size()
		if size <= 100000 {
			sum += dir.Size()
		}
		sum += SumDirs(dir.Dirs)
	}
	return sum
}

func ChangeDir(row string) (doCd bool, dirName string) {
	prefix := "$ cd "
	if strings.HasPrefix(row, prefix) {
		doCd = true
		dirName = strings.Replace(row, prefix, "", 1)
	}

	return
}

func IsDir(row string) (isDir bool, dirName string) {
	prefix := "dir "
	if strings.HasPrefix(row, prefix) {
		isDir = true
		dirName = strings.Replace(row, prefix, "", 1)
	}
	return
}

func ParseFile(row string) (fileName string, fileSize int) {
	fileComponents := strings.Split(row, " ")
	fileName = fileComponents[1]
	fileSize, err := strconv.Atoi(fileComponents[0])
	must(err)
	return
}
