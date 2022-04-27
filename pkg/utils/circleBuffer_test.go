package utils_test

import (
	"testing"

	"github.com/Stalis/birdwatch/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestCircleBuffer_Add(t *testing.T) {
	t.Run("add first item", func(t *testing.T) {
		buf := utils.NewCircleBuffer(2)

		testVal := 0
		buf.Add(testVal)

		require.Equal(t, testVal, buf.Get(0))
	})
}

func TestCircleBuffer_Closed(t *testing.T) {
	t.Run("close buffer with cap = 2", func(t *testing.T) {
		buf := utils.NewCircleBuffer(2)

		testVal1 := 0
		testVal2 := 1
		testVal3 := 3

		buf.Add(testVal1)
		buf.Add(testVal2)

		require.True(t, buf.Closed())

		buf.Add(testVal3)
		require.True(t, buf.Closed())
		require.Equal(t, testVal3, buf.Get(0))
	})
}
