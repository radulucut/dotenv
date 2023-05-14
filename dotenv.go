package dotenv

import (
	"os"
)

func Load(filename string) error {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	var startName, endName, startValue, pos int
	var parsingName = true
	var insideQuotes = false
	var escapeMap = map[int]bool{}

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
			if buf[pos] == '\\' {
				escapeMap[pos+1] = true
				pos++ // skip escape byte
			} else if buf[pos] == '"' {
				insideQuotes = true
				pos++ // skip '='
			}

			startValue = pos

			for pos < len(buf) {
				if !escapeMap[pos] && buf[pos] == '\\' {
					escapeMap[pos+1] = true
				} else if insideQuotes {
					if !escapeMap[pos] && buf[pos] == '"' {
						break
					}
				} else if buf[pos] == '\n' {
					break
				}

				pos++
			}

			// TODO: Remove escape bytes

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

// returns new start pos for the value
// func shiftEscapeBytes(buf *[]byte, escapeMap *map[uint]bool, start uint, end uint)  uint {
// 	for start < end {
// 		if escapeMap[end] {
//
// 		}
// 	}
// }

func isWhiteSpace(b byte) bool {
	switch b {
	case '\t', '\n', '\v', '\f', '\r', ' ':
		return true
	default:
		return false
	}
}
