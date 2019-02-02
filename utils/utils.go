package utils

func Max(elements ...int) int {
	m := elements[0]
	for i := 1; i < len(elements); i++ {
		if elements[i] > m {
			m = elements[i]
		}
	}
	return m
}
