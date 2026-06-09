package automata

type RuleEngine struct {
	Neighbourhood [3]uint8
	NextStep      uint8
}

func GetRuleset() [8][3]uint8 {
	return [8][3]uint8{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
		{1, 0, 0},
		{0, 1, 1},
		{0, 1, 0},
		{0, 0, 1},
		{0, 0, 0},
	}
}

func RuleToBinary(rule uint8) []uint8 {
	bits := make([]uint8, 8)

	for i := range 8 {
		bits[i] = (rule >> (7 - i)) & 1
	}

	return bits
}

func MakeRuleEngine(rule uint8) []RuleEngine {
	binary := RuleToBinary(rule)
	ruleset := GetRuleset()

	engine := make([]RuleEngine, 8)

	for i, neighbourhood := range ruleset {
		engine[i] = RuleEngine{Neighbourhood: neighbourhood, NextStep: binary[i]}
	}

	return engine
}
