package pegchamp

func IgnoreLeadingWhitespace(p Parser) Parser {
	return TakeSecond(
		OptionalWhitespaces(),
		p,
	)
}

// ParserFunc takes a `func() Parser`
// which called in runtime.
// Makes recursive parser and breaks initialization cycle.
func ParserFunc(p func() Parser) Parser {
	return Parser{
		parserFunc: func(ps ParserState) ParserState {
			return p().Parse(ps)
		},
	}
}
