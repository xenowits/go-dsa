package maxprioqueue

import (
	"errors"
	"go-dsa/heap"
	"math"
)

func NewMaxPriorityQueue(input []int) *MaxPriorityQueue {
	return &MaxPriorityQueue{
		heap: heap.NewMaxHeap(input),
	}
}

type MaxPriorityQueue struct {
	heap *heap.MaxHeap
}

func (q *MaxPriorityQueue) HeapMaximum() (int, error) {
	if q.Empty() {
		return 0, errors.New("empty queue")
	}

	return q.heap.Buf[0], nil
}

// HeapExtractMax has T(n) = O(logn) since it performs only a constant amount of work on top of the O(logn) time for MAX-HEAPIFY.
func (q *MaxPriorityQueue) HeapExtractMax() (int, error) {
	if q.Empty() {
		return 0, errors.New("heap underflow")
	}

	// Grab the max value. It is the root node.
	max := q.heap.Buf[0]

	// Magic starts here. We swap the root node to the last leaf node.
	q.heap.Buf[0] = q.heap.Buf[q.heap.Size()-1]
	// Then call max heapify on the new root node.
	q.heap.MaxHeapify(0)

	return max, nil
}

func (q *MaxPriorityQueue) HeapIncreaseKey(index, key int) error {
	if key < q.heap.Buf[index] {
		return errors.New("new key is smaller than current key")
	}

	q.heap.Buf[index] = key

	for index > 0 {
		if q.heap.Buf[parent(index)] >= q.heap.Buf[index] {
			break
		}

		// Child is bigger than the parent. Cannot be. MUSTN'T be. Swap them asap.
		tmp := q.heap.Buf[parent(index)]
		q.heap.Buf[parent(index)] = q.heap.Buf[index]
		q.heap.Buf[index] = tmp

		index = parent(index)
	}

	return nil
}

func (q *MaxPriorityQueue) MaxHeapInsert(key int) error {
	q.heap.Buf = append(q.heap.Buf, math.MinInt) // Our heap just got bigger!

	err := q.HeapIncreaseKey(q.heap.Size()-1, key)
	if err != nil {
		return err
	}

	return nil
}

func (q *MaxPriorityQueue) Empty() bool {
	if q.heap == nil || len(q.heap.Buf) == 0 {
		return true
	}

	return false
}

// Helper functions

func parent(i int) int {
	// Parent of 2 is 0, not 1
	if i&1 == 0 {
		return i>>1 - 1
	}

	// Parent of 1 is 0
	return i >> 1
}
