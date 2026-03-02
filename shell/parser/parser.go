package parser

import "strings"

func Input(input string) []string {
	var (
		quote  rune
		tokens []string
		buf    strings.Builder
	)

	for i, w := range input {
		switch w {
		case ' ':
			if quote == '"' || quote == '\'' {
				buf.WriteRune(w)
				continue
			} else if buf.String() == "" {
				continue
			}

			tokens = append(tokens, buf.String())
			buf.Reset()
		case ';', '>':
			if quote == '"' || quote == '\'' {
				buf.WriteRune(w)
				continue
			}

			if buf.String() != "" {
				tokens = append(tokens, buf.String())
				buf.WriteRune(w)
				buf.Reset()
			}
			tokens = append(tokens, string(w))
		case '&':
			if quote == '"' || quote == '\'' {
				buf.WriteRune(w)
				continue
			} else if i+1 == len([]rune(input)) {
				buf.WriteRune(w)
				tokens = append(tokens, buf.String())
				continue
			} else if input[i+1] == '&' {
				if buf.String() != "" {
					tokens = append(tokens, buf.String())
				}
				buf.WriteRune(w)
				continue
			} else if buf.String() == "&" {
				buf.WriteRune(w)
				tokens = append(tokens, buf.String())
				buf.Reset()
			}

		case '"', '\'':
			if quote == '"' && w == '"' {
				quote = 0
				tokens = append(tokens, buf.String())
				continue
			} else if quote == '\'' && w == '\'' {
				quote = 0
				tokens = append(tokens, buf.String())
				continue
			}

			if quote == 0 {
				quote = w
			} else {
				buf.WriteRune(w)
			}
		default:
			if i+1 == len([]rune(input)) {
				buf.WriteRune(w)
				tokens = append(tokens, buf.String())
			}

			buf.WriteRune(w)
			continue
		}
	}

	return tokens
}
