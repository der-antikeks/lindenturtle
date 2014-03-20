package main

// http://en.wikipedia.org/wiki/L-system
// iterative function
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

// concurrentlindenmayer -> conlindenmayer -> conlinmayer -> coliner!
// channel iterations
func coliner(start []string, rules map[string][]string, iterations int) []string {
	replace := func(in chan string) chan string {
		out := make(chan string)
		go func() {
			defer close(out)
			for s := range in {
				r, ok := rules[s]
				if !ok {
					out <- s
					continue
				}
				for _, c := range r {
					out <- c
				}
			}
		}()
		return out
	}

	c := make(chan string)
	go func(in chan string) {
		for _, s := range start {
			in <- s
		}
		close(in)
	}(c)

	for i := 0; i < iterations; i++ {
		c = replace(c)
	}

	result := []string{}
	for s := range c {
		result = append(result, s)
	}

	return result
}
