package main

// // 两数之和
// 给定一个整数数组nums和一个整数目标值target，请你在该数组中找出和为目标值target的那两个整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 你可以按任意顺序返回答案。

func twoSum(nums []int, target int) []int {
	sumMap := make(map[int]int)
	for index, value := range nums {
		need := target - value
		if p, ok := sumMap[need]; ok {
			return []int{p, index}
		}
		sumMap[value] = index
	}
	return []int{}
}
