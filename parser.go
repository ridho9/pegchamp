package pegchamp

type ParserState struct {
	input  string
	idx    int
	result interface{}
	err    error
}

func (p ParserState) Result() interface{} {
	return p.result
}

func (p ParserState) Error() error {
	return p.err
}

// type Parser func(ParserState) ParserState
type Parser struct {
	f func(ParserState) ParserState
}

// Run the parser with `input`
func (p Parser) Run(input string) ParserState {
	state := ParserState{
		input: input,
		idx:   0,
	}

	return p.f(state)
}
