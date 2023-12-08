package mergetwosortedlists

import "testing"

func TestMergeTwoSortedList(t *testing.T) {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}

	mergedList := mergeTwoLists(list1, list2)
	printList(mergedList)
}
