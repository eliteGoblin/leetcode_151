package dfs

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	prevAddr := make([]string, 0)
	findIp(s, prevAddr, &res)
	return res
}

func findIp(s string, prev []string, res *[]string) {
	if len(prev) == 4 {
		if "" == s {
			(*res) = append(*res, strings.Join(prev, "."))
		}
		return
	}
	for i := 1; i <= len(s); i++ {
		if v, _ := strconv.ParseInt(s[:i], 10, 64); v <= 255 {
			if i > 1 && s[:1] == "0" {
				break
			}
			findIp(s[i:], append(prev, s[:i]), res)
		} else {
			break
		}
	}
}
