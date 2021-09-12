# Structure :
### CommandExecutor :
Struct used to actually execute specific commands. 
Exposes one API : ExecuteCommand(cmd, args) -> This would accept the command alias string for example ls, mkdir etc as well as args list, and executes appropriate command
Will be initialized using its constructor which creates a FileSystem
`GetCommandExecutor`

It also takes in a type of output processor interface, in our case this is standard output printer


### Command :
Its an interface with one method : `Execute()`
Multiple commands inplement this interface :
1) `ListCommand` -> For ls
2) `MkDirCommand` -> For mkdir
3) `TouchCommand` -> For touch
4) `CdCommand` -> For cd

>More Can be added as and when required


### Services : FileSystem
Currently, exposed one Interface(service) for File System which exposes below APIs :
1. `Ls(isRecursive)`
2. `Mkdir(directoryName)`
3. `Pwd()`
4. `Cd(dirName)`
5. `Touch(fileName)`


### InMemoryFileSystem :
Contains core Data structure and business logic for our different functionalities.
Contains:
1) `currentPath` : Current Path of our users cmd prompt
2) `root` : Root Directory of type `Dir` representing the root of FS
3) `currDirectory` : Represents the current Directory corresponding to current path

`InMemoryFileSystem` uses the below `Dir` data model :
```
type Dir struct {
    parent *Dir
    dirMap map[string]Dir
    fileMap map[string]string
}
```

[Here, parent represents parent directory, dirMap represents the directories mapping from directory map in current directory, and fileMap represents mapping from file name to file content in the current directory]


`InMemoryFileSystem` Implements all the APIs exposed by the FileSystem Interface, whichever is of our concern


## Execution and commands :
Currently, the project only has defined interfaces and APIs. Any framework for creating command line tool, for example cobra can be used to generate cmd tools, and then call our APIs for required functionalities. 

They would have to instantiate `CommandExecutor` and then call the below API with cmd(string) and args(string array) :

```CommandExecutor.ExecuteCommand(cmd string, args []string)```


