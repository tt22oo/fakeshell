package command

import "github.com/tt22oo/fakeshell/shell"

func ls(s *shell.Shell, cmd []string) (string, int) {
	var output string
	for name := range s.Entry.Children {
		output += name + " "
	}
	output += "\r\n"

	return output, 0
}
