package stringproblem

var mp13 = map[string]int{
	"I":  1,
	"V":  5,
	"X":  10,
	"L":  50,
	"C":  100,
	"D":  500,
	"M":  1000,
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}

func romanToInt(s string) int {
	if s == "" {
		return 0
	}
	if len(s) == 1 {
		return mp13[s]
	}
	if v, ok := mp13[s[:2]]; ok {
		return v + romanToInt(s[2:])
	}
	return mp13[s[0:1]] + romanToInt(s[1:])
}
