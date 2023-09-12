package main

// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 返回容器可以储存的最大水量。
// 说明：你不能倾斜容器。

// 思路：使用两个指针一次从左向右遍历 记最左边长度为left 遍历时记录最长长度的索引 有以下情况
// cur小于left 继续向前
// cur大于left 判断是否大于最大索引
// 想到这里 感觉这种思路比较复杂 看了下解答 可以让指针从左右两边向中间收缩
// 计算当前面积 并比较左右大小 较小的那边向中间收缩
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	max := 0
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		if area > max {
			max = area
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return max
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
// 证明部分 大致意思是 围成的面积在索引向中间收缩时 取决于较小值的长度
// 改变较小值 则面积可能会增大 但是改变较大值 面积则一定不会增大 只会不变或减小 所以要较小的那边向中间收缩