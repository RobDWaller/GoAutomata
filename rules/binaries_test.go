package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuleToBinary(t *testing.T) {
	t.Run("valid rule binaries", func(t *testing.T) {
		result := RuleToBinary(30)
		assert.Equal(t, "00011110", result)
	})
}
