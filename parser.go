package pegchamp

type ParserState struct {
	input  string
	idx    int
	result interface{}
	err    error
}

func (ps ParserState) OutOfBound() bool {
	return ps.idx >= len(ps.input)
}

func (p ParserState) Head() string {
	return p.input[p.idx:]
}

func (p ParserState) Index() int {
	return p.idx
}

func (p ParserState) AdvanceIndex(val int) ParserState {
	p.idx += val
	return p
}

func (p ParserState) InputLen() int {
	return len(p.input)
}

func (p ParserState) Result() interface{} {
	return p.result
}

func (p ParserState) SetResult(res interface{}) ParserState {
	p.result = res
	return p
}

func (p ParserState) Error() error {
	return p.err
}

func (p ParserState) SetError(val error) ParserState {
	p.err = val
	return p
}

type Parser struct {
	parserFunc func(ParserState) ParserState
}

func NewParser(f func(ParserState) ParserState) Parser {
	return Parser{
		parserFunc: f,
	}
}

// Run the parser with `input`
func (p Parser) Run(input string) ParserState {
	state := ParserState{
		input: input,
		idx:   0,
	}

	return p.Parse(state)
}

func (p Parser) Parse(ps ParserState) ParserState {
	if ps.err != nil {
		return ps
	}
	return p.parserFunc(ps)
}

// Map takes a mapper function that takes the parser result and return new result.
// Other part of the ParserState won't be changed.
// The mapper function won't be run in the case of a parsing error
// so it could be expected only a successful result is passed.
// The mapper function could also return an error which get assigned into the returned
// parser state.
func (p Parser) Map(mapper func(ps ParserState) (interface{}, error)) Parser {
	return Parser{
		parserFunc: func(ps ParserState) ParserState {
			res := p.Parse(ps)
			if res.err != nil {
				return res
			}

			mapRes, err := mapper(res)
			res.err = err
			res.result = mapRes
			return res
		},
	}
}

// MapConstant replaces the result with `val` when the parser succeeds
func (p Parser) MapConstant(val interface{}) Parser {
	return Parser{
		parserFunc: func(ps ParserState) ParserState {
			res := p.Parse(ps)
			if res.err != nil {
				return res
			}

			res.result = val
			return res
		},
	}
}
