package file

import "github.com/tt22oo/fakeshell/shell"

const (
	touchMissing string = "touch: missing file operand\r\nTry 'touch --help' for more information.\r\n"
)

func Touch(s *shell.Shell, cmd []string) (string, int) {
	if len(cmd) < 2 {
		return touchMissing, 1
	}

	for _, name := range cmd[1:] {
		s.Dir.Entry.NewFile(name, "")
	}

	return "", 0
}
