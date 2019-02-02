package stringproblem

// const (
// 	start = iota
// 	afterSign
// 	afterDot
// 	afterE
// 	afterESign
// )

// func isNumber(s string) bool {
// 	if "" == s {
// 		return true
// 	}
// 	state := start
// 	i := 0
// 	isExpectDigit := false
// 	for i < len(s) {
// 		switch state {
// 		case start:
// 			if isSign(s[i]) {
// 				state = afterSign
// 				isExpectDigit = true
// 			} else if isDigit(s[i]) {
// 				state = afterSign
// 			} else {
// 				return false
// 			}
// 		case afterSign:
// 			if isExpectDigit {
// 				if !isDigit(s[i]) {
// 					return false
// 				}
// 				isExpectDigit = false
// 				break
// 			}
// 			if '.' == s[i] {
// 				state = afterDot
// 				isExpectDigit = true
// 			} else if 'e' == s[i] {
// 				state = afterE
// 				isExpectDigit = true
// 			} else if !isDigit(s[i]) {
// 				return false
// 			}
// 		case afterDot:
// 			if isExpectDigit {
// 				if !isDigit(s[i]) {
// 					return false
// 				}
// 				isExpectDigit = false
// 				break
// 			}
// 			if 'e' == s[i] {
// 				state = afterE
// 			} else if !isDigit(s[i]) {
// 				return false
// 			}
// 		case afterE:
// 			if isDigit(s[i]) {
// 				state = afterESign
// 				isExpectDigit = false
// 			} else if isSign(s[i]) {
// 				state = afterESign
// 				isExpectDigit = true
// 			} else {
// 				return false
// 			}
// 		case afterESign:
// 		}
// 		i++
// 	}
// 	return ifDotHasDigit
// }

// func isSign(c byte) bool {
// 	return '+' == c || '-' == c
// }
// func isDigit(c byte) bool {
// 	return c >= '0' && c <= '9'
// }
