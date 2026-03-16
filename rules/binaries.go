package rules

import "fmt"

func RuleToBinary(rule int16) string {
	return fmt.Sprintf("%08b", rule)
}
