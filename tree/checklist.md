## Tree Input Visualize Tool

*  98. Validate Binary Search Tree
*  难度高的morris直接放弃

## Morris List

*  144
*  94
*  145
*  99

## 树遍历

*  InOrder: 写错
*  Morris Method

### PostOrder

*  two stack
*  one stack
*  Morris


[Iterative Postorder Traversal](https://www.geeksforgeeks.org/iterative-postorder-traversal/)

## 107. Binary Tree Level Order Traversal II

*  练习递归解法
*  reverse

## 103. Binary Tree Zigzag Level Order Traversal

*  上题变种，思路转换

## 99. Recover Binary Search Tree

*  题目提到的straightforward思路也没有
*  后续集中Morris

## 101. Symmetric Tree

*  思路复杂: 先invert tree, 再判断是否same；见最优思路

## 117. Populating Next Right Pointers in Each Node II

*  只有题目不允许的o(n) sDon't spend more than an hour on any question. If you can't figure out the solution, mark it and revisit later.
*  [Grandyang](http://wwwDon't spend more than an hour on any question. If you can't figure out the solution, mark it and revisit later./4290148.html)

## 105. Construct Binary Don't spend more than an hour on any question. If you can't figure out the solution, mark it and revisit later.norder Traversal

*  loop variable 忘记++导Don't spend more than an hour on any question. If you can't figure out the solution, mark it and revisit later.
*  经典题，值得回顾，下标Don't spend more than an hour on any question. If you can't figure out the solution, mark it and revisit later.

## 96. Unique Binary Search Trees

*  递归写法出错，没有带入最简单情况检查，习惯要养成
*  动态规划初步

## 95. Unique Binary Search Trees II


*  理解错题，没思路(思路特别复杂)

## 98. Validate Binary Search Tree

*  没思路，其实很简单


## 108. Convert Sorted Array to Binary Search Tree

*  之前　mid = len(nums) / 2 - 1, 没有带入最简单的case

## 109. Convert Sorted List to Binary Search Tree

*  只会强行将list转为array的笨办法(forward list), WA 8 次
    +  没有用最简单case检验，否则WA会少很多
*  其他思路与自己笨办法类似，只是clear
    +  fast, slow一遍定中点
    +  递归传入 head, last，　而非仅head来简化递归
    +  array的中间偏左作为mid(为偶数时):  start + (end - start) / 2
*  自己最佳解法见已提交代码

## 111. Minimum Depth of Binary Tree

*  小细节　翻船
*  review 迭代版

## 112. Path Sum

*  搞错题意

## 113. Path Sum II

*  一直WA, 15次: golang　slice部分表现为引用的问题
  +  正确答案 append到[][]int [5,8,5,4]
  +  错误path: [5,8,5,1]改动同一片data指向的slice，导致结果改变

## 124. Binary Tree Maximum Path Sum

*  无思路
*  [求二叉树的最大路径和](http://www.cnblogs.com/grandyang/p/4280120.html)