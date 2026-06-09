package automata

func Step(engines []RuleEngine, row []uint8) []uint8 {
	result := make([]uint8, len(row))

	for i := range row {
		n := GetNeighbourhood(row, uint8(i))
		engine, _ := FindNeighbourhood(engines, n)
		result[i] = engine.NextStep
	}

	return result
}

func Steps(steps uint, rule uint8, seed []uint8) map[uint][]uint8 {
	engines := MakeRuleEngine(rule)
	result := make(map[uint][]uint8)

	result[0] = seed
	var i uint
	for i = 0; i < steps; i++ {
		result[i+1] = Step(engines, result[i])
	}

	return result
}
