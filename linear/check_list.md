
## Linear

### Array

#### 26

*  思路不清晰，第一次一直出错


#### 33. Search in Rotated Sorted Array

*  思路混乱，WA
*  什么是rotate: 循环左移或者右移，结果是一样的，不分左右
*  讲解:  http://www.cnblogs.com/grandyang/p/4325648.html

81 Search in Rotated Sorted Array 特殊情况处理


#### 4. Median of Two Sorted Arrays

median:  middle element when n is odd; average of middle two elements when n is even

*  step 1: 理解两个same sized array，find median  https://www.geeksforgeeks.org/median-of-two-sorted-arrays/
*  step 2: 两个different sized array, find median, find kth element of sorted array
    +  https://www.geeksforgeeks.org/k-th-element-two-sorted-arrays/
    +  https://www.geeksforgeeks.org/median-of-two-sorted-arrays-of-different-sizes/
    +  https://leetcode.com/articles/median-of-two-sorted-arrays/?page=2


#### 15 3sum

https://www.geeksforgeeks.org/find-a-triplet-that-sum-to-a-given-value/
两个关键:

*  ksum: 用 O(n^(k-1))复杂度找到
*  不重复的三元组: 如何跳过重复元素: leetcpde.pdf 有些判断不必要，见自己提交的code.
    +  仅当找到解的时候，需要跳到下一个，否则如果不是解，肯定要往前或者往后走，并不会重复, 重复的case: [-2,0,0,2,2]



#### 18 4sum

*  夹逼原理类似，再一次练习: 去重逻辑，什么情况index因重复滑动


#### 31. Next Permutation

*  没思路，数学问题

#### 42. Trapping Rain Water

*  求凹陷区域和的思路

#### 48. Rotate Image

*  翻转替代旋转的算法
*  翻转时 i, j 的范围。全部翻转，相当于不变

#### 70. Climbing Stairs

*  思路
*  TLE问题


#### 135. Candy

*  思路错误：自己找最小值，但最小值可以多个，如 [1, 3, 2, 2, 1]

#### 137. Single Number II

*  无思路，除了用O(n) space
*  思路一: golang int OJ 上64bit,不同于playground上的32bit, int(unsafe.Sizeof(var)) * 8 来求
*  

[[LeetCode] Single Number II 单独的数字之二](http://www.cnblogs.com/grandyang/p/4263927.html)  
[Find the element that appears once](https://www.geeksforgeeks.org/find-the-element-that-appears-once/)  


