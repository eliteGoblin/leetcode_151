package array

func Candy(ratings []int) int {
	candyArr := make([]int, len(ratings))
	candyArr[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candyArr[i] = candyArr[i-1] + 1
		} else {
			candyArr[i] = 1
		}
	}
	res := candyArr[len(ratings)-1]
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candyArr[i] = max(candyArr[i], candyArr[i+1]+1)
		}
		res += candyArr[i]
	}
	return res
}
