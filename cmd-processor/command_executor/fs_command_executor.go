package command_executor

import (
	"errors"
	"os"

	"github.com/cmd-processor/commands"
	"github.com/cmd-processor/file_systems"
	"github.com/cmd-processor/output_processor"
	"github.com/cmd-processor/services"
)

// CommandExecutor defines the main struct which would be used to call a particular command
type CommandExecutor struct {
	commands.Command
	fileSystem services.FileSystem
}

func GetCommandExecutor() *CommandExecutor {
	fs := file_systems.Constructor()
	return &CommandExecutor{fileSystem: fs}
}

func (ex *CommandExecutor) ExecuteCommand(cmd string, args []string) {
	processor := output_processor.GetStdOutputPrinter()
	switch cmd {
	case "ls":
		{
			isRecursive := len(args) > 0 && args[0] == "r"
			ex.Command = commands.GetListCommand(ex.fileSystem, processor, isRecursive)
			ex.Execute()
		}
	case "mkdir":
		{
			dirName := ""
			if len(args) > 0 {
				dirName = args[0]
			}
			ex.Command = commands.GetMkDirCommand(ex.fileSystem, processor, dirName)
			ex.Execute()
		}
	case "cd":
		{
			dirName := ""
			if len(args) > 0 {
				dirName = args[0]
			}
			ex.Command = commands.GetCdCommand(ex.fileSystem, processor, dirName)
			ex.Execute()
		}
	case "touch":
		{
			fileName := ""
			if len(args) > 0 {
				fileName = args[0]
			}
			ex.Command = commands.GetTouchCommand(ex.fileSystem, processor, fileName)
			ex.Execute()
		}
	case "quit":
		{
			os.Exit(0)
		}
	default:
		processor.ProcessError(errors.New("Incorrect Command" + cmd))
	}
}
