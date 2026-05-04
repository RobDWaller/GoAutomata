package automata

import (
	"go_automata/rules"

	"testing"

	"github.com/stretchr/testify/assert"
)

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
