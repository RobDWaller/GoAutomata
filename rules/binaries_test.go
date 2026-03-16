package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuleToBinary(t *testing.T) {

	type testCase struct {
		rule     uint8
		expected string
	}

	t.Run("valid rule binaries", func(t *testing.T) {

		tests := []testCase{
			{rule: 30, expected: "00011110"},
			{rule: 45, expected: "00101101"},
			{rule: 5, expected: "00000101"},
			{rule: 101, expected: "01100101"},
			{rule: 0, expected: "00000000"},
			{rule: 255, expected: "11111111"},
		}

		for _, test := range tests {
			result, err := RuleToBinary(test.rule)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	})
}
