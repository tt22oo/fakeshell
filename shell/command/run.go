package command

import (
	"fmt"

	"github.com/tt22oo/fakeshell/shell"
	"github.com/tt22oo/fakeshell/shell/command/file"
	"github.com/tt22oo/fakeshell/shell/proc"
)

type command func(s *shell.Shell, cmd []string) (string, int)

var commands = map[string]command{
	"ls":    file.Ls,
	"cd":    file.Cd,
	"touch": file.Touch,
	"mkdir": file.Mkdir,
}

func runCommand(s *shell.Shell, cmd []string) (string, int) {
	if len(cmd) == 0 {
		return "", 0
	}

	run := commands[cmd[0]]
	if run == nil {
		return fmt.Sprintf("%s: command not found\r\n", cmd[0]), 127
	}

	process := &proc.Process{
		User: s.User,
		Cmd:  cmd[0],
		Args: cmd[1:],
	}
	pid := process.New(&s.Proc.Table)
	defer s.Proc.Table.Delete(pid)

	return run(s, cmd)
}

func RunCommands(s *shell.Shell, tokens []string) string {
	var (
		outputs string
		command []string
	)

	for i, token := range tokens {
		switch token {
		case ";":
			result, _ := runCommand(s, command)
			outputs += result
			command = nil
		case "&&":
			result, stat := runCommand(s, command)
			outputs += result
			command = nil

			if stat != 0 {
				return outputs
			}
		default:
			if token != "" {
				if command == nil {
					command = []string{}
				}
				command = append(command, token)
			}

			if i == len(tokens)-1 {
				result, _ := runCommand(s, command)
				outputs += result
			}
		}
	}

	return outputs
}
