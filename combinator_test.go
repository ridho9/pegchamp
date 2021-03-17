package pegchamp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		assert.Error(t, actual.Error())
	})
}

func TestChoice(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		p := Choice(
			String("hello"),
			String("world"),
		)

		expected := "world"
		actual := p.Run("world hello")
		assert.Equal(t, expected, actual.Result())
		assert.Nil(t, actual.Error())
	})

	t.Run("case 2", func(t *testing.T) {
		p := Choice(
			String("hello"),
			String("world"),
		)

		expected := "hello"
		actual := p.Run("hello world")
		assert.Equal(t, expected, actual.Result())
		assert.Nil(t, actual.Error())
	})

	t.Run("case 3", func(t *testing.T) {
		p := Choice(
			String("hello"),
			String("world"),
		)

		actual := p.Run("no dont")
		assert.Error(t, actual.Error())
	})
}
