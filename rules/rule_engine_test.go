package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
					{neigbourhood: [3]uint8{1, 1, 1}, next_step: 0},
					{neigbourhood: [3]uint8{1, 1, 0}, next_step: 0},
					{neigbourhood: [3]uint8{1, 0, 1}, next_step: 0},
					{neigbourhood: [3]uint8{1, 0, 0}, next_step: 1},
					{neigbourhood: [3]uint8{0, 1, 1}, next_step: 1},
					{neigbourhood: [3]uint8{0, 1, 0}, next_step: 1},
					{neigbourhood: [3]uint8{0, 0, 1}, next_step: 1},
					{neigbourhood: [3]uint8{0, 0, 0}, next_step: 0},
				},
			},
			{
				rule: 45,
				expected: []RuleEngine{
					{neigbourhood: [3]uint8{1, 1, 1}, next_step: 0},
					{neigbourhood: [3]uint8{1, 1, 0}, next_step: 0},
					{neigbourhood: [3]uint8{1, 0, 1}, next_step: 1},
					{neigbourhood: [3]uint8{1, 0, 0}, next_step: 0},
					{neigbourhood: [3]uint8{0, 1, 1}, next_step: 1},
					{neigbourhood: [3]uint8{0, 1, 0}, next_step: 1},
					{neigbourhood: [3]uint8{0, 0, 1}, next_step: 0},
					{neigbourhood: [3]uint8{0, 0, 0}, next_step: 1},
				},
			},
		}

		for _, test := range tests {
			result := MakeRuleEngine(test.rule)
			assert.Equal(t, test.expected, result)
		}
	})
}
