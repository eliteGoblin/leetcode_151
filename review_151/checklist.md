
## 33. Search in Rotated Sorted Array

*  迷糊，之前没真正弄懂
*  不能传入slice, 需要返回之前的index
*  mid := first + ( last - first) / 2

## 4. Median of Two Sorted Arrays

*  不要恐惧，比想象中简单
*  关键问题: 
    +  都有 k / 2时，为什么扔 mid 小的array的左半边
    +  一个没有 k / 2时，为什么扔长数组的左半边，为什么不是扔右半边？
    +  s[k / 2 ] == l[k / 2] 时，如何处理? 会把最终值扔掉吗? 


## 18 4sum

*  小WA,自己代码为最优

## 31. Next Permutation

*  没注意允许重复，小问题导致WA

## 48. Rotate Image

*  花了比较久一次AC，注意:
    +  对角线旋转的规律
    +  invert注意范围，全部invert == 不invert，刚好invert一半

## 135 candy

*  没真正理解解题思路，其实挺简单

## 137 Single Number II

*  依旧无思路
*  符号问题，需要搞清楚int包含的bits数为 32或者64，否则期望负数，给出了大整数

golang int OJ 上64bit,不同于playground上的32bit, int(unsafe.Sizeof(var)) * 8 来求

错误，花了好多时间找到: 

```golang
for _, v := range nums {
    if (v & (1<<uint(i))) > 0 {
        countWrong++
    }
}
```

golang并不会自动转换类型， 


## 2. Add Two Numbers

*  WA 多次，醉了状态不好还是真的菜?
*  看清题目: non-empty, 考虑特殊情况　0 + 0
*  follow up: 不是reverse order

这问题看不出来
```golang
value = value % 10
carry = value / 10
```

## 25 Reverse Nodes in k-Group

*  递归： WA
    +  向后移动到子链表末尾，
    +  抽象问题: 链表分两段: k 和　n - k

*  复杂情况分离子函数

## 146. LRU Cache

*  golang double link list 使用
*  自己搞混了 key, value; 很多WA; list不能只存value, 必须存key, 才能实现O(1)删除key
*  interface 指向struct, 不能直接
*  WA 了 20次，一定回顾; 小聪明想复用节点，只改了value,并没有改key; 修补程序时，之前没想到struct还需要存储key,越改越乱。

```
e.(pair).Value = 2 //编译失败
```

## 98. Validate Binary Search Tree

变量类型转换
```
UINT_MAX = uint(^0) // overflow, -1 --> uint
UINT_MAX = ^uint(0) // pass, 必须是uint类型，否则会认为是-1
INT_MAX = int(UINT_MAX >> 1)
```
* 注意题意：仅有不等时合法

## Tree Traversal

*  stack 自己实现: 熟悉，用这个模板: stack/stack.go
掌握用stack遍历
*  Preorder: 144
*  Inorder: 94
*  Postorder: 145

## 99. Recover Binary Search Tree

*  掌握利用递归的解法，不管morris
*  双指针: 当逆序的两节点相邻时，与不相邻时，逻辑不一样，这里WA。答案参见自己代码

[Recover Binary Search Tree 复原二叉搜索树](http://www.cnblogs.com/grandyang/p/4298069.html)

## 107. Binary Tree Level Order Traversal II

*  双缓存，不必每次考虑边界
*  递归解法思路不错，见第一次提交代码, 递归解法: 传入数组pointer后，不能新建变量: res := *pres, 这样改的还是拷贝，并没有改origin数组，会导致overflow

## 117. Populating Next Right Pointers in Each Node II

*  题目不允许space O(n)的解法
*  还是不会，见自己的解答; 注意WA, root -> right 漏了 -> next

## 96 Unique Binary Search Trees

*  不难题，dp简单练习
*  简单代码问题上一直WA: n - j 应该为 i - j, 验算没发现，一直打log才发现

## 95. Unique Binary Search Trees II


*  上题的基础上输出结果树
*  自己的解答参考简化思路，巧妙的点:
    +  只用下标，不用传递数组
    +  []*TreeNode{nil}的妙用，即使是空，也会执行至少一次，而不会发生自己担心的跳过循环的情况

## 98. Validate Binary Search Tree

*  注意相等的情况，定义不包含相等，意味着相等的情况不满足题意

## TODO

*  Morris Inorder
*  Convert Sorted List to Binary Search Tree


## General

*  -- 写成 ++
*  doc写的 type mismatch, 没有取数组下标
*  取数组下标赋值时切记判断是否overflow, 比如只想初始化数组，res[0] = xx res[1] = xx 确保数组len
*  for i < j{} 总忘记写i++,j--