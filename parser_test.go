package pegchamp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	t.Run("one map", func(t *testing.T) {
		parser := String("hello").Map(func(ps ParserState) (interface{}, error) {
			return []interface{}{ps.result}, nil
		})
		expected := []interface{}{"hello"}
		actual := parser.Run("hello world")
		assert.Equal(t, expected, actual.Result())
		assert.Nil(t, actual.Error())
	})

	t.Run("two map", func(t *testing.T) {
		parser := String("hello").Map(func(ps ParserState) (interface{}, error) {
			return []interface{}{ps.result}, nil
		}).Map(func(ps ParserState) (interface{}, error) {
			return ps.result.([]interface{})[0], nil
		})
		expected := "hello"
		actual := parser.Run("hello world")
		assert.Equal(t, expected, actual.Result())
		assert.Nil(t, actual.Error())
	})

	t.Run("mapping error result", func(t *testing.T) {
		parser := String("hello").Map(func(ps ParserState) (interface{}, error) {
			return []interface{}{ps.result}, nil
		}).Map(func(ps ParserState) (interface{}, error) {
			return ps.result.([]interface{})[0], nil
		})
		actual := parser.Run("world")
		assert.Error(t, actual.Error())
	})
}

func TestMapConstant(t *testing.T) {
	t.Run("one map", func(t *testing.T) {
		parser := String("true").MapConstant(true)
		expected := true
		actual := parser.Run("true world")
		assert.Equal(t, expected, actual.Result())
		assert.Nil(t, actual.Error())
	})
}
