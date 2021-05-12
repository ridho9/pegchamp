package pegchamp

// ParserPointer takes a pointer to a Parser, which returns a lazy Parser
// which only dereference the pointer in runtime.
// Makes recursive parser and breaks initialization cycle.
func ParserPointer(p *Parser) Parser {
	return Parser{
		parserFunc: func(ps ParserState) ParserState {
			return (*p).Parse(ps)
		},
	}
}

func IgnoreLeadingWhitespace(p Parser) Parser {
	return TakeSecond(
		OptionalWhitespaces(),
		p,
	)
}
