package metrics

import "github.com/shirou/gopsutil/v3/mem"

func GetRamMemory() (float64, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}
	
	return v.UsedPercent, nil
}