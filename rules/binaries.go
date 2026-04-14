package rules

import "fmt"

// RuleToBinary converts a Wolfram rule number to an 8-bit binary string.
// eg: Rule 30 to 00011110
func RuleToBinary(rule uint8) string {
	return fmt.Sprintf("%08b", rule)
}
