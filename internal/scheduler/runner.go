package scheduler

import (
	"log"
	"monitor-agent/internal/metrics"
	"monitor-agent/internal/sender"
	"time"
)

type Config struct {
	Interval time.Duration
	ApiUrl   string
	Token    string
}

func Start(cfg Config) {
	ticker := time.NewTicker(cfg.Interval)
	defer ticker.Stop()

	log.Printf("Iniciando monitoramento. Envio a cada %v para %s", cfg.Interval, cfg.ApiUrl)

	// Loop infinito controlado pelo Ticker
	for range ticker.C {
		collectAndSend(cfg)
	}
}

func collectAndSend(cfg Config) {
	// 1. Coleta (sequencial ou goroutines se fosse crítico, aqui sequencial basta)
	cpuUse, err := metrics.GetCpuUsage()
	if err != nil {
		log.Printf("Erro ao ler CPU: %v", err)
		return
	}

	ramUse, err := metrics.GetRamMemory()
	if err != nil {
		log.Printf("Erro ao ler RAM: %v", err)
		return
	}

	diskUse, err := metrics.GetDiskUsage()
	if err != nil {
		log.Printf("Erro ao ler Disco: %v", err)
		return
	}

	// 2. Monta o Payload
	payload := sender.Payload{
		Token:     cfg.Token,
		UsageCPU:  cpuUse,
		UsageMemory:  ramUse,
		UsageDisk: diskUse,
	}

	// 3. Envia
	log.Printf("Enviando métricas: CPU: %.2f%%, RAM: %.2f%%, DISK: %.2f%%", cpuUse, ramUse, diskUse)
	if err := sender.SendMetrics(cfg.ApiUrl, payload); err != nil {
		log.Printf("Falha ao enviar dados para API: %v", err)
	} else {
		log.Println("Dados enviados com sucesso!")
	}
}