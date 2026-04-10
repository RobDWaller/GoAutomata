package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRuleset(t *testing.T) {

	type testCase struct {
		expected [1]string
	}

	t.Run("valid rule binaries", func(t *testing.T) {

		tests := []testCase{
			{expected: [1]string{"111"}},
		}

		for _, test := range tests {
			result := GetRuleset()
			assert.Equal(t, test.expected, result)
		}
	})
}
