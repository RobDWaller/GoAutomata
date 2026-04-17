package automata

import "go_automata/rules"

func FindNeighbourhood(rule uint8, neighbourhood [3]uint8) rules.RuleEngine {
	return rules.RuleEngine{
		Neighbourhood: [3]uint8{1, 1, 1},
		NextStep:      1,
	}
}
