package pegchamp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSequenceOf(t *testing.T) {
	tests := []struct {
		name        string
		parsers     []Parser
		input       string
		expected    []interface{}
		shouldError bool
	}{
		{
			name:     "single string",
			parsers:  []Parser{String("hello")},
			input:    "hello world",
			expected: []interface{}{"hello"},
		},
		{
			name:     "multiple string",
			parsers:  []Parser{String("hello"), String(" "), String("world")},
			input:    "hello world",
			expected: []interface{}{"hello", " ", "world"},
		},
		{
			name:        "error",
			parsers:     []Parser{String("hello"), String("world")},
			input:       "hello world",
			shouldError: true,
		},
		{
			name:     "char and string",
			parsers:  []Parser{Char('h'), String("ello")},
			input:    "hello world",
			expected: []interface{}{byte('h'), "ello"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := SequenceOf(tt.parsers...).Run(tt.input)

			if !tt.shouldError {
				assert.EqualValues(t, tt.expected, actual.Result())
				assert.Nil(t, actual.Error())
			} else {
				assert.Error(t, actual.Error())
			}
		})
	}
}

func TestChoice(t *testing.T) {
	tests := []struct {
		name        string
		parsers     []Parser
		input       string
		expected    interface{}
		shouldError bool
	}{
		{
			name:     "user first parser",
			parsers:  []Parser{String("hello"), String("world")},
			input:    "hello",
			expected: "hello",
		},
		{
			name:     "user second parser",
			parsers:  []Parser{String("hello"), String("world")},
			input:    "world",
			expected: "world",
		},
		{
			name:        "error",
			parsers:     []Parser{String("a"), String("b")},
			input:       "dont",
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Choice(tt.parsers...).Run(tt.input)

			if !tt.shouldError {
				assert.Equal(t, tt.expected, actual.Result())
				assert.Nil(t, actual.Error())
			} else {
				assert.Error(t, actual.Error())
			}
		})
	}
}

func TestTakeSecond(t *testing.T) {
	tests := []struct {
		name        string
		first       Parser
		second      Parser
		input       string
		expected    interface{}
		shouldError bool
	}{
		{
			name:     "correct",
			first:    String("h"),
			second:   String("ello"),
			input:    "hello",
			expected: "ello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := TakeSecond(tt.first, tt.second).Run(tt.input)

			if !tt.shouldError {
				assert.Equal(t, tt.expected, actual.Result())
				assert.Nil(t, actual.Error())
			} else {
				assert.Error(t, actual.Error())
			}
		})
	}
}
