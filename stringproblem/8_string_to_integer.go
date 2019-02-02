package stringproblem

const (
	start = iota
	sign
)

const (
	INT_MAX = 1<<31 - 1
	INT_MIN = 1 << 31
)

func MyAtoi(str string) int {
	res := 0
	curPos := 0
	for ; curPos < len(str); curPos++ {
		if str[curPos] != ' ' {
			break
		}
	}
	if curPos == len(str) {
		return res
	}
	state := start
	isPositive := true
	for {
		switch state {
		case start:
			if str[curPos] >= '0' && str[curPos] <= '9' {
				res = res*10 + int(str[curPos]-'0')
				state = sign
			} else if str[curPos] == '+' || str[curPos] == '-' {
				if str[curPos] == '-' {
					isPositive = false
				}
				state = sign
			} else {
				return resWithSign(res, isPositive)
			}
			curPos++
		case sign:
			for curPos < len(str) {
				if str[curPos] >= '0' && str[curPos] <= '9' {
					res = res*10 + int(str[curPos]-'0')
					if res&(^0x7fffffff) != 0 {
						if isPositive {
							return INT_MAX
						}
						return INT_MIN
					}
				} else {
					return resWithSign(res, isPositive)
				}
				curPos++
			}
			return resWithSign(res, isPositive)
		default:
			return resWithSign(res, isPositive)
		}
	}
}

func resWithSign(res int, isPositive bool) int {
	if isPositive {
		return res
	}
	return -1 * res
}
