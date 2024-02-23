type Path struct {
    city  int
    price int
    k     int
}

type minHeap []Path
func (h minHeap) Len() int { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].price < h[j].price }
func (h minHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(Path)) }
func (h *minHeap) Pop() interface{} {
    tmp := (*h)[len(*h) - 1]
    *h = (*h)[:len(*h) - 1]
    return tmp
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
    cache := make(map[int][][]int)
    for _, v := range flights {
        if _, ok := cache[v[0]]; ok {
            cache[v[0]] = append(cache[v[0]], []int{v[1], v[2]})
        } else {
            cache[v[0]] = [][]int{[]int{v[1], v[2]}}
        }
    }
    h := &minHeap{}
    heap.Init(h)
    heap.Push(h, Path{src, 0, K + 1})
    for h.Len() > 0 {
        path := heap.Pop(h).(Path)
        if path.city == dst { return path.price }
        if path.k > 0 {
            for _, dstAndPrice := range cache[path.city] {
                heap.Push(h, Path{dstAndPrice[0], path.price + dstAndPrice[1], path.k - 1})
            }
        }
    }
    return -1
}
