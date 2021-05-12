package pegchamp

// SequenceOf takes a list of parsers and applies them sequentially. Returning the result in an array.
func SequenceOf(parsers ...Parser) Parser {
	return Parser{
		parserFunc: func(ps ParserState) ParserState {
			totalResult := []interface{}{}
			currentState := ps

			for _, parser := range parsers {
				currentState = parser.Parse(currentState)
				if currentState.err != nil {
					break
				}

				totalResult = append(totalResult, currentState.Result())
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
		parserFunc: func(ps ParserState) ParserState {
			triedResult := []ParserState{}

			for _, parser := range parsers {
				res := parser.Parse(ps)
				if res.err == nil {
					return res
				}
				triedResult = append(triedResult, res)
			}

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
		parserFunc: func(ps ParserState) ParserState {
			res1 := first.Parse(ps)
			if res1.err != nil {
				return res1
			}

			return second.Parse(res1)
		},
	}
}

// Many will run `parser` for 0 or more times until it errors, and accumulate it in an array.
func Many(parser Parser) Parser {
	return Parser{
		parserFunc: func(ps ParserState) ParserState {
			totalResult := []interface{}{}
			currentState := ps

			for {
				currentState = parser.Parse(currentState)
				if currentState.err != nil {
					currentState.err = nil
					break
				}

				totalResult = append(totalResult, currentState.Result())
			}

			currentState.result = totalResult
			return currentState
		},
	}
}

// Many1 will run `parser` for 1 or more times until it errors, and accumulate it in an array.
// Not matching any will returns the first error.
func Many1(parser Parser) Parser {
	return Parser{
		parserFunc: func(ps ParserState) ParserState {
			totalResult := []interface{}{}
			currentState := ps

			for {
				currentState = parser.Parse(currentState)
				if currentState.err != nil {
					if len(totalResult) > 0 {
						currentState.err = nil
					}
					break
				}

				totalResult = append(totalResult, currentState.Result())
			}

			currentState.result = totalResult
			return currentState
		},
	}
}

// Optional will tries `parser`, returning `nil` when fails.
func Optional(parser Parser) Parser {
	return Parser{
		parserFunc: func(ps ParserState) ParserState {
			res := parser.Parse(ps)
			if res.err != nil {
				res = ps
				res.result = nil
			}
			return res
		},
	}
}
