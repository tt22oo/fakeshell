package parser

func Input(input string) []string {
	var (
		quote  rune
		tokens []string
	)

	buf := ""
	for i, w := range input {
		switch w {
		case ' ':
			if quote == '"' || quote == '\'' {
				buf += string(w)
				continue
			} else if buf == "" {
				continue
			}

			tokens = append(tokens, buf)
			buf = ""
		case ';', '>':
			if quote == '"' || quote == '\'' {
				buf += string(w)
				continue
			}

			if buf != "" {
				tokens = append(tokens, buf)
				buf = ""
			}
			tokens = append(tokens, string(w))
		case '&':
			if quote == '"' || quote == '\'' {
				buf += string(w)
				continue
			} else if i+1 == len([]rune(input)) {
				buf += string(w)
				tokens = append(tokens, buf)
				continue
			} else if input[i+1] == '&' {
				if buf != "" {
					tokens = append(tokens, buf)
				}
				buf = string(w)
				continue
			} else if buf == "&" {
				buf += string(w)
				tokens = append(tokens, buf)
				buf = ""
			}

		case '"', '\'':
			if quote == '"' && w == '"' {
				quote = 0
				tokens = append(tokens, buf)
				continue
			} else if quote == '\'' && w == '\'' {
				quote = 0
				tokens = append(tokens, buf)
				continue
			}

			if quote == 0 {
				quote = w
			} else {
				buf += string(w)
			}
		default:
			if i+1 == len([]rune(input)) {
				buf += string(w)
				tokens = append(tokens, buf)
			}

			buf += string(w)
			continue
		}
	}

	return tokens
}
