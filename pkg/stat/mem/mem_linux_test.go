//go:build linux

package mem

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_parseLine(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name      string
		args      args
		wantKey   string
		wantValue int
	}{
		{
			"Base test",
			args{
				"MemTotal:       16176776 kB",
			},
			"MemTotal",
			16176776,
		},
		{
			"Empty input",
			args{
				"",
			},
			"",
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, gotValue := parseLine(tt.args.input)

			require.Equal(t, tt.wantKey, gotKey)
			require.Equal(t, tt.wantValue, gotValue)
		})
	}
}
