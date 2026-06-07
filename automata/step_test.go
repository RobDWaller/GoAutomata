package automata

import (
	"go_automata/rules"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStep(t *testing.T) {
	tests := []struct {
		name     string
		rule     uint8
		row      []uint8
		expected []uint8
	}{
		{
			name:     "rule 30 on five element row",
			rule:     30,
			row:      []uint8{0, 1, 0, 1, 1},
			expected: []uint8{0, 1, 0, 1, 0},
		},
		{
			name:     "rule 30 on three element row",
			rule:     30,
			row:      []uint8{1, 0, 0},
			expected: []uint8{1, 1, 1},
		},
		{
			name:     "rule 110 on three element row",
			rule:     110,
			row:      []uint8{1, 0, 0},
			expected: []uint8{1, 0, 1},
		},
		{
			name:     "rule 30 single element 1 wraps to itself",
			rule:     30,
			row:      []uint8{1},
			expected: []uint8{0},
		},
		{
			name:     "rule 30 single element 0 wraps to itself",
			rule:     30,
			row:      []uint8{0},
			expected: []uint8{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Step(rules.MakeRuleEngine(tt.rule), tt.row)
			assert.Equal(t, tt.expected, result)
		})
	}
}
