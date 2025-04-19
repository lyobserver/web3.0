package main

import (
	"fmt"
	"main/task1/questions"
)

func main() {

	nums := []int{2, 2, 1}
	fmt.Println("{2, 2, 1}方法1=", questions.SingleNumber1(nums))
	fmt.Println("{2, 2, 1}方法2=", questions.SingleNumber2(nums))
	nums1 := []int{4, 1, 2, 1, 2}
	fmt.Println("{4, 1, 2, 1, 2}方法1=", questions.SingleNumber1(nums1))
	fmt.Println("{4, 1, 2, 1, 2}方法2=", questions.SingleNumber2(nums1))
	nums2 := []int{1}
	fmt.Println("{1}方法1=", questions.SingleNumber1(nums2))
	fmt.Println("{1}方法2=", questions.SingleNumber2(nums2))
	q198s := []int{2, 9, 1, 3, 7, 8}
	fmt.Println("打家劫舍：", questions.Rob(q198s))

	var list1 questions.ListNode = questions.ListNode{Val: 1, Next: &questions.ListNode{Val: 2, Next: &questions.ListNode{Val: 4}}}
	var list2 questions.ListNode = questions.ListNode{Val: 1, Next: &questions.ListNode{Val: 3, Next: &questions.ListNode{Val: 4}}}
	var list3 *questions.ListNode = questions.MergeTwoLists(&list1, &list2)
	fmt.Println("合并两个有序链表：")
	for list3 != nil {
		fmt.Print(list3.Val, " ")
		list3 = list3.Next
	}

	var ordernums = []int{1, 2, 3}
	fmt.Println("全排列：", questions.Permute(ordernums))

	var reverseString = []byte{'h', 'e', 'l', 'l', 'o'}
	questions.ReverseString(reverseString)
	fmt.Println("反转字符串：", string(reverseString))

	var x = 8
	fmt.Println("x的平方根：", questions.MySqrt(x))

	var dupnums = []int{1, 1, 2, 2, 3, 4, 4, 5}
	dupnums = dupnums[0:questions.RemoveDuplicates(dupnums)]
	fmt.Println("删除有序数组中的重复项：", dupnums)

	var intervals = [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println("合并区间：", questions.Merge(intervals))

	node1 := &questions.Node{Val: 1}
	node2 := &questions.Node{Val: 2}
	node3 := &questions.Node{Val: 3}
	node4 := &questions.Node{Val: 4}
	node5 := &questions.Node{Val: 5}

	node1.Next = node2
	node2.Prev = node1
	node2.Next = node3
	node3.Prev = node2

	node1.Child = node4
	node4.Next = node5
	node5.Prev = node4
	flattenedList := questions.Flatten(node1)
	fmt.Println("扁平化多级双向链表：")
	for flattenedList != nil {
		fmt.Print(flattenedList.Val, " ")
		flattenedList = flattenedList.Next
	}
	fmt.Print("\n")

	calendar := questions.Constructor()
	fmt.Println(calendar.Book(10, 20)) // 返回 true
	fmt.Println(calendar.Book(15, 25)) // 返回 false，因为与已有日程重叠
	fmt.Println(calendar.Book(20, 30)) // 返回 true
}
