package examples

func RomanToInt(s string) int {

	total := 0
	ant := 4000

	values := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for _, v := range s {
		if values[string(v)] <= ant {
			total += values[string(v)]

		} else {
			total += values[string(v)] - (ant * 2)

		}
		ant = values[string(v)]
	}

	return total
}
