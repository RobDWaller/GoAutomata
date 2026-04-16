package rules

type RuleEngine struct {
	neigbourhood [3]uint8
	next_step    uint8
}

// MakeRuleEngine builds a lookup table for a Wolfram rule number.
// Keys are 3-cell neighborhoods (for example, "101") and values are next-state bits (0 or 1).
func MakeRuleEngine(rule uint8) []RuleEngine {
	binary := RuleToBinary(rule)
	ruleset := GetRuleset()

	engine := make([]RuleEngine, 8)

	for i, neighbourhood := range ruleset {
		engine[i] = RuleEngine{neigbourhood: neighbourhood, next_step: binary[i]}
	}

	return engine
}
