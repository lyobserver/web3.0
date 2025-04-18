package task1

import "fmt"

/*
任务1：
136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
可以使用 `for` 循环遍历数组，结合 `if` 条件判断和 `map` 数据结构来解决，例如通过 `map` 记录每个元素出现的次数，然后再遍历 `map` 找到出现次数为1的元素。
*/

//方法1
func singleNumber1(nums []int) int {
	var result int
	for _, num := range nums {
		result ^= num
	}
	return result
}

//方法2
func singleNumber2(nums []int) int {
	var result int
	var numCount = make(map[int]int)
	for _, num := range nums {
		numCount[num]++
	}
	for num, count := range numCount {
		if count == 1 {
			result = num
			break
		}
	}
	return result
}

func main() {
	nums := []int{2, 2, 1}
	fmt.Println("{2, 2, 1}方法1=", singleNumber1(nums))
	fmt.Println("{2, 2, 1}方法2=", singleNumber2(nums))
	nums1 := []int{4, 1, 2, 1, 2}
	fmt.Println("{4, 1, 2, 1, 2}方法1=", singleNumber1(nums1))
	fmt.Println("{4, 1, 2, 1, 2}方法2=", singleNumber2(nums1))
	nums2 := []int{1}
	fmt.Println("{1}方法1=", singleNumber1(nums2))
	fmt.Println("{1}方法2=", singleNumber2(nums2))
}
