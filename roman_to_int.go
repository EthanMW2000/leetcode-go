package main

func romanToInt(s string) int {
	numeralsAsInts := map[byte]int {
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	numerals := []byte(s)
	total, i := 0, 0
	length := len(s)
	
	for i < length {
		if len(s) - 1 == i {
			total += numeralsAsInts[numerals[i]]
		} else if numerals[i] == 'I' && (numerals[i + 1] == 'V' || numerals[i + 1] == 'X') {
			total += numeralsAsInts[numerals[i + 1]] - numeralsAsInts[numerals[i]]
			i++
		} else if numerals[i] == 'X' && (numerals[i + 1] == 'L' || numerals[i + 1] == 'C') {
			total += numeralsAsInts[numerals[i + 1]] - numeralsAsInts[numerals[i]]
			i++
		} else if numerals[i] == 'C' && (numerals[i + 1] == 'D' || numerals[i + 1] == 'M') {
			total += numeralsAsInts[numerals[i + 1]] - numeralsAsInts[numerals[i]]
			i++
		} else {
			total += numeralsAsInts[numerals[i]]
		}
		i++
	}

	return total
}
