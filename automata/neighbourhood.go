package automata

import "go_automata/rules"

func GetNeighbourhood(row []uint8, position uint8) [3]uint8 {
	n := len(row)
	i := int(position)
	prev := row[(i+n-1)%n]
	curr := row[i]
	next := row[(i+1)%n]
	return [3]uint8{prev, curr, next}
}

func FindNeighbourhood(engines []rules.RuleEngine, neighbourhood [3]uint8) (rules.RuleEngine, bool) {

	for _, engine := range engines {
		if engine.Neighbourhood == neighbourhood {
			return engine, true
		}
	}

	return rules.RuleEngine{}, false
}
