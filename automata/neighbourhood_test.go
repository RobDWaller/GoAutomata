package automata

import (
	"go_automata/rules"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNeighbourhood(t *testing.T) {
	t.Run("get neighbourhood", func(t *testing.T) {
		tests := []struct {
			name     string
			row      []uint8
			position uint8
			expected [3]uint8
		}{
			{
				name:     "middle position",
				row:      []uint8{0, 1, 0, 1, 1},
				position: 2,
				expected: [3]uint8{1, 0, 1},
			},
			{
				name:     "start of row wraps previous to end",
				row:      []uint8{0, 1, 0, 1, 1},
				position: 0,
				expected: [3]uint8{1, 0, 1},
			},
			{
				name:     "end of row wraps next to start",
				row:      []uint8{0, 1, 0, 1, 1},
				position: 4,
				expected: [3]uint8{1, 1, 0},
			},
			{
				name:     "single element row wraps both prev and next to itself",
				row:      []uint8{1},
				position: 0,
				expected: [3]uint8{1, 1, 1},
			},
			{
				name:     "two element row position 0",
				row:      []uint8{0, 1},
				position: 0,
				expected: [3]uint8{1, 0, 1},
			},
			{
				name:     "two element row position 1",
				row:      []uint8{0, 1},
				position: 1,
				expected: [3]uint8{0, 1, 0},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result := GetNeighbourhood(test.row, test.position)
				assert.Equal(t, test.expected, result)
			})
		}
	})
}

func TestFindNeighbourhood(t *testing.T) {

	type testCase struct {
		rule          uint8
		neighbourhood [3]uint8
		expected      rules.RuleEngine
		found         bool
	}

	t.Run("find neighbourhood", func(t *testing.T) {

		tests := []testCase{
			{
				rule:          30,
				neighbourhood: [3]uint8{1, 1, 1},
				expected:      rules.RuleEngine{Neighbourhood: [3]uint8{1, 1, 1}, NextStep: 0},
				found:         true,
			},
			{
				rule:          45,
				neighbourhood: [3]uint8{0, 1, 0},
				expected:      rules.RuleEngine{Neighbourhood: [3]uint8{0, 1, 0}, NextStep: 1},
				found:         true,
			},
			{
				rule:          45,
				neighbourhood: [3]uint8{2, 1, 0},
				expected:      rules.RuleEngine{},
				found:         false,
			},
		}

		for _, test := range tests {
			result, found := FindNeighbourhood(test.rule, test.neighbourhood)
			assert.Equal(t, test.expected, result)
			assert.Equal(t, test.found, found)
		}
	})
}
