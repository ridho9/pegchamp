package pegchamp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlpha(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		want        string
		shouldError bool
	}{
		{
			name:  "success",
			input: "hello world",
			want:  "h",
		},
		{
			name:        "fail",
			input:       "1234",
			shouldError: true,
		},
		{
			name:        "empty",
			input:       "",
			shouldError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Alpha().Run(tt.input)

			if !tt.shouldError {
				assert.Equal(t, tt.want, actual.Result())
				assert.Nil(t, actual.Error())
			} else {
				assert.Error(t, actual.Error())
			}
		})
	}
}

func TestAlphas(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		want        string
		shouldError bool
	}{
		{
			name:  "success",
			input: "hello world",
			want:  "hello",
		},
		{
			name:  "success 2",
			input: "hello",
			want:  "hello",
		},
		{
			name:        "minimal one",
			input:       "1",
			shouldError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Alphas().Run(tt.input)

			if !tt.shouldError {
				assert.Equal(t, tt.want, actual.Result())
				assert.Nil(t, actual.Error())
			} else {
				assert.Error(t, actual.Error())
			}
		})
	}
}

func TestNumber(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		want        string
		shouldError bool
	}{
		{
			name:  "success",
			input: "123",
			want:  "1",
		},
		{
			name:        "fail",
			input:       "hello",
			shouldError: true,
		},
		{
			name:        "empty",
			input:       "",
			shouldError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Number().Run(tt.input)

			if !tt.shouldError {
				assert.Equal(t, tt.want, actual.Result())
				assert.Nil(t, actual.Error())
			} else {
				assert.Error(t, actual.Error())
			}
		})
	}
}

func TestNumbers(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		want        string
		shouldError bool
	}{
		{
			name:  "success",
			input: "123 456",
			want:  "123",
		},
		{
			name:  "success 2",
			input: "123",
			want:  "123",
		},
		{
			name:        "minimal one",
			input:       "a",
			shouldError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Numbers().Run(tt.input)

			if !tt.shouldError {
				assert.Equal(t, tt.want, actual.Result())
				assert.Nil(t, actual.Error())
			} else {
				assert.Error(t, actual.Error())
			}
		})
	}
}

func TestOptionalWhitespaces(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		want        string
		shouldError bool
	}{
		{
			name:  "success",
			input: " 456",
			want:  " ",
		},
		{
			name:  "success 2",
			input: "asd",
			want:  "",
		},
		{
			name:  "success 3",
			input: " \t\n\r123",
			want:  " \t\n\r",
		},
		{
			name:  "success 4",
			input: " \t\n\r",
			want:  " \t\n\r",
		},
		{
			name:  "success 5",
			input: "",
			want:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := OptionalWhitespaces().Run(tt.input)

			if !tt.shouldError {
				assert.Equal(t, tt.want, actual.Result())
				assert.Nil(t, actual.Error())
			} else {
				assert.Error(t, actual.Error())
			}
		})
	}
}
