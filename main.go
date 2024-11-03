package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
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
		Next: AddTwoNumbers(l1.Next, l2.Next, retCarry),
	}
}

func MakeList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	return &ListNode{
		Val:  nums[0],
		Next: MakeList(nums[1:]),
	}
}

func ToInt(l *ListNode) int {
	if l.Next == nil {
		return l.Val
	}

	return (ToInt(l.Next) * 10) + l.Val
}

func ListsEqual(l1 *ListNode, l2 *ListNode) bool {
  if l1 == nil && l2 == nil {
    return true
  }
  if l1 == nil && l2 != nil {
    return false
  }
  if l1 != nil && l2 == nil {
    return false
  }
  if l1.Val != l2.Val {
    return false
  }
  return ListsEqual(l1.Next, l2.Next)
}
