package automata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSteps(t *testing.T) {
	tests := []struct {
		name     string
		steps    uint
		rule     uint8
		seed     []uint8
		expected map[uint][]uint8
	}{
		{
			name:     "rule 30, 0 steps returns just seed",
			steps:    0,
			rule:     30,
			seed:     []uint8{0, 1, 0, 1, 1},
			expected: map[uint][]uint8{0: {0, 1, 0, 1, 1}},
		},
		{
			name:  "rule 30, 1 step on five element row",
			steps: 1,
			rule:  30,
			seed:  []uint8{0, 1, 0, 1, 1},
			expected: map[uint][]uint8{
				0: {0, 1, 0, 1, 1},
				1: {0, 1, 0, 1, 0},
			},
		},
		{
			name:  "rule 30, 2 steps on three element row",
			steps: 2,
			rule:  30,
			seed:  []uint8{1, 0, 0},
			expected: map[uint][]uint8{
				0: {1, 0, 0},
				1: {1, 1, 1},
				2: {0, 0, 0},
			},
		},
		{
			name:  "rule 110, 1 step on three element row",
			steps: 1,
			rule:  110,
			seed:  []uint8{1, 0, 0},
			expected: map[uint][]uint8{
				0: {1, 0, 0},
				1: {1, 0, 1},
			},
		},
		{
			name:     "rule 30, single element, 0 steps",
			steps:    0,
			rule:     30,
			seed:     []uint8{1},
			expected: map[uint][]uint8{0: {1}},
		},
		{
			name:  "rule 30, single element, 1 step wraps to 0",
			steps: 1,
			rule:  30,
			seed:  []uint8{1},
			expected: map[uint][]uint8{
				0: {1},
				1: {0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Steps(tt.steps, tt.rule, tt.seed)
			assert.Equal(t, tt.expected, result)
		})
	}
}
