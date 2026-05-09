# Remote System Pulse Agent

A lightweight, high-performance system monitoring agent written in **Go**. 

The **Pulse Agent** is designed to run as a background service on host machines. It gathers real-time hardware telemetry (CPU, RAM, and Disk usage) and dispatches the data to a centralized API for monitoring and analysis.

---

## 🔗 Project Ecosystem
This repository contains the **Monitoring Agent**. To see the full project in action, check out the other components:

* **Backend API:** [System Pulse (Spring Boot)](https://github.com/CaioEd/system-pulse) - Manages states, connections and CRUD of servers.
* **Frontend Dashboard:** [ServersHealth-Next](https://github.com/CaioEd/ServersHealth-Next) (Next.js) - Frontend dashboard to visualize the data in real-time.
---

## 🚀 Features

* **Low Footprint:** Built with Go for minimal CPU and RAM overhead.
* **Real-time Telemetry:** Collects precise metrics using `gopsutil`.
* **Automated Scheduling:** Configurable check intervals (currently set to 10s).
* **Docker Ready:** Easily deployable across distributed environments.
* **Secure:** Supports token-based authentication for API ingestion.

---

## 🛠 Tech Stack

* **Language:** [Go](https://go.dev/) (Golang)
* **Metrics Engine:** [gopsutil](https://github.com/shirou/gopsutil)
* **Communication:** JSON over HTTP (POST)

---

## 📂 Project Structure

Following a clean and modular architecture:

```text
.
├── main.go                      # Entry point & Configuration loader
├── internal/
│   ├── metrics/                 # Hardware data collection logic
│   │   ├── cpu.go
│   │   ├── memory.go
│   │   └── disk.go
│   ├── scheduler/               # Orchestration and timing
│   │   └── runner.go
│   └── sender/                  # HTTP transport layer
│       └── client.go
├── Dockerfile.simulated-server  # Lightweight Linux image for simulation
├── docker-compose.yml           # Local simulated server environment
├── go.mod
└── .gitignore

```

---

## How to run locally

```bash
1. Download dependencies:  go mod tidy
2. Run the agent:          go run main.go
```

---

## 🐳 Simulated production server with Docker Compose

Use the compose stack to spin up a lightweight Linux server where you can run the agent as if it were in production.

### 1) Start the simulated server

```bash
docker compose up -d --build
```

This starts a container named `metrics-agent-simulated-server` based on Alpine Linux + Go.

### 2) Configure the agent environment

Inside the project root, create a `.env` file used by the agent:

```env
API_URL=http://host.docker.internal:8080/metrics
AGENT_TOKEN=your-agent-token
```

> `host.docker.internal` points to services running on your host machine (for example your local backend API).

### 3) Open a shell in the simulated server

```bash
docker compose exec simulated-server bash
```

### 4) Initialize and run the agent inside the container

```bash
go mod tidy
go run main.go
```

The agent will start collecting real-time CPU, memory and disk metrics from the simulated server container and send them to your configured API every 10 seconds.

### 5) (Optional) Generate synthetic load to see metrics changing in real time

In another terminal, run:

```bash
docker compose exec simulated-server sh -lc "stress-ng --cpu 2 --vm 1 --vm-bytes 256M --timeout 120s"
```

### 6) Stop the simulated server

```bash
docker compose down
```
