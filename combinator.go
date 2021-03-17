package pegchamp

// SequenceOf takes a list of parsers and applies them sequentially. Returning the result in an array.
func SequenceOf(parsers ...Parser) Parser {
	return Parser{
		f: func(ps ParserState) ParserState {
			if ps.err != nil {
				return ps
			}

			totalResult := []interface{}{}
			currentState := ps

			for _, parser := range parsers {
				result := parser.f(currentState)
				if result.err != nil {
					currentState.err = result.err
					break
				}

				totalResult = append(totalResult, result.Result())
				currentState = result
			}

			currentState.result = totalResult
			return currentState
		},
	}
}
