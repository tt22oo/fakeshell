package command

import "github.com/tt22oo/fakeshell/shell"

func touch(s *shell.Shell, cmd []string) (string, int) {
	if len(cmd) < 2 {
		return "Try 'touch --help' for more information.\r\n", 1
	}

	for _, name := range cmd[1:] {
		s.Entry.NewFile(name, "")
	}

	return "", 0
}
