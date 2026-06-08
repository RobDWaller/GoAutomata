package automata

import "go_automata/rules"

func Steps(steps uint, rule uint8, seed []uint8) map[uint][]uint8 {
	engines := rules.MakeRuleEngine(rule)
	result := make(map[uint][]uint8)

	result[0] = seed
	var i uint
	for i = 0; i < steps; i++ {
		result[i+1] = Step(engines, result[i])
	}

	return result
}
