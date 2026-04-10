package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRuleset(t *testing.T) {

	type testCase struct {
		expected [8]string
	}

	t.Run("valid rule binaries", func(t *testing.T) {

		tests := []testCase{
			{expected: [8]string{"111", "110", "101", "100", "011", "010", "001", "000"}},
		}

		for _, test := range tests {
			result := GetRuleset()
			assert.Equal(t, test.expected, result)
		}
	})
}
