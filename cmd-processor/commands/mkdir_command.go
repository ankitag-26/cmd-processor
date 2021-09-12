package commands

import (
	"github.com/cmd-processor/output_processor"
	"github.com/cmd-processor/services"
)

type MkDirCommand struct {
	fileSystem      services.FileSystem
	directoryName   string
	outputProcessor output_processor.OutputProcessor
}

func GetMkDirCommand(fs services.FileSystem, processor output_processor.OutputProcessor, dirName string) MkDirCommand {
	return MkDirCommand{fileSystem: fs, directoryName: dirName, outputProcessor: processor}
}

func (cmd MkDirCommand) Execute() {
	err := cmd.fileSystem.Mkdir(cmd.directoryName)
	if err != nil {
		cmd.outputProcessor.ProcessError(err)
	}
}
