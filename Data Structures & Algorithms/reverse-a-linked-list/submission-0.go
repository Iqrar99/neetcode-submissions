/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

    currNode := head
	var prevNode *ListNode
	for currNode.Next !=  nil {
		nextNode := currNode.Next
		currNode.Next = prevNode
		prevNode = currNode
		currNode = nextNode
	}
	if currNode.Next == nil {
		currNode.Next = prevNode
	}

	return currNode
}
