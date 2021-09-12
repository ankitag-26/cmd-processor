package commands

import (
	"github.com/cmd-processor/output_processor"
	"github.com/cmd-processor/services"
)

type ListCommand struct {
	fileSystem      services.FileSystem
	isRecursive     bool
	outputProcessor output_processor.OutputProcessor
}

func GetListCommand(fs services.FileSystem, processor output_processor.OutputProcessor, isRecursive bool) ListCommand {
	return ListCommand{fileSystem: fs, isRecursive: isRecursive, outputProcessor: processor}
}

func (cmd ListCommand) Execute() {
	results, err := cmd.fileSystem.Ls(cmd.isRecursive)
	if err != nil {
		cmd.outputProcessor.ProcessError(err)
	} else {
		for _, path := range results {
			cmd.outputProcessor.ProcessOutput(path)
		}
	}
}
