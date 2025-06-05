type UnionFind struct{
    parent []int
}

func newUnionFind() *UnionFind{
    parents := make([]int, 26)
    for i :=0; i < 26; i++{
        parents[i] = i
    }
    
    return &UnionFind{
        parent: parents,
    }
}

func (u *UnionFind) find(x int) int{
    if u.parent[x] != x{
        u.parent[x] = u.find(u.parent[x])
    }
    
    return u.parent[x]
}

func (u *UnionFind) union (x, y int){
    xr, yr := u.find(x), u.find(y)
    
    if xr < yr{
        u.parent[yr] = xr
    } else{
        u.parent[xr] = yr
    }
}

func smallestEquivalentString(A string, B string, S string) string {
    uf := newUnionFind()
    for i :=0; i < len(A); i++{
        uf.union(int(A[i]-'a'), int(B[i]-'a'))
    }
    
    res := make([]byte, 0,len(S))
    
    for i :=0; i < len(S); i++{
        res =append(res, byte(uf.find(int(S[i]-'a'))+'a'))
    }
    
    return string(res)
}
