package rules

// GetRuleset returns all 3-cell neighborhoods in Wolfram order.
func GetRuleset() [8]string {
	return [8]string{"111", "110", "101", "100", "011", "010", "001", "000"}
}
