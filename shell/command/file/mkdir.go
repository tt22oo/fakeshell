package file

import "github.com/tt22oo/fakeshell/shell"

const (
	mkdirMissing string = "mkdir: missing operand\r\nTry 'mkdir --help' for more information.\r\n"
)

func Mkdir(s *shell.Shell, cmd []string) (string, int) {
	if len(cmd) < 2 {
		return mkdirMissing, 1
	}

	for _, name := range cmd[1:] {
		s.Dir.Entry.NewDirectory(name)
	}

	return "", 0
}
