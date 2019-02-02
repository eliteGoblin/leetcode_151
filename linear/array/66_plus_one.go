package array

func PlusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{}
	}
	carry := 0
	digits[len(digits)-1]++
	res := make([]int, len(digits))
	for j := len(digits) - 1; j >= 0; j-- {
		sum := digits[j] + carry
		res[j] = sum % 10
		carry = sum / 10
	}
	if carry > 0 {
		res = append([]int{carry}, res...)
	}
	return res
}
