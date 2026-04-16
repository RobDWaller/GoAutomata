package rules

// GetRuleset returns all 3-cell neighborhoods in Wolfram order.
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
