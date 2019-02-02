## 78. Subsets

*  排序输入AC也是醉了，不明白依据
*  迭代法，也挺简单

## 90. Subsets II

*  WA了N次，golang特性，边界条件小问题，unit test
```golang
//向range增加元素并不会陷入死循环
arr := []int{1,2,3,4,5}
for i := range arr {
    arr = append(arr, -1)
}
// 死循环
for i := 0; i < len(arr); i ++ {
    arr = append(arr, -1)
}
```

## 47. Permutations II

*  无思路，复习next_permutation: 47. Permutations II
*  递归很复杂，也不想看

## 77. Combinations

*  无思路
*  DFS搜索, 递归公式(注意{{}} 和 {} 的区别)

[Combinations 组合项](http://www.cnblogs.com/grandyang/p/4332522.html)

## 17. Letter Combinations of a Phone Number

*  简单题别翻船，练习匿名函数递归 