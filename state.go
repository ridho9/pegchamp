package pegchamp

type ParserState struct {
	input      string
	matchedLen int
	result     interface{}
	err        error
}

func (p ParserState) Result() interface{} {
	return p.result
}

func (p ParserState) Error() error {
	return p.err
}

type Parser func(ParserState) ParserState

func (p Parser) Run(input string) ParserState {
	state := ParserState{
		input: input,
	}

	return p(state)
}
