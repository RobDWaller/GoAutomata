package rules

import (
	"strconv"
	"strings"
)

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
