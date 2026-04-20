# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

**Pulse Agent** — a Go background service that collects CPU / RAM / Disk telemetry via `gopsutil` and POSTs it as JSON every 10s to a central API. It is one piece of a larger ecosystem: a Spring Boot backend (`system-pulse`) and a Next.js dashboard (`ServersHealth-Next`). This repo only contains the agent.

Go module name: `monitor-agent` (note it differs from the repo directory name — internal imports use `monitor-agent/internal/...`).

## Commands

```bash
go mod tidy          # install deps
go run main.go       # run agent (requires .env)
go build             # produce ./monitor-agent binary
```

There is no test suite and no linter configured. `go vet ./...` and `go build ./...` are the only static checks.

### Simulated server (Docker Compose)

`docker-compose.yml` + `Dockerfile.simulated-server` spin up an Alpine+Go container used as a fake production host. It mounts the repo at `/workspace` and defines `host.docker.internal` so the agent can reach a backend running on the host.

```bash
docker compose up -d --build
docker compose exec simulated-server bash       # shell into container
# inside: go mod tidy && go run main.go
docker compose exec simulated-server sh -lc "stress-ng --cpu 2 --vm 1 --vm-bytes 256M --timeout 120s"   # synthetic load
docker compose down
```

Note: compose sets `API_URL` / `AGENT_TOKEN` via `environment:`, while local runs read them from `.env` via `godotenv`. Both must be present for the corresponding mode — `main.go` calls `log.Fatal` if either is missing.

## Architecture

Three-layer pipeline driven by a ticker, all under `internal/`:

- `internal/metrics/` — collection. Each file wraps one `gopsutil/v3` subpackage (`cpu`, `mem`, `disk`, `host`) and returns a scalar or a small tuple. `cpu.GetCpuUsage` blocks for 1s sampling the average across cores.
- `internal/scheduler/` — orchestration. `scheduler.Start(cfg)` owns the `time.Ticker` loop and calls `collectAndSend` each tick. Collection is sequential; a single error aborts that tick (nothing is sent partially).
- `internal/sender/` — transport. `sender.Payload` is the wire format and `SendMetrics` POSTs JSON with a 10s HTTP timeout. Non-200 responses are returned as errors.

`main.go` is just a config loader: reads env vars, builds `scheduler.Config`, and hands off.

### Host info is sent once, not every tick

The scheduler tracks `hostInfoSent` as a local `*bool` in `Start`. `OperatingSystem` / `KernelVersion` are populated on the first successful tick only; on subsequent ticks those fields stay empty and are omitted by the `omitempty` JSON tags. Any change to payload fields that should be sent-once must follow the same pattern (add `omitempty` + flip a flag in `collectAndSend`) — do not move this state into `Config`, which is passed by value.

### Wire contract

`sender.Payload` JSON keys are consumed by the Spring Boot backend — do not rename `usageCpu`, `usageRam`, `usageDisk`, `operatingSystem`, `kernelVersion`, or `token` without a coordinated backend change. `token` is sent inside the JSON body, not as an `Authorization` header.

## Conventions

- Logs and inline comments are in Portuguese (pt-BR). Match the existing style when editing those files; user-facing API field names stay in English.
- The interval is hardcoded to 10s in `main.go` (`intervalSec`). If making it configurable, thread it through env vars rather than flags to stay consistent with the existing config style.
