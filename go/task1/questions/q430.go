package questions

/*
430. 扁平化多级双向链表：多级双向链表中，除了指向下一个节点和前一个节点指针之外，它还有一个子链表指针，可能指向单独的双向链表。
这些子列表也可能会有一个或多个自己的子项，依此类推，生成多级数据结构，如下面的示例所示。给定位于列表第一级的头节点，请扁平化列表，即将这样的多级双向链表展平成普通的双向链表，使所有结点出现在单级双链表中。
可以定义一个结构体来表示链表节点，包含 `val`、`prev`、`next` 和 `child` 指针，然后使用递归的方法来扁平化链表，先处理当前节点的子链表，再将子链表插入到当前节点和下一个节点之间。
*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */
type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func Flatten(root *Node) *Node {
	flattenTail(root)
	return root
}

// flattenTail 辅助函数，递归扁平化链表并返回链表的尾节点
func flattenTail(head *Node) *Node {
	if head == nil {
		return nil
	}
	if head.Child == nil {
		if head.Next == nil {
			return head
		}
		// 递归处理下一个节点
		return flattenTail(head.Next)
	}
	child := head.Child
	next := head.Next
	// 断开当前节点与子节点的连接
	head.Child = nil

	// 递归处理子链表
	childTail := flattenTail(child)

	// 将子链表插入到当前节点和下一个节点之间
	head.Next = child
	child.Prev = head
	if next != nil {
		childTail.Next = next
		next.Prev = childTail
		// 递归处理下一个节点
		return flattenTail(next)
	}
	return childTail
}
