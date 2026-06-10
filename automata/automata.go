package automata

func Generate(steps uint, rule uint8, seed string) map[uint][]uint8 {
	seedBytes := make([]uint8, len(seed))
	for i, c := range seed {
		seedBytes[i] = uint8(c - '0')
	}
	return Steps(steps, rule, seedBytes)
}
