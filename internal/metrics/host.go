package metrics

import "github.com/shirou/gopsutil/v3/host"

func GetOperationSystem() (string, error) {
	h, err := host.Info()
	if err != nil {
		return "", err
	}
	return h.OS, nil
}

func GetHostName() (string, error) {
	h, err := host.Info()
	if err != nil {
		return "", err
	}
	return h.Hostname, nil
}

func GetKernelVersion() (string, error) {
	h, err := host.Info()
	if err != nil {
		return "", err
	}
	return h.KernelVersion, nil
}