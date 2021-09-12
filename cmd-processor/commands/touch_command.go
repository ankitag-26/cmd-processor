package commands

import (
	"github.com/cmd-processor/output_processor"
	"github.com/cmd-processor/services"
)

type TouchCommand struct {
	fileSystem      services.FileSystem
	fileName        string
	outputProcessor output_processor.OutputProcessor
}

func GetTouchCommand(fs services.FileSystem, processor output_processor.OutputProcessor, file string) TouchCommand {
	return TouchCommand{fileSystem: fs, fileName: file, outputProcessor: processor}
}

func (cmd TouchCommand) Execute() {
	err := cmd.fileSystem.Touch(cmd.fileName)
	if err != nil {
		cmd.outputProcessor.ProcessError(err)
	}
}
