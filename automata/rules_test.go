package automata

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

func TestMakeRuleEngine(t *testing.T) {
	type testCase struct {
		rule     uint8
		expected []RuleEngine
	}

	t.Run("valid rule engines", func(t *testing.T) {
		tests := []testCase{
			{
				rule: 30,
				expected: []RuleEngine{
					{Neighbourhood: [3]uint8{1, 1, 1}, NextStep: 0},
					{Neighbourhood: [3]uint8{1, 1, 0}, NextStep: 0},
					{Neighbourhood: [3]uint8{1, 0, 1}, NextStep: 0},
					{Neighbourhood: [3]uint8{1, 0, 0}, NextStep: 1},
					{Neighbourhood: [3]uint8{0, 1, 1}, NextStep: 1},
					{Neighbourhood: [3]uint8{0, 1, 0}, NextStep: 1},
					{Neighbourhood: [3]uint8{0, 0, 1}, NextStep: 1},
					{Neighbourhood: [3]uint8{0, 0, 0}, NextStep: 0},
				},
			},
			{
				rule: 45,
				expected: []RuleEngine{
					{Neighbourhood: [3]uint8{1, 1, 1}, NextStep: 0},
					{Neighbourhood: [3]uint8{1, 1, 0}, NextStep: 0},
					{Neighbourhood: [3]uint8{1, 0, 1}, NextStep: 1},
					{Neighbourhood: [3]uint8{1, 0, 0}, NextStep: 0},
					{Neighbourhood: [3]uint8{0, 1, 1}, NextStep: 1},
					{Neighbourhood: [3]uint8{0, 1, 0}, NextStep: 1},
					{Neighbourhood: [3]uint8{0, 0, 1}, NextStep: 0},
					{Neighbourhood: [3]uint8{0, 0, 0}, NextStep: 1},
				},
			},
		}

		for _, test := range tests {
			result := MakeRuleEngine(test.rule)
			assert.Equal(t, test.expected, result)
		}
	})
}
