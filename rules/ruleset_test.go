package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRuleset(t *testing.T) {

	type testCase struct {
		expected [8][3]uint8
	}

	t.Run("valid ruleset", func(t *testing.T) {

		tests := []testCase{
			{expected: [8][3]uint8{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}, {1, 0, 0}, {0, 1, 1}, {0, 1, 0}, {0, 0, 1}, {0, 0, 0}}},
		}

		for _, test := range tests {
			result := GetRuleset()
			assert.Equal(t, test.expected, result)
		}
	})
}
