package file

import (
	"fmt"

	"github.com/tt22oo/fakeshell/shell"
)

const (
	cdNoSuch  string = "-bash: cd: %s: No such file or directory\r\n"
	cdToomany string = "-bash: cd: too many arguments\r\n"
)

func Cd(s *shell.Shell, cmd []string) (string, int) {
	switch len(cmd) {
	case 1:
		entry, _ := s.Dir.GetEntry(s.Dir.HomePath)
		s.Dir.Entry = entry
	case 2:
		entry, err := s.Dir.GetEntry(cmd[1])
		if err != nil {
			return fmt.Sprintf(cdNoSuch, cmd[1]), 1
		}

		s.Dir.Entry = entry
	default:
		return cdToomany, 1
	}

	return "", 0
}
