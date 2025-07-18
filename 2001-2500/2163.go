func minimumDifference(nums []int) int64 {

n := len(nums) / 3



maxH, minH := maxHeap(make([]int, 0, n)), minHeap(make([]int, 0, n))

leftSum, rightSum := 0, 0

for i := 0; i < n; i++ {

leftSum += nums[i]

maxH = append(maxH, nums[i])



j := len(nums) - i - 1

rightSum += nums[j]

minH = append(minH, nums[j])

}



heap.Init(&maxH)

heap.Init(&minH)



diff := make([]int, n+1)

for i := range diff {

diff[i] += leftSum

diff[len(diff)-i-1] -= rightSum



l, r := n+i, len(nums)-n-i-1

heap.Push(&maxH, nums[l])

heap.Push(&minH, nums[r])



leftSum += nums[l]

leftSum -= heap.Pop(&maxH).(int)



rightSum += nums[r]

rightSum -= heap.Pop(&minH).(int)

}



res := math.MaxInt64

for _, d := range diff {

if d < res {

res = d

}

}



return int64(res)

}



type minHeap []int



func (h minHeap) Len() int { return len(h) }

func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h minHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *minHeap) Pop() (v interface{}) {

v = (*h)[len(*h)-1]

*h = (*h)[:len(*h)-1]



return v

}



type maxHeap []int



func (h maxHeap) Len() int { return len(h) }

func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h maxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *maxHeap) Pop() (v interface{}) {

v = (*h)[len(*h)-1]

*h = (*h)[:len(*h)-1]



return v

}
