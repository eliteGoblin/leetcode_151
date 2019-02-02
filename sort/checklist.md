

## 21. Merge Two Sorted Lists

*  WA 两次

## 23. Merge k Sorted Lists

下面死循环看不出
```golang
for  node, end := nextNode(lists); !end; {
    preNode.Next = node
    preNode = preNode.Next
}
```
*  耗时: naive 350-400 ms; heap: 15ms
*  heap解法
*  分治法

## 147. Insertion Sort List

*  简单但是WA 3次，不验算

## 148. Sort List

*  非常值得review, 自己的代码算作最佳模板, WA卡了1 hour
*  mergeList, 使list都以nil结尾code简单(起始是单节点head.Next=nil), 不要纠结与不用nil结尾，仅用len排序，循环终止条件和最终node置为nil都比较麻烦

## 41. First Missing Positive

*  毫无思路: 原地换，bucket sort

[首个缺失的正数](http://www.cnblogs.com/grandyang/p/4395963.html)

## 75. Sort Colors

*  最简单思路: count sort 没想到
*  partition: 分别用0, 1, 2 作为pivot分割
*  双指针, 除了count sort,其他都是partition变种
*  参考 leetcode-cpp.pdf