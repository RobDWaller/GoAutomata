package automata

import "go_automata/rules"

func FindNeighbourhood(rule uint8, neighbourhood [3]uint8) rules.RuleEngine {
	ruleEngines := rules.MakeRuleEngine(rule)

	var ruleEngine rules.RuleEngine

	for _, engine := range ruleEngines {
		if engine.Neighbourhood == neighbourhood {
			ruleEngine = engine
			break
		}
	}

	return ruleEngine
}
