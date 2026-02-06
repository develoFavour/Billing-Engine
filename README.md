# 💰 Billing & Metering Engine (Golang)

A high-performance, scalable billing microservice built with Golang. This service is designed to turn real-time usage metrics into billable amounts, providing a standardized capability for enterprise-grade billing cycles.

## 🚀 High-Performance Architecture
This engine is built for scale, using a **Dual-Persistence Strategy**:
- **Hot Path (Ingestion)**: Usage metrics are atomically incremented in **Redis** via `INCRBYFLOAT`. This allows for sub-millisecond response times and handles high-concurrency metric ingestion without overloading the primary database.
- **Cold Path (Persistence)**: A background **Aggregator Worker** (running in a dedicated goroutine) periodically flushes totals from Redis to **PostgreSQL**. This ensures long-term reliability and data integrity while keeping the "Hot Path" fast.

## 🏗️ Technical Highlights
- **Distributed Metering**: Real-time usage tracking with atomic Redis operations.
- **Worker Pattern**: Decoupled background processing using Go's concurrency primitives (`goroutines`, `channels`, `tickers`).
- **Repository Pattern**: Clean abstraction of data layers (Postgres / Redis / Mock).
- **Graceful Shutdown**: Handles OS signals to ensure no data loss during worker flushes or server restarts.
- **Scalable Schema**: UUID-based multi-tenant design with support for flexible pricing tiers.

## 🔌 Tech Stack
- **Language**: Golang 1.21+
- **API Framework**: Gin Gonic
- **Primary DB**: PostgreSQL (pgx/v5)
- **Caching/Metering**: Redis (go-redis/v9)
- **Infrastructure**: Docker & Docker Compose
- **Concurrency**: Goroutines, Tickers, Context Management

## 🛠️ Getting Started

### Prerequisites
- Go 1.21+
- Docker & Docker Compose
- Make (optional)

### Setup
1. Clone the repository.
2. Setup environment variables:
   ```bash
   cp .env.example .env
   ```
3. Start the infrastructure:
   ```bash
   make docker-up
   ```
4. Run the application:
   ```bash
   make run
   ```

## 🔌 API Endpoints
- `GET /health`: Basic health check.

## 🧪 Testing
```bash
make test
```

---
*Created by Favour Opia as part of a technical portfolio for the Canonical Golang Software Engineer position.*
