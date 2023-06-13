package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}

	if l1 == nil {
		l1 = &ListNode{
			Val:  0,
			Next: nil,
		}
	}

	if l2 == nil {
		l2 = &ListNode{
			Val:  0,
			Next: nil,
		}
	}

  retCarry := 0
	val := l1.Val + l2.Val + carry

	if val >= 10 {
		val -= 10
		retCarry = 1
	}

	return &ListNode{
		Val:  val,
		Next: addTwoNumbers(l1.Next, l2.Next, retCarry),
	}
}

func makeList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	return &ListNode{
		Val:  nums[0],
		Next: makeList(nums[1:]),
	}
}

func toInt(l *ListNode) int {
	if l.Next == nil {
		return l.Val
	}

	return (toInt(l.Next) * 10) + l.Val
}
