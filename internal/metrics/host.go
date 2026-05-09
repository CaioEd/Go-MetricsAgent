package metrics

import "github.com/shirou/gopsutil/v3/host"

func GetOperatingSystemAndKernel() (string, string, error) {
	h, err := host.Info()
	if err != nil {
		return "", "", err
	}

	return h.OS, h.KernelVersion, nil
}

func GetOperationSystem() (string, error) {
	osName, _, err := GetOperatingSystemAndKernel()
	if err != nil {
		return "", err
	}

	return osName, nil
}

func GetHostName() (string, error) {
	h, err := host.Info()
	if err != nil {
		return "", err
	}
	return h.Hostname, nil
}

func GetKernelVersion() (string, error) {
	_, kernel, err := GetOperatingSystemAndKernel()
	if err != nil {
		return "", err
	}

	return kernel, nil
}
