package models

type Dir struct {
	parent  *Dir
	dirMap  map[string]Dir
	fileMap map[string]string
}

func NewDir(parent *Dir) Dir {
	return Dir{
		parent:  parent,
		dirMap:  make(map[string]Dir),
		fileMap: make(map[string]string),
	}
}

func (this *Dir) GetParent() *Dir {
	return this.parent
}

func (this *Dir) GetCurrentDirectoriesMap() map[string]Dir {
	return this.dirMap
}

func (this *Dir) GetCurrentFilesMap() map[string]string {
	return this.fileMap
}
