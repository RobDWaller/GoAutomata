package rules

import "fmt"

func RuleToBinary(rule uint8) (string, error) {
	return fmt.Sprintf("%08b", rule), nil
}
