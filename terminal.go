package pegchamp

import (
	"fmt"
	"unicode/utf8"
)

// Alpha takes a single alphabetical `/[a-zA-Z]/`` character and return it in a string.
func Alpha() Parser {
	return Parser{
		parserFunc: func(ps ParserState) ParserState {
			if ps.err != nil {
				return ps
			}

			if ps.OutOfBound() {
				ps.err = fmt.Errorf("expected alphabetical but found end of input")
				return ps
			}

			head, size := utf8.DecodeRuneInString(ps.input[ps.idx:])
			if isRuneAlpha(head) {
				ps.result = string(head)
				ps.idx += size
				return ps
			}

			ps.err = fmt.Errorf("expected alphabetical but found \"%.16s\"", ps.input[ps.idx:])
			return ps
		},
	}
}

// Alphas takes more than one alphabetical `/[a-zA-Z]+/` characters and return it in a string.
func Alphas() Parser {
	return Parser{
		parserFunc: func(ps ParserState) ParserState {
			if ps.err != nil {
				return ps
			}

			if ps.OutOfBound() {
				ps.err = fmt.Errorf("expected alphabetical but found end of input")
				return ps
			}

			totalLen := 0
			for _, runeValue := range ps.input[ps.idx:] {
				if !isRuneAlpha(runeValue) {
					break
				}
				totalLen += utf8.RuneLen(runeValue)
			}

			ps.result = ps.input[ps.idx : ps.idx+totalLen]
			ps.idx += totalLen
			if ps.result == "" {
				ps.err = fmt.Errorf("expected alphabetical but found \"%.16s\"", ps.input[ps.idx:])
			}
			return ps
		},
	}
}

// Number takes a single numerical `/[0-9]/`` character and return it in a string.
func Number() Parser {
	return Parser{
		parserFunc: func(ps ParserState) ParserState {
			if ps.err != nil {
				return ps
			}

			if ps.OutOfBound() {
				ps.err = fmt.Errorf("expected numerical but found end of input")
				return ps
			}

			head, size := utf8.DecodeRuneInString(ps.input[ps.idx:])
			if isRuneNumber(head) {
				ps.result = string(head)
				ps.idx += size
				return ps
			}

			ps.err = fmt.Errorf("expected numerical but found \"%.16s\"", ps.input[ps.idx:])
			return ps
		},
	}
}

// Numbers takes more than one numerical `/[0-9]+/` characters and return it in a string.
func Numbers() Parser {
	return Parser{
		parserFunc: func(ps ParserState) ParserState {
			if ps.err != nil {
				return ps
			}

			if ps.OutOfBound() {
				ps.err = fmt.Errorf("expected numerical but found end of input")
				return ps
			}

			totalLen := 0
			for _, runeValue := range ps.input[ps.idx:] {
				if !isRuneNumber(runeValue) {
					break
				}
				totalLen += utf8.RuneLen(runeValue)
			}

			ps.result = ps.input[ps.idx : ps.idx+totalLen]
			ps.idx += totalLen
			if ps.result == "" {
				ps.err = fmt.Errorf("expected numerical but found \"%.16s\"", ps.input[ps.idx:])
			}
			return ps
		},
	}
}

// OptionalWhitespaces takes zero or more whitespaces characters and return it in a string.
func OptionalWhitespaces() Parser {
	return Parser{
		parserFunc: func(ps ParserState) ParserState {
			if ps.err != nil {
				return ps
			}

			totalLen := 0
			for _, runeValue := range ps.input[ps.idx:] {
				if !isRuneWhitespace(runeValue) {
					break
				}
				totalLen += utf8.RuneLen(runeValue)
			}

			ps.result = ps.input[ps.idx : ps.idx+totalLen]
			ps.idx += totalLen
			return ps
		},
	}
}

func isRuneAlpha(r rune) bool {
	insideLower := ('a' <= r) && (r <= 'z')
	insideUpper := ('A' <= r) && (r <= 'Z')
	return insideLower || insideUpper
}

func isRuneNumber(r rune) bool {
	return ('0' <= r) && (r <= '9')
}

func isRuneWhitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}
