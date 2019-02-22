
## review 

### 120. Triangle

*  另make一个minSum [][]int, 没必要; 另外造成额外的混乱，赋值时弄混minSum和triangle，没必要

### 53. Maximum Subarray

*  经典题，N多解法
*  自己不用DP，卡了半天，纠结于全是负数的数组和结果是空集的conflict
*  应掌握: naive方法(注意不可算空集)与DP方法

### 132 palinfrome partition II 

*  无思路，pdf实现，最简洁

### 85. Maximal Rectangle

*  无思路
*  grandyang: 转化法，转化称为求直方图最大矩形: [Largest Rectangle in Histogram]\(https://leetcode.com/problems/largest-rectangle-in-histogram/), 本身也是hard题目
    +  WA到麻木: 没注意 '0', 以为 0; 直接每层累加得height, 没想到 [0 1] [1 0]情况
*  dp法: [Maximal Rectangle](https://segmentfault.com/a/1190000004239097)


### 123. Best Time to Buy and Sell Stock III

*  直观的算法o(n2)也能AC
*  O(n)算法需要掌握

## 97. Interleaving String

*  grandyang讲解到位:  http://www.cnblogs.com/grandyang/p/4298664.html, code实现参考自己的
*  DP画出图来，清晰很多，尤其是二维数组如何更新
*  Memo法: 用hash来存储不可达解也是一种思路

## 87 Scramble String

*  Recursive Method也WA多次 250 / 283 test cases passed
*  leet_code.pdf分析不错:
*  递归 == 深搜, 加速策略
    +  剪枝
    +  Memo(能memo也能DP)


## 64. Minimum Path Sum

*  Memo WA了多次
*  DP runtime error，下列代码风格容易出错

```golang
var i, j int
	for i = 1; i < len(grid); i++ {
		for j = 1; j < len(grid[0]); j++ {
			grid[i][j] = min(grid[i-1][j], grid[i][j-1])
		}
	}
return grid[i-1][j-1]
// 当[[0]] 时， i == 1, j == 0, j-1 overflow
```

## 72. Edit Distance

*  DP 练习，无思路，解法很经典, grandyang: http://www.cnblogs.com/grandyang/p/4344107.html

## 91. Decode Ways

*  用Memo　WA 8次，　醉了
*  DP WA了4次，看答案改对的，一定Review

## 139. Word Break

*  Memo法　WA了一次，用小case, TDD

## 遗留

* 5
* 96. Unique Binary Search Trees
* 62. Unique Paths
* 63. Unique Paths ii
* 55. Jump Game

