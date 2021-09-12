package commands

import (
	"github.com/cmd-processor/output_processor"
	"github.com/cmd-processor/services"
)

type CdCommand struct {
	fileSystem      services.FileSystem
	directoryName   string
	outputProcessor output_processor.OutputProcessor
}

func GetCdCommand(fs services.FileSystem, processor output_processor.OutputProcessor, dirName string) CdCommand {
	return CdCommand{fileSystem: fs, directoryName: dirName, outputProcessor: processor}
}

func (cmd CdCommand) Execute() {
	err := cmd.fileSystem.Cd(cmd.directoryName)
	if err != nil {
		cmd.outputProcessor.ProcessError(err)
	}
}
