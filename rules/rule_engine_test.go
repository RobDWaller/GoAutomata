package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeRuleEngine(t *testing.T) {

	type testCase struct {
		rule     uint8
		expected map[string]uint8
	}

	t.Run("valid rule engines", func(t *testing.T) {

		tests := []testCase{
			{rule: 30, expected: map[string]uint8{"111": 0, "110": 0, "101": 0, "100": 1, "011": 1, "010": 1, "001": 1, "000": 0}},
			{rule: 45, expected: map[string]uint8{"111": 0, "110": 0, "101": 1, "100": 0, "011": 1, "010": 1, "001": 0, "000": 1}},
		}

		for _, test := range tests {
			result, err := MakeRuleEngine(test.rule)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	})
}
