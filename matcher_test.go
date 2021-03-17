package pegchamp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	t.Run("one string", func(t *testing.T) {
		input := "abc"
		str := "a"
		expected := ParserState{
			input:  "bc",
			result: "a",
		}

		actual := String(str).Run(input)
		assert.Equal(t, expected.Result(), actual.Result())
	})

	t.Run("empty string", func(t *testing.T) {
		input := "abc"
		str := ""
		expected := ParserState{
			input:  "abc",
			result: "",
		}

		actual := String(str).Run(input)
		assert.Equal(t, expected.Result(), actual.Result())
	})

	t.Run("invalid input", func(t *testing.T) {
		input := "hello"
		str := "world"

		actual := String(str).Run(input)
		assert.NotNil(t, actual.Error())
	})

	t.Run("empty input", func(t *testing.T) {
		input := ""
		str := "world"

		actual := String(str).Run(input)
		assert.NotNil(t, actual.Error())
	})
}

func TestChar(t *testing.T) {
	tests := []struct {
		name        string
		c           byte
		input       string
		want        byte
		shouldError bool
	}{
		{
			name:  "case 1",
			c:     'h',
			input: "hello world",
			want:  'h',
		},
		{
			name:        "case 2",
			c:           'h',
			input:       "",
			shouldError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Char(tt.c).Run(tt.input)

			if !tt.shouldError {
				assert.Equal(t, tt.want, actual.Result())
				assert.Nil(t, actual.Error())
			} else {
				assert.Error(t, actual.Error())
			}
		})
	}
}
