package main

// 给你一个整数数组 nums ，判断是否存在三元组 
// [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
// 你返回所有和为 0 且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。

func threeSum(nums []int) [][]int {
	first := 0
	res := make([][]int, 0)
	for first < len(nums) - 2 {
		leftValue := nums[first]
		mid := first + 1
		for mid < len(nums) - 1 {
			midValue := nums[mid]
			last := mid + 1
			for last < len(nums) {
				lastValue := nums[last]
				if leftValue + midValue + lastValue == 0 {
					cur := []int{leftValue, midValue, lastValue}
					if !checkExist(res, cur) {
						res = append(res, cur)
					}
				}
				last++
			}
			mid++
		}
		first++
	}
	return res
}

func checkExist(res [][]int, cur []int) bool {
	for _, arr := range res {
		l := arr[0]
		m := arr[1]
		r := arr[2]
		count := 0
		for _, v := range cur {
			if v != l && v != m && v != r {
				count++
			}
		}
		if count == 0 {
			return true
		}
	}
	return false
}