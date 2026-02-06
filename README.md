# Remote System Pulse Agent

A lightweight, high-performance system monitoring agent written in **Go**. 

The **Pulse Agent** is designed to run as a background service on host machines. It gathers real-time hardware telemetry (CPU, RAM, and Disk usage) and dispatches the data to a centralized API for monitoring and analysis.

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
├── main.go                 # Entry point & Configuration loader
├── internal/
│   ├── metrics/            # Hardware data collection logic
│   │   ├── cpu.go
│   │   ├── memory.go
│   │   └── disk.go
│   ├── scheduler/          # Orchestration and timing
│   │   └── runner.go
│   └── sender/             # HTTP transport layer
│       └── client.go
├── go.mod
└── .gitignore