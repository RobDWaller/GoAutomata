package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuleToBinary(t *testing.T) {

	type testCase struct {
		rule     uint8
		expected []uint8
	}

	t.Run("valid rule binaries", func(t *testing.T) {

		tests := []testCase{
			{rule: 30, expected: []uint8{0, 0, 0, 1, 1, 1, 1, 0}},
			{rule: 45, expected: []uint8{0, 0, 1, 0, 1, 1, 0, 1}},
			{rule: 5, expected: []uint8{0, 0, 0, 0, 0, 1, 0, 1}},
			{rule: 101, expected: []uint8{0, 1, 1, 0, 0, 1, 0, 1}},
			{rule: 0, expected: []uint8{0, 0, 0, 0, 0, 0, 0, 0}},
			{rule: 255, expected: []uint8{1, 1, 1, 1, 1, 1, 1, 1}},
		}

		for _, test := range tests {
			result := RuleToBinary(test.rule)
			assert.Equal(t, test.expected, result)
		}
	})
}
