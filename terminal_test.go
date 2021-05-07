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
