package pegchamp

import (
	"fmt"
	"strings"
)

// String takes `str` and only match that string one time.
func String(str string) Parser {
	return func(ps ParserState) ParserState {
		if ps.err != nil {
			return ps
		}

		if strings.HasPrefix(ps.input[ps.idx:], str) {
			ps.result = str
			ps.idx += len(str)
		} else {
			ps.err = fmt.Errorf("expected '%s' but found '%.16s'", str, ps.input)
		}

		return ps
	}
}

// SequenceOf takes a list of parsers and applies them sequentially. Returning the result in an array.
func SequenceOf(parsers ...Parser) Parser {
	return func(ps ParserState) ParserState {
		if ps.err != nil {
			return ps
		}

		totalResult := []interface{}{}
		currentState := ps

		for _, parser := range parsers {
			result := parser(currentState)
			if result.err != nil {
				currentState.err = result.err
				break
			}

			totalResult = append(totalResult, result.Result())
			currentState = result
		}

		currentState.result = totalResult
		return currentState
	}
}
