package mem

type MemoryStat struct {
	Available int
	Total     int
	Used      int

	LinuxVMStat  *LinuxVMStat
	DarwinVMStat *DarwinVMStat
}

type DarwinVMStat struct {
	Total          int
	PagesFree      int
	PagesActive    int
	PagesInactive  int
	PagesWiredDown int
}

type LinuxVMStat struct {
	MemTotal     int
	MemFree      int
	MemAvailable int
}
