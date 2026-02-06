package sender

import (
	"net/http"
	"time"
	"encoding/json"
	"bytes"
	"fmt"
)

type Payload struct {
	Token string `json:"token"`
	UsageCPU float64 `json:"usageCpu"`
	UsageMemory float64 `json:"usageRam"`
	UsageDisk float64 `json:"usageDisk"`
}

func SendMetrics(apiUrl string, data Payload) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonData))

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error sending metrics: %s", resp.Status)
	}

	return nil
}