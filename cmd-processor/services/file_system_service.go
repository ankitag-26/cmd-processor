package services

type FileSystem interface {
	Ls(isRecursive bool) ([]string, error)
	Mkdir(directoryName string) error
	Pwd() (string, error)
	Cd(dirName string) error
	Touch(fileName string) error
}
