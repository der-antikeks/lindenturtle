package main

// http://en.wikipedia.org/wiki/L-system
func lindenmayer(start []string, rules map[string][]string, iterations int) []string {
	if iterations == 0 {
		return start
	}

	result := []string{}
	for _, c := range start {
		if r, ok := rules[c]; ok {
			result = append(result, r...)
		} else {
			result = append(result, c)
		}
	}

	if iterations--; iterations > 0 {
		return lindenmayer(result, rules, iterations)
	}

	return result
}
