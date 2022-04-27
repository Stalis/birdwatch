package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCircleBuffer_AddInternal(t *testing.T) {
	t.Run("tirst item added", func(t *testing.T) {
		buf := NewCircleBuffer(2)

		testVal := 0
		buf.Add(testVal)

		require.Equal(t, testVal, buf.items[0])
	})

	t.Run("two item added with cap 1", func(t *testing.T) {
		buf := NewCircleBuffer(1)

		testVal1 := 0
		buf.Add(testVal1)

		testVal2 := 1
		buf.Add(testVal2)

		require.Equal(t, testVal2, buf.items[0])
	})
}

func TestCircleBuffer_GetInternal(t *testing.T) {
	t.Run("get result equal slice item", func(t *testing.T) {
		buf := NewCircleBuffer(1)

		buf.Add(4)

		require.Equal(t, buf.items[0], buf.Get(0))
	})
}
