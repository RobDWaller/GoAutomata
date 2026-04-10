package rules

import (
	"strconv"
	"strings"
)

func MakeRuleEngine(rule uint8) (map[string]uint8, error) {
	binary := strings.Split(RuleToBinary(rule), "")
	ruleset := GetRuleset()

	engine := make(map[string]uint8)

	for i, rule := range ruleset {
		result, err := strconv.ParseUint(binary[i], 2, 8)

		engine[rule] = uint8(result)

		if err != nil {
			return engine, err
		}
	}

	return engine, nil
}
