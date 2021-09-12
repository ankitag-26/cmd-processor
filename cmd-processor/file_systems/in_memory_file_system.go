package file_systems

import (
	"errors"
	"sort"
	"strings"

	"github.com/cmd-processor/models"
)

type InMemoryFileSystem struct {
	currentPath   string
	root          models.Dir
	currDirectory models.Dir
}

func Constructor() *InMemoryFileSystem {
	fs := &InMemoryFileSystem{
		root:        models.NewDir(nil),
		currentPath: "/",
	}
	fs.currDirectory = fs.root
	return fs
}

func (this InMemoryFileSystem) getCurrentDir() models.Dir {
	return this.currDirectory
}

func (this InMemoryFileSystem) Ls(isRecursive bool) ([]string, error) {
	if !isRecursive {
		return lsInPath(this.root, this.currentPath)
	}
	results := make([]string, 0)
	err := listRecursively(this.root, this.currentPath, &results)
	return results, err
}

func listRecursively(root models.Dir, path string, results *[]string) error {
	paths := strings.Split(path, "/")
	for i := 1; i < len(paths); i++ {
		currPath := "/"
		for j := 1; j < i; j++ {
			currPath += paths[j]
		}
		currResultsList, err := lsInPath(root, currPath)
		if err != nil {
			return err
		}
		*results = append(*results, currResultsList...)
	}
	return nil
}

func lsInPath(root models.Dir, path string) ([]string, error) {
	curr := root
	files := []string{}
	if path != "/" {
		d := strings.Split(path, "/")
		for i := 1; i < len(d)-1; i++ {
			curr = curr.GetCurrentDirectoriesMap()[d[i]]
		}
		if _, ok := curr.GetCurrentFilesMap()[d[len(d)-1]]; ok {
			files = append(files, d[len(d)-1])
			return files, nil
		} else {
			curr = curr.GetCurrentDirectoriesMap()[d[len(d)-1]]
		}
	}
	for file, _ := range curr.GetCurrentFilesMap() {
		files = append(files, file)
	}
	for dir, _ := range curr.GetCurrentDirectoriesMap() {
		files = append(files, dir)
	}
	sort.Strings(files)
	return files, nil
}

func (this InMemoryFileSystem) Mkdir(directoryName string) error {
	curr := this.root
	path := this.currentPath + "/" + directoryName
	d := strings.Split(path, "/")
	for i := 1; i < len(d); i++ {
		if _, ok := curr.GetCurrentDirectoriesMap()[d[i]]; !ok {
			curr.GetCurrentDirectoriesMap()[d[i]] = models.NewDir(&curr)
		}
		curr = curr.GetCurrentDirectoriesMap()[d[i]]
	}
	return nil
}

func (this InMemoryFileSystem) Pwd() (string, error) {
	return this.currentPath, nil
}

func (this InMemoryFileSystem) Cd(dirName string) error {
	if this.currentPath == "/" || dirName == "." {
		return nil
	}
	currentPathSplit := strings.Split(this.currentPath, "/")
	if dirName == ".." {
		currentPathSplit = currentPathSplit[:len(currentPathSplit)-1]
		this.currentPath = strings.Join(currentPathSplit, "/")
		this.currDirectory = *this.currDirectory.GetParent()
		return nil
	}
	currDir := this.getCurrentDir()
	if _, ok := currDir.GetCurrentDirectoriesMap()[dirName]; !ok {
		return errors.New("Directory Not Found")
	} else {
		this.currentPath += "/" + dirName
		this.currDirectory = currDir.GetCurrentDirectoriesMap()[dirName]
	}
	return nil
}

func (this InMemoryFileSystem) Touch(fileName string) error {
	this.currDirectory.GetCurrentFilesMap()[fileName] = ""
	return nil
}
