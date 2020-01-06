Qu.  Convert Binary Number in a Linked List to Integer

Given head which is a reference node to a singly-linked list. The value of each node in the linked list is either 0 or 1. The linked list holds the binary representation of a number.

Return the decimal value of the number in the linked list.

Ans --------

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
    i := 0
    a := head
    for a != nil {
        i++
        a = a.Next
    }
    k := 0
    for head != nil {
        if head.Val != 0 {
            k += int(math.Pow(2,float64(i-1)))
        }
        i = i-1
        head = head.Next
    }
    return k
}
