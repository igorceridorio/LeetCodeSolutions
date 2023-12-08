package addtwonumbers

import (
	"fmt"
	"math/big"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func makeListFromString(s string) *ListNode {
	list := &ListNode{}
	tail := list

	for i := 0; i < len(s); i++ {
		number, _ := strconv.Atoi(string(s[i]))
		tail.Next = &ListNode{Val: number}
		tail = tail.Next
	}

	return list.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1str := ""
	l2str := ""

	for l1 != nil {
		l1str += strconv.Itoa(l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		l2str += strconv.Itoa(l2.Val)
		l2 = l2.Next
	}

	l1str = reverse(l1str)
	l2str = reverse(l2str)

	l1number := new(big.Int)
	l1number, _ = l1number.SetString(l1str, 10)
	l2number := new(big.Int)
	l2number, _ = l2number.SetString(l2str, 10)

	sum := big.NewInt(0)
	sum = sum.Add(l1number, l2number)

	sumRev := reverse(sum.String())
	return makeListFromString(sumRev)
}
