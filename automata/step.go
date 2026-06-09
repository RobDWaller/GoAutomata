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
