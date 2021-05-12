package pegchamp

import (
	"fmt"
	"strings"
)

// String takes `str` and only match that string one time.
func String(str string) Parser {
	return Parser{
		Func: func(ps ParserState) ParserState {
			if ps.err != nil {
				return ps
			}

			if ps.OutOfBound() {
				ps.err = fmt.Errorf("expected \"%s\" but found end of input", str)
				return ps
			}

			if strings.HasPrefix(ps.input[ps.idx:], str) {
				ps.result = str
				ps.idx += len(str)
			} else {
				ps.err = fmt.Errorf("expected \"%s\" but found \"%.16s\"", str, ps.input[ps.idx:])
			}

			return ps
		},
	}
}

// Char only matches one byte at a time. Works weirdly due to golang using byte as char.
func Char(c byte) Parser {
	return Parser{
		Func: func(ps ParserState) ParserState {
			if ps.err != nil {
				return ps
			}

			if ps.OutOfBound() {
				ps.err = fmt.Errorf("expected '%c' but found end of input", c)
				return ps
			}

			if ps.input[ps.idx] == c {
				ps.result = ps.input[ps.idx]
				ps.idx += 1
			} else {
				ps.err = fmt.Errorf("expected '%c' but found '%.16s'", c, ps.input[ps.idx:])
			}

			return ps
		},
	}
}
