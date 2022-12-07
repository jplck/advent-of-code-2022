package main

type IFs interface {
	Cd(dirName string)
	CdUp()
}

type Fs struct {
	CurrentDir *Dir
}

type IDir interface {
	List()
	Size() int
	AddFile(file *File)
	AddDir(dir *Dir)
	Cd(dirName string)
	CdUp()
}

type Dir struct {
	Parent *Dir
	Name   string
	Dirs   map[string]*Dir
	Files  []*File
}

type File struct {
	Name   string
	Parent *Dir
	Size   int
}
