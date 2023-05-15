package dotenv

import (
	"os"
)

// Reads your env file and loads it into ENV for this process.
//
//	It overrides existing variables
func Load(filename string) error {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	var startName, endName, startValue, pos int
	var parsingName = true
	var insideQuotes = false
	var quoteType byte

	for pos < len(buf) {
		if parsingName {
			startName = pos

			for pos < len(buf) && buf[pos] != '=' {
				pos++
			}

			endName = pos
			parsingName = false
			pos++ // skip '='
		} else {
			if buf[pos] == '"' || buf[pos] == '\'' {
				insideQuotes = true
				quoteType = buf[pos]
				pos++ // skip quote
			}

			startValue = pos

			for pos < len(buf) {
				if insideQuotes {
					if buf[pos] == quoteType {
						break
					}
				} else if buf[pos] == '\n' || buf[pos] == '\r' {
					break
				}

				pos++
			}

			os.Setenv(string(buf[startName:endName]), string(buf[startValue:pos]))

			pos++ // skip end byte

			// skip whitespace bytes
			for pos < len(buf) && isWhiteSpace(buf[pos]) {
				pos++
			}

			parsingName = true
			insideQuotes = false
		}
	}

	return nil
}

func isWhiteSpace(b byte) bool {
	switch b {
	case '\t', '\n', '\v', '\f', '\r', ' ':
		return true
	default:
		return false
	}
}
