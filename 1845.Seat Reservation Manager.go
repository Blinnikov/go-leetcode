// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type SeatManager struct {
    h *IntHeap
}

func Constructor(n int) SeatManager {
    sm := new(SeatManager)

    sm.h = &IntHeap{}

    heap.Init(sm.h)

    for i := 1; i <= n; i++ {
		heap.Push(sm.h, i)
	}

    return *sm
}


func (this *SeatManager) Reserve() int {
    res := heap.Pop(this.h).(int)
    return res
}


func (this *SeatManager) Unreserve(seatNumber int)  {
    heap.Push(this.h, seatNumber)
}


/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */
