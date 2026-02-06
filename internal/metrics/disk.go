package metrics

import "github.com/shirou/gopsutil/v3/disk"

func GetDiskUsage() (float64, error) {
	// "/" pega o root.
	d, err := disk.Usage("/")
	if err != nil {
		return 0, err
	}
	return d.UsedPercent, nil
}