//go:build darwin

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
				"Pages reactivated:                      1710876.",
			},
			"Pages reactivated",
			1710876,
		},
		{
			"Test with trim doube quotes",
			args{
				"\"Translation faults\":                  39725408.",
			},
			"Translation faults",
			39725408,
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

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			gotKey, gotValue := parseLine(testCase.args.input)

			require.Equal(t, testCase.wantKey, gotKey)
			require.Equal(t, testCase.wantValue, gotValue)
		})
	}
}
