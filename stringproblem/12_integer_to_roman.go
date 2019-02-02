package stringproblem

var arr = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var mp = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

func intToRoman(num int) string {
	res := ""
	for num > 0 {
		var n int
		for _, v := range arr {
			if num >= v {
				n = v
				break
			}
		}
		res += mp[n]
		num -= n
	}
	return res
}
