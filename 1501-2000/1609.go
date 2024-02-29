func isEvenOddTree(root *TreeNode) bool {
    q := make([]*TreeNode, 0)
    q = append(q, root)
    level := 0
    
    for len(q) > 0 {
        size := len(q)
        
        prev := 0      
        if level%2 == 1 {
            prev = math.MaxInt32
        }
        
        for i := 0; i < size; i++ {
            cur := q[0]
            q = q[1:]
            if level%2 == 0 && (cur.Val%2 == 0 || cur.Val <= prev) {
                return false
            }
            
            if level%2 ==1 && (cur.Val%2 == 1 || cur.Val >= prev) {
                return false
            }
            
            prev = cur.Val
            if cur.Left != nil {
                q = append(q, cur.Left)
            }
            
            if cur.Right != nil {
                q = append(q, cur.Right)
            }
        
        }
        
        level++
    }
    
    return true
}
