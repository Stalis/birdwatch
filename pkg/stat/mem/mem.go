package mem

import "context"

type MemoryStat struct {
	Available int
	Total     int
	Used      int
}

func GetMemoryStat(ctx context.Context) *MemoryStat {
	res := &MemoryStat{}
	vmStat := GetDarwinVMStat(ctx)

	res.Available = vmStat.PagesFree + vmStat.PagesInactive
	res.Total = vmStat.Total
	res.Used = vmStat.Total - res.Available

	return res
}
