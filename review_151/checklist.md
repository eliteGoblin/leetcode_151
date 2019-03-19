
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


## General

*  -- 写成 ++
*  doc写的 type mismatch, 没有取数组下标
*  取数组下标赋值时切记判断是否overflow, 比如只想初始化数组，res[0] = xx res[1] = xx 确保数组len