package maxprioqueue

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaxPriorityQueue(t *testing.T) {
	input := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}
	q := NewMaxPriorityQueue(input)
	maxExpected, err := q.HeapMaximum()
	require.NoError(t, err)
	require.Equal(t, maxExpected, 16)

	maxExpected, err = q.HeapExtractMax()
	require.NoError(t, err)
	require.Equal(t, maxExpected, 16)

	// Extract again
	maxExpected, err = q.HeapExtractMax()
	require.NoError(t, err)
	require.Equal(t, maxExpected, 14)
}

func TestMaxPriorityQueueInsert(t *testing.T) {
	input := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}
	q := NewMaxPriorityQueue(input)
	err := q.MaxHeapInsert(21)
	require.NoError(t, err)

	maxExpected, err := q.HeapExtractMax()
	require.NoError(t, err)
	require.Equal(t, maxExpected, 21)
}