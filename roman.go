package roman

var romanSymbols = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func ToNumber(roman string) int {
	var sum = 0
	var length = len(roman)

	for i := 0; i < length; i++ {
		if i != length-1 && romanSymbols[roman[i]] < romanSymbols[roman[i+1]] {
			sum += romanSymbols[roman[i+1]] - romanSymbols[(roman[i])]
			i++
		} else {
			sum += romanSymbols[roman[i]]
		}
	}
	return sum
}