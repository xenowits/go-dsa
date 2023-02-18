package heap

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLeftRight(t *testing.T) {
	require.Equal(t, left(2), 5)
	require.Equal(t, right(2), 6)

	require.Equal(t, left(3), 7)
	require.Equal(t, right(3), 8)
}
