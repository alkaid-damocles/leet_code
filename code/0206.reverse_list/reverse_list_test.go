package code

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	l1 := ListNode{Val: 6}
	l2 := ListNode{Val: 9, Next: &l1}
	l1.Next = &l2
	fmt.Println(reverseList(&l1))
}

func TestXxx(t *testing.T) {

}
