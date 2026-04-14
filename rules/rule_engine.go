package rules

// MakeRuleEngine builds a lookup table for a Wolfram rule number.
// Keys are 3-cell neighborhoods (for example, "101") and values are next-state bits (0 or 1).
func MakeRuleEngine(rule uint8) (map[string]uint8, error) {
	binary := RuleToBinary(rule)
	ruleset := GetRuleset()

	engine := make(map[string]uint8)

	for i, value := range ruleset {
		engine[value] = binary[i]
	}

	return engine, nil
}
