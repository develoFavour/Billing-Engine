# 💰 Billing & Metering Engine (Golang)

A high-performance, scalable billing microservice built with Golang. This service is designed to turn real-time usage metrics into billable amounts, providing a standardized capability for enterprise-grade billing cycles.

## 🚀 Built With
- **Language**: Golang 1.21+
- **API Framework**: Gin Gonic
- **Persistence**: PostgreSQL (pgx/v5 for high performance)
- **Caching**: Redis (Real-time metering cache)
- **Infrastructure**: Docker & Docker Compose
- **Quality**: Graceful shutdown, Repository Pattern, Dependency Injection

## 🏗️ Architecture
The project follows the standard Go project layout:
- `cmd/`: Entry points for the application.
- `internal/`: Private application and library code.
- `pkg/`: Public library code that can be used by other projects.
- `migrations/`: SQL migration files.

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
