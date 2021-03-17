package pegchamp

import (
	"fmt"
	"strings"
)

// String takes `str` and only match that string one time.
func String(str string) Parser {
	return Parser{
		f: func(ps ParserState) ParserState {
			if ps.err != nil {
				return ps
			}

			if strings.HasPrefix(ps.input[ps.idx:], str) {
				ps.result = str
				ps.idx += len(str)
			} else {
				ps.err = fmt.Errorf("expected \"%s\" but found \"%.16s\"'", str, ps.input)
			}

			return ps
		},
	}
}
