package heap

func NewMaxHeap(input []int) *MaxHeap {
	return buildMaxHeap(input)
}

type MaxHeap struct {
	Buf []int
}

func (m *MaxHeap) Size() int {
	return len(m.Buf)
}

// MaxHeapify
// T(n) = O(logn). Alternatively, we can characterize the running time of MAXHEAPIFY on a node of height h as O(h).
// Since height of heap is logn.
func (m *MaxHeap) MaxHeapify(i int) {
	if m.Size() == 0 {
		return
	}

	l := left(i)  // Left child
	r := right(i) // Right child

	largest := i
	if l < m.Size() && m.Buf[l] > m.Buf[i] {
		largest = l
	}
	if r < m.Size() && m.Buf[r] > m.Buf[largest] {
		largest = r
	}

	if largest != i {
		// Exchange A[i] with A[largest].
		tmp := m.Buf[i]
		m.Buf[i] = m.Buf[largest]
		m.Buf[largest] = tmp

		// Call maxheapify with largest
		m.MaxHeapify(largest)
	}
}

// T(n) = O(n*logn). Since O(n) iterations of MaxHeapify(logn)
func buildMaxHeap(input []int) *MaxHeap {
	if len(input) == 0 {
		return &MaxHeap{}
	}

	m := &MaxHeap{Buf: input}

	// Call MaxHeapify for all non-leaves starting with the last leaf.
	startLeafIndex := len(input) / 2 // For 3 elements, leaves start at 1. For 4, leaves start from 2.
	for i := startLeafIndex; i >= 0; i-- {
		m.MaxHeapify(i)
	}

	return m
}

// Helper functions

func left(i int) int {
	return i<<1 + 1 // 2*i+1
}

func right(i int) int {
	return i<<1 + 2 // 2*i+2
}
