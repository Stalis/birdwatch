package mem

import (
	"context"
	"testing"
	"time"

	"github.com/Stalis/birdwatch/pkg/utils"
	"github.com/stretchr/testify/require"
)

func getCircleBufferWithItems(capacity int, items ...interface{}) *utils.CircleBuffer {
	if capacity < 0 {
		capacity = len(items)
	}

	res := utils.NewCircleBuffer(capacity)
	for _, v := range items {
		res.Add(v)
	}

	return res
}

func TestWatcher_Avg(t *testing.T) {
	type fields struct {
		buffer            *utils.CircleBuffer
		averagingInterval time.Duration
		scanInterval      time.Duration
		cancelFunc        context.CancelFunc
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *MemoryStat
	}{
		{
			"Base using",
			fields{
				buffer: getCircleBufferWithItems(-1, &MemoryStat{
					Available: 100,
					Used:      50,
					Total:     25,
				}, &MemoryStat{
					Available: 200,
					Used:      100,
					Total:     50,
				}),
			},
			args{
				context.Background(),
			},
			&MemoryStat{
				Available: 150,
				Used:      75,
				Total:     37,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Watcher{
				buffer:            tt.fields.buffer,
				averagingInterval: tt.fields.averagingInterval,
				scanInterval:      tt.fields.scanInterval,
				cancelFunc:        tt.fields.cancelFunc,
			}

			got := w.Avg(tt.args.ctx)
			require.Equal(t, tt.want, got)
		})
	}
}
