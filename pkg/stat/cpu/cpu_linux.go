// go:build linux

package cpu

func Load() int {
	return 5
}

func GetStat() *Stat {
	return &Stat{
		Load: Load(),
	}
}
