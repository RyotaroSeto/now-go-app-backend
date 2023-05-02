package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name      string
		password  string
		assertion assert.ErrorAssertionFunc
	}{
		{
			name:      "OK",
			password:  RandomString(6),
			assertion: assert.NoError,
		},
		{
			name:      "NG",
			password:  RandomString(1000),
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := HashPassword(tt.password)
			tt.assertion(t, err)
		})
	}
}
