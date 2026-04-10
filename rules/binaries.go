package rules

import "fmt"

func RuleToBinary(rule uint8) string {
	return fmt.Sprintf("%08b", rule)
}
