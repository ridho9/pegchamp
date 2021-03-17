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

// Choice takes a list of parsers and returns the result of first parsers that match.
// When all parsers fail, returns the result of the parser that matched the longest.
func Choice(parsers ...Parser) Parser {
	return Parser{
		f: func(ps ParserState) ParserState {
			if ps.err != nil {
				return ps
			}

			triedResult := []ParserState{}

			for _, parser := range parsers {
				res := parser.f(ps)
				if res.err == nil {
					return res
				}
				triedResult = append(triedResult, res)
			}

			// all
			longest := ParserState{
				idx: -1,
			}
			for _, res := range triedResult {
				if res.idx > longest.idx {
					longest = res
				}
			}

			return longest
		},
	}
}
