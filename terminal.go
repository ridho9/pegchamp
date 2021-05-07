package pegchamp

import "fmt"

// Alpha takes a single alphabetical `/[a-zA-Z]/`` character and return it in a string.
func Alpha() Parser {
	return Parser{
		Func: func(ps ParserState) ParserState {
			if ps.err != nil {
				return ps
			}

			if ps.outOfBound() {
				ps.err = fmt.Errorf("expected alphabetical but found end of input")
				return ps
			}

			head := ps.input[ps.idx]
			insideLower := ('a' <= head) && (head <= 'z')
			insideUpper := ('A' <= head) && (head <= 'Z')
			if insideLower || insideUpper {
				ps.result = ps.input[ps.idx : ps.idx+1]
				ps.idx += 1
				return ps
			}

			ps.err = fmt.Errorf("expected alphabetical but found \"%.16s\"", ps.input[ps.idx:])
			return ps
		},
	}
}
