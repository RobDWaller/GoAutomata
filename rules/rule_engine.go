package rules

type RuleEngine struct {
	Neighbourhood [3]uint8
	NextStep      uint8
}

// MakeRuleEngine builds a lookup table for a Wolfram rule number.
// Keys are 3-cell neighborhoods (for example, "101") and values are next-state bits (0 or 1).
func MakeRuleEngine(rule uint8) []RuleEngine {
	binary := RuleToBinary(rule)
	ruleset := GetRuleset()

	engine := make([]RuleEngine, 8)

	for i, neighbourhood := range ruleset {
		engine[i] = RuleEngine{Neighbourhood: neighbourhood, NextStep: binary[i]}
	}

	return engine
}
