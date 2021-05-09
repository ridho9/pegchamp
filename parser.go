package pegchamp

type ParserState struct {
	input  string
	idx    int
	result interface{}
	err    error
}

func (ps ParserState) outOfBound() bool {
	return ps.idx >= len(ps.input)
}

func (p ParserState) Result() interface{} {
	return p.result
}

func (p ParserState) Error() error {
	return p.err
}

// type Parser func(ParserState) ParserState
type Parser struct {
	Func func(ParserState) ParserState
}

// Map takes a mapper function that takes the parser result and return new result.
// Other part of the ParserState won't be changed.
// The mapper function won't be run in the case of an error
// so it could be expected only a successful result is passed.
func (p Parser) Map(mapper func(ps ParserState) interface{}) Parser {
	return Parser{
		Func: func(ps ParserState) ParserState {
			res := p.Func(ps)
			if res.err != nil {
				return res
			}

			res.result = mapper(res)
			return res
		},
	}
}

// Run the parser with `input`
func (p Parser) Run(input string) ParserState {
	state := ParserState{
		input: input,
		idx:   0,
	}

	return p.Func(state)
}

// MapConstant replaces the result with `val` when the parser succeeds
func (p Parser) MapConstant(val interface{}) Parser {
	return Parser{
		Func: func(ps ParserState) ParserState {
			res := p.Func(ps)
			if res.err != nil {
				return res
			}

			res.result = val
			return res
		},
	}
}
