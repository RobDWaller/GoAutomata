package rules

// RuleToBinary converts a Wolfram rule number to eight binary digits (MSB to LSB).
// e.g. Rule 30 => []uint8{0, 0, 0, 1, 1, 1, 1, 0}
// Code uses bitwise operators to return correct binary bit.
func RuleToBinary(rule uint8) []uint8 {
	bits := make([]uint8, 8)

	for i := range 8 {
		bits[i] = (rule >> (7 - i)) & 1
	}

	return bits
}
