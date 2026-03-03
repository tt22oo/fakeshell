package file

import (
	"strings"

	"github.com/tt22oo/fakeshell/shell"
)

func Ls(s *shell.Shell, cmd []string) (string, int) {
	var output string
	if len(cmd) == 1 {
		for name := range s.Dir.Entry.Children {
			if strings.HasPrefix(name, ".") {
				continue
			}
			output += name + " "
		}
	} else if len(cmd) > 1 {
		switch cmd[1] {
		case "-a":
			for name := range s.Dir.Entry.Children {
				output += name + " "
			}
		}
	}

	output += "\r\n"
	return output, 0
}
