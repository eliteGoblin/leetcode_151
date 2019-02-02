package stringproblem

func AddBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	if len(b) == 0 {
		return a
	}
	var carry byte
	ia, ib := len(a)-1, len(b)-1
	bytesRes := make([]byte, len(a))
	for ia >= 0 {
		var curB byte
		if ib >= 0 {
			curB = b[ib] - '0'
		} else {
			curB = 0
		}
		res := carry + a[ia] - '0' + curB
		carry = res / 2
		res = res % 2
		bytesRes[ia] = res + '0'
		ia--
		ib--
	}
	if carry > 0 {
		bytesRes = append([]byte{1}, bytesRes...)
	}
	return string(bytesRes)
}
