## 125. Valid Palindrome

*  思路需要精炼，啰嗦
*  golang不要考虑unicode

## 28. Implement strStr

*  特殊情况: "" ""
*  "a" "a" 证明boundary出错，要用最简单情况验证boundary
*  数组长度:  长len: [start, tail]; tail - start + 1 == len; 最后valid的needle pos: == len(hayStack) - len(needle); 因此虚幻条件:  i < len(hayStack) - len(needle) + 1


## 8. String to Integer (atoi)

*  用FSM解此类题不易出错
*  重点在如何判断溢出，WA了3次: 只要按32bit的最高位，变为1,就算溢出
    +  特殊情况: 最小INT -2^31, 虽然数字刚好能表示，但也可以认为其溢出，返回 INT_MIN，不用纠结这种情况怎么办

## 67. Add Binary

*  阴沟翻船，WA了5次，给出错误case找不到错误:
    +  先res /= 2 再算carry，carry总是0
    +  append([]byte{1})　没有append字符，golang没报错


## 67. Add Binary

*  阴沟翻船，WA了5次，给出错误case找不到错误:
    +  先res /= 2 再算carry，carry总是0
    +  append([]byte{1})　没有append字符，golang没报错

## 10. Regular Expression Matching

*  大致有思路，但是一直写不对，思路不清晰，边界判断不清楚


## 44. Wildcard Matching

*  根据上一题，WA了几次写出递归版，TLE
*  不会分析递归版时间复杂度

## 14. Longest Common Prefix

*  翻车: 空数组

## 65. Valid Number

*  思路不清晰，FSM的map实现，比代码简洁很多

## 49. Group Anagrams

*  最笨的思路: map[byte]count 统计字母频率，判断两个word是否互为anagram; 

