package main

// 45. 跳跃游戏 II

func jump(nums []int) int {
	start, end, step := 0, 1, 0
	for end < len(nums) {
		maxPos := 0
		for i := start; i < end; i++ {
			if i+nums[i] > maxPos {
				maxPos = i + nums[i]
			}
		}
		start = end
		end = maxPos + 1
		step++
	}
	return step
}

func jump2(nums []int) int {
	end, maxPos, step := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > maxPos {
			maxPos = i + nums[i]
		}
		if i == end {
			end = maxPos
			step++
		}
	}
	return step
}

func main() {

}
