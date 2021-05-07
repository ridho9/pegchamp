package pegchamp

// SequenceOf takes a list of parsers and applies them sequentially. Returning the result in an array.
func SequenceOf(parsers ...Parser) Parser {
	return Parser{
		Func: func(ps ParserState) ParserState {
			if ps.err != nil {
				return ps
			}

			totalResult := []interface{}{}
			currentState := ps

			for _, parser := range parsers {
				result := parser.Func(currentState)
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
		Func: func(ps ParserState) ParserState {
			if ps.err != nil {
				return ps
			}

			triedResult := []ParserState{}

			for _, parser := range parsers {
				res := parser.Func(ps)
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

// TakeSecond takes two parser and run them sequentially. Returns the result of the second parser.
// Effectively ignores the result of the first parser.
func TakeSecond(first Parser, second Parser) Parser {
	return Parser{
		Func: func(ps ParserState) ParserState {
			if ps.err != nil {
				return ps
			}

			res1 := first.Func(ps)
			if res1.err != nil {
				return res1
			}

			return second.Func(res1)
		},
	}
}
