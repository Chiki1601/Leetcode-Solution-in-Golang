type Solution struct {
    node *ListNode
}


/** @param head The linked list's head.
        Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
    return Solution{node: head}
}


/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
    ptr := this.node
    res := this.node.Val
    cnt := 1
    for ptr.Next != nil {
        if rand.Intn(cnt + 1) == 0 {
            res = ptr.Next.Val
        }
        cnt++
        ptr = ptr.Next
    }
    return res
}
