package data

import (
	"container/heap"
	"github.com/vsliouniaev/aoc/util/nav"
)

//https: //pkg.go.dev/container/heap

// PointHeap can bs used in combination with a points on a grid to help with djikstra path-finding (See 2021/15)
type PointHeap struct {
	Heap  []PointVal
	Index map[nav.Point]int
}

type PointVal struct {
	Point nav.Point
	Val   int
}

func NewPointHeap() *PointHeap {
	return &PointHeap{
		Index: make(map[nav.Point]int),
	}
}

func (h PointHeap) Len() int           { return len(h.Heap) }
func (h PointHeap) Less(i, j int) bool { return h.Heap[i].Val < h.Heap[j].Val }
func (h PointHeap) Swap(i, j int) {
	h.Index[h.Heap[i].Point] = j
	h.Index[h.Heap[j].Point] = i
	h.Heap[i], h.Heap[j] = h.Heap[j], h.Heap[i]
}

func (h *PointHeap) Set(pv PointVal) {
	i, ok := h.Index[pv.Point]
	if !ok {
		panic("Set is not a push")
	}
	h.Heap[i] = pv
	heap.Fix(h, i)
}

func (h *PointHeap) Delete(p nav.Point) {
	i := h.Index[p]
	delete(h.Index, p)
	heap.Remove(h, i)
}

func (h PointHeap) ContainsPoint(x nav.Point) bool {
	_, ok := h.Index[x]
	return ok
}

func (h *PointHeap) Push(x interface{}) {
	v := x.(PointVal)
	h.Heap = append(h.Heap, v)
	h.Index[v.Point] = len(h.Heap) - 1
}

func (h *PointHeap) Pop() interface{} {
	old := h.Heap
	n := len(old)
	x := old[n-1]
	h.Heap = old[0 : n-1]
	delete(h.Index, x.Point)
	return x
}
