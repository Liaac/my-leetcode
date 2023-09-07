package main

// 2023.9.7
// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

// 考虑使用一个map key为nums中的每个num value为当前num序列的长度
// 对nums进行遍历 对当前的num 寻找map中是否有相邻num-1和num+1
// 如果存在 比如存在num-1 它的value读出来是3 即序列长度为3 那么就将num num-1到num-3的value全更新为3+1=4 再继续看num+1 同理
// 如果不存在 则使其为1

// 以下为第一次写的代码
// 反思一下  又考虑复杂了 而且逻辑有些问题 《《《《《《《《《《《《《《——————————————————————————————
// 我考虑了每个数 在遍历时向左或者向右 但实际上只需要在nums的所有num记录之后 从最左侧即num-1不存在的数开始向右统计即可《《《《《《《《《《————————————
// type Way int

// const (
// 	LEFT Way = iota
// 	RIGHT
// 	BOTH
// )

// func longestConsecutive(nums []int) int {
// 	numMap := make(map[int]int, 0)
// 	max := 0
// 	for _, value := range nums {
// 		if _, ok := numMap[value]; !ok {
// 			numMap[value] = 1
// 		}
// 		dealLink(numMap, value, BOTH)
// 		if numMap[value] > max {
// 			max = numMap[value]
// 		}
// 	}
// 	return max
// }

// func dealLink(numMap map[int]int, value int, way Way) {
// 	pre := value - 1
// 	next := value + 1
// 	_, preok := numMap[pre]
// 	if way != RIGHT && preok{
// 		len := numMap[pre]
// 		numMap[value] += len
// 		dealLink(numMap, value - len, LEFT)
// 		for i := 1; i <= len; i++ {
// 			numMap[value - i] = numMap[value]
// 		}
// 	}
// 	_, nextok := numMap[next]
// 	if way != LEFT && nextok {
// 		len := numMap[next]
// 		numMap[value] += len
// 		dealLink(numMap, value + len, RIGHT)
// 		for i := 1; i <= len; i++ {
// 			numMap[value + i] = numMap[value]
// 		}
// 	}
// }

func longestConsecutive(nums []int) int {
	numMap := make(map[int]bool, 0)
	for _, num := range nums {
		numMap[num] = true
	}
	max := 0
	for num := range numMap {
		if !numMap[num-1] {
			curLen := 1
			curNum := num
			for numMap[curNum+1] {
				curLen++
				curNum++
			}
			if curLen > max {
				max = curLen
			}
		}
	}
	return max
}