package main

import (
	"fmt"
	"math"
)

type SegmentTree struct {
	sumTree []int
	minTree []int
	maxTree []int
	arr     []int
	n       int
}

func (segTree *SegmentTree) buildSum(node int, start int, end int) {
	if start == end {
		(*segTree).sumTree[node] = (*segTree).arr[start]
	} else {
		mid := start + (end-start)/2
		left := 2*node + 1
		right := 2*node + 2
		(*segTree).buildSum(left, start, mid)
		(*segTree).buildSum(right, mid+1, end)
		(*segTree).sumTree[node] = (*segTree).sumTree[left] + (*segTree).sumTree[right]
	}
}
func (segTree *SegmentTree) buildMin(node int, start int, end int) {
	if start == end {
		(*segTree).minTree[node] = (*segTree).arr[start]
	} else {
		mid := start + (end-start)/2
		left := 2*node + 1
		right := 2*node + 2
		segTree.buildMin(left, start, mid)
		segTree.buildMin(right, mid+1, end)
		segTree.minTree[node] = min(segTree.minTree[left], segTree.minTree[right])
	}
}

func (segTree *SegmentTree) buildMax(node int, start int, end int) {
	if start == end {
		segTree.maxTree[node] = segTree.arr[start]
	} else {
		mid := start + (end-start)/2
		left := 2*node + 1
		right := 2*node + 2
		segTree.buildMax(left, start, mid)
		segTree.buildMax(right, mid+1, end)
		segTree.maxTree[node] = max(segTree.maxTree[left], segTree.maxTree[right])
	}
}
func (segTree *SegmentTree) queryMax(node int, start int, end int, L int, R int) int {
	if R < start || end < L {
		return math.MinInt64
	}
	if L <= start && end <= R {
		return segTree.maxTree[node]
	}
	mid := start + (end-start)/2
	left := 2*node + 1
	right := 2*node + 2
	return max(segTree.queryMax(left, start, mid, L, R), segTree.queryMax(right, mid+1, end, L, R))
}
func (segTree *SegmentTree) queryMin(node int, start int, end int, L int, R int) int {
	if R < start || end < L {
		return math.MaxInt64
	}
	if L <= start && end <= R {
		return segTree.minTree[node]
	}
	mid := start + (end-start)/2
	left := 2*node + 1
	right := 2*node + 2
	return min(segTree.queryMin(left, start, mid, L, R), segTree.queryMin(right, mid+1, end, L, R))
}
func (segTree *SegmentTree) querySum(node int, start int, end int, L int, R int) int {
	if R < start || end < L {
		return 0
	}
	if start <= L && end <= R {
		return segTree.sumTree[node]
	}
	mid := start + (end-start)/2
	left := 2*node + 1
	right := 2*node + 2
	return segTree.querySum(left, start, mid, L, R) + segTree.querySum(right, mid+1, end, L, R)
}
func (segTree *SegmentTree) updateSum(node int, start int, end int, index int, value int) {
	if start == end {
		segTree.sumTree[node] = value
	} else {
		mid := start + (end-start)/2
		left := 2*node + 1
		right := 2*node + 2
		if index <= mid {
			segTree.updateSum(left, start, mid, index, value)
		} else {
			segTree.updateSum(right, mid+1, end, index, value)
		}
		segTree.sumTree[node] = segTree.sumTree[left] + segTree.sumTree[right]
	}
}
func (segTree *SegmentTree) updateMin(node int, start int, end int, index int, value int) {
	if start == end {
		segTree.minTree[node] = value
	} else {
		mid := start + (end-start)/2
		left := 2*node + 1
		right := 2*node + 2
		if index <= mid {
			segTree.updateMin(left, start, mid, index, value)
		} else {
			segTree.updateMin(right, mid+1, end, index, value)
		}
		segTree.minTree[node] = min(segTree.minTree[left], segTree.minTree[right])
	}
}
func (segTree *SegmentTree) updateMax(node int, start int, end int, index int, value int) {
	if start == end {
		segTree.maxTree[node] = value
	} else {
		mid := start + (end-start)/2
		left := 2*node + 1
		right := 2*node + 2
		if index <= mid {
			segTree.updateMax(left, start, mid, index, value)
		} else {
			segTree.updateMax(right, mid+1, end, index, value)
		}
		segTree.maxTree[node] = max(segTree.maxTree[left], segTree.maxTree[right])
	}
}
func NewSegmentTree(arr []int) *SegmentTree {
	segTree := &SegmentTree{
		arr:     arr,
		n:       len(arr),
		sumTree: make([]int, 4*len(arr)),
		minTree: make([]int, 4*len(arr)),
		maxTree: make([]int, 4*len(arr)),
	}
	segTree.buildSum(0, 0, segTree.n-1)
	segTree.buildMin(0, 0, segTree.n-1)
	segTree.buildMax(0, 0, segTree.n-1)
	return segTree
}
func (segTree *SegmentTree) rangeSum(L int, R int) int {
	return segTree.querySum(0, 0, segTree.n-1, L, R)
}
func (segTree *SegmentTree) rangeMax(L int, R int) int {
	return segTree.queryMax(0, 0, segTree.n-1, L, R)
}
func (segTree *SegmentTree) rangeMin(L int, R int) int {
	return segTree.queryMin(0, 0, segTree.n-1, L, R)
}
func (segTree *SegmentTree) updateValue(index int, value int) {
	segTree.updateMax(0, 0, segTree.n-1, index, value)
	segTree.updateMin(0, 0, segTree.n-1, index, value)
	segTree.updateSum(0, 0, segTree.n-1, index, value)
}

func main() {
	var n int
	fmt.Scan(&n)
	var arr []int
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		arr = append(arr, x)
	}
	segtree := NewSegmentTree(arr)
	ansSum := segtree.rangeSum(1, 4)
	ansMax := segtree.rangeMax(1, 4)
	ansMin := segtree.rangeMin(1, 4)
	fmt.Println("The range sum in the array from index 1 to index 4 : ", ansSum)
	fmt.Println("The range max in the array from index 1 to index 4 : ", ansMax)
	fmt.Println("The range min in the array from index 1 to index 4 : ", ansMin)
	fmt.Print("\n\nArray has been updated\n\n\n")
	segtree.updateValue(3, 10)
	ansSum = segtree.rangeSum(1, 4)
	ansMax = segtree.rangeMax(1, 4)
	ansMin = segtree.rangeMin(1, 4)
	fmt.Println("The range sum in the updated array from index 1 to index 4 : ", ansSum)
	fmt.Println("The range max in the updated array from index 1 to index 4 : ", ansMax)
	fmt.Println("The range min in the updated array from index 1 to index 4 : ", ansMin)
}
