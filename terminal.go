package pegchamp

import (
	"fmt"
	"strings"
)

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
			if isByteAlpha(head) {
				ps.result = ps.input[ps.idx : ps.idx+1]
				ps.idx += 1
				return ps
			}

			ps.err = fmt.Errorf("expected alphabetical but found \"%.16s\"", ps.input[ps.idx:])
			return ps
		},
	}
}

// Alphas takes more than one alphabetical `/[a-zA-Z]+/` characters and return it in a string.
func Alphas() Parser {
	return Parser{
		Func: func(ps ParserState) ParserState {
			if ps.err != nil {
				return ps
			}

			if ps.outOfBound() {
				ps.err = fmt.Errorf("expected alphabetical but found end of input")
				return ps
			}

			builder := strings.Builder{}
			for ps.idx >= len(ps.input) {
				head := ps.input[ps.idx]
				if !isByteAlpha(head) {
					break
				}

				builder.WriteByte(head)
				ps.idx += 1
			}

			ps.result = builder.String()
			if ps.result == "" {
				ps.err = fmt.Errorf("expected alphabetical but found \"%.16s\"", ps.input[ps.idx:])
			}
			return ps
		},
	}
}

func isByteAlpha(b byte) bool {
	insideLower := ('a' <= b) && (b <= 'z')
	insideUpper := ('A' <= b) && (b <= 'Z')
	return insideLower || insideUpper
}
