package main

import (
	"log"
	"monitor-agent/internal/scheduler"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// Carrega as variáveis de ambiente
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado")
		return
	}

	apiUrl := os.Getenv("API_URL")
	if apiUrl == "" {
		log.Fatal("API_URL não definida")
	}

	token := os.Getenv("AGENT_TOKEN")
	if token == "" {
		log.Fatal("Token de agente não definido")
	}

	intervalSec := time.Duration(10) * time.Second	
	
	cfg := scheduler.Config{
		Interval: intervalSec,
		ApiUrl:   apiUrl,
		Token:    token,
	}

	log.Printf("Pulse Agent iniciado com sucesso!")
    log.Printf("Monitorando a cada %v e enviando para: %s", intervalSec, apiUrl)

	// Inicia o agente
	scheduler.Start(cfg)
}