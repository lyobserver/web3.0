package questions

/*
46. 全排列：给定一个不含重复数字的数组 `nums` ，返回其所有可能的全排列。可以使用回溯算法，定义一个函数来进行递归操作，
在函数中通过交换数组元素的位置来生成不同的排列，使用 `for` 循环遍历数组，每次选择一个元素作为当前排列的第一个元素，然后递归调用函数处理剩余的元素。
*/
func Permute(nums []int) [][]int {
	var result [][]int
	backtrack(nums, 0, &result)
	return result
}

// backtrack 是回溯函数，用于递归生成全排列
func backtrack(nums []int, start int, result *[][]int) {
	if start == len(nums) {
		// 复制当前排列到结果集中
		temp := make([]int, len(nums))
		copy(temp, nums)
		*result = append(*result, temp)
		return
	}

	for i := start; i < len(nums); i++ {
		// 交换元素位置
		nums[start], nums[i] = nums[i], nums[start]
		// 递归处理剩余元素
		backtrack(nums, start+1, result)
		// 回溯操作，恢复元素位置
		nums[start], nums[i] = nums[i], nums[start]
	}
}
