package main

func (f *Dir) Size() int {
	sizeInDir := 0
	for _, file := range f.Files {
		sizeInDir += file.Size
	}

	for _, dir := range f.Dirs {
		sizeInDir += dir.Size()
	}
	return sizeInDir
}

func (f *Dir) AddDir(dir *Dir) {
	if f.Dirs == nil {
		f.Dirs = make(map[string]*Dir)
	}
	dir.Parent = f
	f.Dirs[dir.Name] = dir
}

func (f *Dir) AddFile(file *File) {
	file.Parent = f
	f.Files = append(f.Files, file)
}

func (f *Dir) Cd(dirName string) *Dir {
	if newDir, ok := f.Dirs[dirName]; ok {
		return newDir
	}
	return nil
}

func (f *Fs) Cd(dirName string) {
	f.CurrentDir = f.CurrentDir.Dirs[dirName]
}

func (f *Fs) CdUp() {
	f.CurrentDir = f.CurrentDir.Parent
}
