package metrics

import (
	"time"
	"github.com/shirou/gopsutil/v3/cpu"
)

func GetCpuUsage() (float64, error) {
	// Pega a média de todos os cores (false) num intervalo de 1s
	percentages, err := cpu.Percent(1*time.Second, false)
	if err != nil {
		return 0, err
	}
	if len(percentages) > 0 {
		return percentages[0], nil
	}
	return 0, nil
}