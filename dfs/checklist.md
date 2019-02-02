## 131. Palindrome Partitioning

*  没思路: [Palindrome Partitioning 拆分回文串](http://www.cnblogs.com/grandyang/p/4270008.html)

## 62. Unique Paths

*  深搜超时
*  备忘录: 注意自己代码两点，自己第一次写的代码一直runtime, 1 index, 代码问题导致数组溢出
*  动归，数学公式

## 51 N Queens

* 经典深搜
* 没思路, 看答案还WA

## 93. Restore IP Addresses

*  思路有，漏掉 前导0需要去掉的case: 0.010.xx.xx 第二段不合法

## 39. Combination Sum

* WA 重复结果
* 40 一次AC，但是review好题

## 22. Generate Parentheses

*  小错误

## 37. Sudoku Solver

* 自己WA了几次做出来，但是代码太复杂, 过早优化，是弱点!!!! don't be too clever一开始: 37_xxx_premature.go
* 简洁代码 参见pdf
自己优化的代码:
Runtime: 48 ms, faster than 18.18% of Go online submissions for Sudoku Solver.
Memory Usage: 1 MB, less than 33.33% of Go online submissions for Sudoku Solver.
简洁代码
Runtime: 8 ms, faster than 100.00% of Go online submissions for Sudoku Solver.
Memory Usage: 929.8 KB, less than 100.00% of Go online submissions for Sudoku Solver.
别自己瞎jb"优化", 9*9的矩阵不需要用map cache

## 79. Word Search

* 回溯没啥问题，回溯过程自己会清除visited数组，不需要自己每次clear甚至每次new, 见代码