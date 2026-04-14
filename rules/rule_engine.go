package rules

import (
	"strconv"
	"strings"
)

// MakeRuleEngine builds a lookup table for a Wolfram rule number.
// Keys are 3-cell neighborhoods (for example, "101") and values are next-state bits (0 or 1).
func MakeRuleEngine(rule uint8) (map[string]uint8, error) {
	binary := strings.Split(RuleToBinary(rule), "")
	ruleset := GetRuleset()

	engine := make(map[string]uint8)

	for i, value := range ruleset {
		result, err := strconv.ParseUint(binary[i], 2, 8)

		if err != nil {
			return engine, err
		}

		engine[value] = uint8(result)
	}

	return engine, nil
}
