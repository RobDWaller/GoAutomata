package automata

import "go_automata/rules"

func FindNeighbourhood(rule uint8, neighbourhood [3]uint8) (rules.RuleEngine, bool) {

	for _, engine := range rules.MakeRuleEngine(rule) {
		if engine.Neighbourhood == neighbourhood {
			return engine, true
		}
	}

	return rules.RuleEngine{}, false
}
