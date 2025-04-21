package pagefilter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_LimitDefault(t *testing.T) {
	v := limitDefault()
	require.Equal(t, 100, v)
}

func Test_LimitMin(t *testing.T) {
	v := limitMin()
	require.Equal(t, 1, v)
}

func Test_LimitMax(t *testing.T) {
	v := limitMax()
	require.Equal(t, 20000, v)
}
