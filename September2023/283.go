package main

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。

// 思路：使用left记录最左边的0值 向后遍历时遇到不为0的值时 将其移到left处
func moveZeroes(nums []int) {
	len := len(nums)
	left := -1
	for i := 0; i < len; i++ {
		if left >= 0 && nums[i] != 0 {
			swap(nums, left, i)
			left += 1
		} else {
			if nums[i] == 0 && left == -1 {
				left = i
			}
		}
	}
}

func swap(nums []int, left, right int) {
	temp := nums[left]
	nums[left] = nums[right]
	nums[right] = temp
}

// 另一种解法可以使用快慢指针 快指针前进
// 前进时若当前数不为0 与赋值给慢指针索引位置 慢指针前进
// func moveZeroes2(nums []int)  {
// 	slow, fast := 0, 0
// 	for fast < len(nums) {
// 		if nums[fast] != 0 {
// 			nums[slow] = nums[fast]
// 			slow++
// 		}
// 		fast++
// 	}
// 	for slow < len(nums) {
// 		nums[slow] = 0
// 		slow++
// 	}
// }
