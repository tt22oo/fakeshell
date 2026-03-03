## Fake Shell
Supported Commands:
`ls`, `cd`, `touch`, `mkdir`

## Example
```go
package main

import (
	"fmt"
	"log"

	"github.com/tt22oo/fakeshell/loader"
	"github.com/tt22oo/fakeshell/shell"
	"github.com/tt22oo/fakeshell/shell/command"
	"github.com/tt22oo/fakeshell/shell/parser"
)

func main() {
	// initialize shell configuration
	cfg := &shell.Config{
		HomePath: "/root",
		DirPath:  "configs/dir.json",
		Proc: &loader.ProcConfig{
			ProcessPath: "configs/proc/process.json",
			CpuInfoPath: "configs/proc/cpuinfo.txt",
			MemInfoPath: "configs/proc/meminfo.txt",
			VersionPath: "configs/proc/version.txt",
		},
	}

	// create a new shell
	s, err := cfg.New()
	if err != nil {
		log.Printf("Shell Error: %s", err.Error())
		return
	}

	// input string (commands)
	input := "touch a; ls"
	tokens := parser.Input(input)             // tokenize input
	fmt.Print(command.RunCommands(s, tokens)) // run commands
}
```
Output:
`test1 test2`
## Data Format
**Directory JSON:**
```json
{
    "name": {
        "name": "name",
        "type": "directory",
        "children": {
            "name": {
                "name": "name",
                "type": "file",
                "data": "data"
            }
        }
    }
}
```
**Process JSON:** 
```json
{
    "pid": {
        "user": "user",
        "cmd": "command",
        "args": ["arg1", "arg2"]
    }
}
```
