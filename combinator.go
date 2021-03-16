package pegchamp

import (
	"fmt"
	"strings"
)

func String(str string) Parser {
	return func(state ParserState) ParserState {
		if state.err != nil {
			return state
		}

		if strings.HasPrefix(state.input, str) {
			state.result = str
			state.input = state.input[len(str):]
			state.matchedLen += len(str)
		} else {
			state.err = fmt.Errorf("expected '%s' but found '%.16s'", str, state.input)
		}

		return state
	}
}

func SequenceOf(parsers ...Parser) Parser {
	return func(ps ParserState) ParserState {
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
