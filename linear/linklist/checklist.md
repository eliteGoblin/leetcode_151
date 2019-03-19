## 2. Add Two Numbers

*  简单粗暴方法一次AC，下次采用： 一个for循环搞定，而不是 l1 && l2，然后再来一次for.

## 92. Reverse Linked List II

*  一次AC，值得重做，要点
    +  头插法
    +  伪头部

## 61. Rotate List

*  一直WA, 各种小问题

## 25 Reverse Nodes in k-Group

*  迭代: 大致思路有，做不出来。思路有小问题
    +  必须先检查下一段是否足够k节点，不能错误的优化：希望一次搞定：逆序+向后移动
    +  自己的代码目前最优，供之后参考
*  递归：没想到

## 138 Copy List with Random Pointer

*  自己只有o(n^2), naive思路
*  答案提供了inplace的o(n)思路

##  142. Linked List Cycle II

扩展：　有环快慢指针一定能追上吗？

最容易理解的思路:
*  先求出环长：快慢指针处遍历一次，相遇前走的距离
*  一前以后指针，每次一step，直到相遇，相遇处便是环入口
[判断单链表是否有环、求环长和环入口最优算法](https://www.tomorrow.wiki/archives/1886)

## 143. Reorder List

*  思路有，WA
*  思路啰嗦, 对比最佳答案


## 146. LRU Cache

*  第一次没思路，WA，细节
*  双链表的不同
    +  心里有数，修改哪些link: insert, 四个；　cur.next cur.pre;pre.next;next.pre
    +  记录tail时，任何修改node的: move, delete都需要判断tail是否需要变化(自己就是在moveToHead时，忘记修改tail)

## 460. LFU Cache

*  没做过，一起练习