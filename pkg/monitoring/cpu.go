package monitoring

import "math/rand"

func CPULoad() int {
	return rand.Intn(100) //nolint:gosec
}
