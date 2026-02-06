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

## 🧪 Manual Walkthrough (Live Demo)

You can demonstrate the full "Meter-to-Bill" lifecycle using these commands:

### 1. 📥 Record Usage (Hits Redis)
Send a metric event for a customer. This is processed instantly via Redis.
```powershell
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/usage" -Method Post -Body '{"customer_id": "YOUR_CUSTOMER_ID", "resource_type": "api_call", "quantity": 100}' -ContentType "application/json"
```

### 2. 💰 Check Estimated Bill (Real-time Calculation)
Calculate the bill by combining Redis usage with Postgres pricing metadata.
```powershell
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/billing/YOUR_CUSTOMER_ID" -Method Get
```

### ⚙️ 3. Background Aggregation (Automatic)
Watch the server logs. Every 60 seconds, the **Aggregator Worker** will:
1. Scan Redis for active meters.
2. Flush the totals to the `usage_records` table in PostgreSQL.
3. Reset Redis meters to zero.

---

## 🏗️ Technical Highlights
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

