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

func TestSequenceOf(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		p := SequenceOf(
			String("hello"),
		)

		expected := []interface{}{"hello"}
		actual := p.Run("hello world").Result()
		assert.Equal(t, expected, actual)
	})

	t.Run("case 2", func(t *testing.T) {
		p := SequenceOf(
			String("hello"),
			String(" "),
			String("world"),
		)

		expected := []interface{}{"hello", " ", "world"}
		actual := p.Run("hello world").Result()
		assert.Equal(t, expected, actual)
	})

	t.Run("case 3", func(t *testing.T) {
		p := SequenceOf(
			String("hello"),
			String("world"),
		)

		actual := p.Run("hello world")
		assert.NotNil(t, actual.Error())
	})
}
