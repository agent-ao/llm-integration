# 🧱 Message Service - Go (Gin) Microservice

A production-ready Go service using the [Gin](https://github.com/gin-gonic/gin) framework, structured with clean architecture principles, and integrated with MongoDB and RabbitMQ.

---

## 📁 Project Architecture

message-service/
├── cmd/ # Entry point (main.go)
│ └── server/
├── internal/ # Internal business logic
│ ├── config/ # Env & config loading
│ ├── handler/ # HTTP handlers
│ ├── model/ # Domain models (DTOs, types)
│ ├── repository/ # MongoDB, RabbitMQ, or other persistence
│ ├── router/ # HTTP router setup
│ └── service/ # Business logic (use cases)
├── pkg/ # Public utilities (e.g., logger)
├── go.mod / go.sum # Go dependencies
└── .env # Environment variables


---

## ⚙️ Tech Stack

| Layer        | Technology                      |
|--------------|----------------------------------|
| HTTP Server  | [Gin](https://github.com/gin-gonic/gin) |
| Database     | MongoDB                         |
| Messaging    | RabbitMQ                        |
| Config       | `.env` via [`godotenv`](https://github.com/joho/godotenv) |
| Logging      | Built-in Go logger              |

---

## 🚀 Development Workflow

### 1. Clone the Repo

```bash
git clone https://github.com/your-org/message-service.git
cd message-service

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Set Up Environment Variables
``` bash
PORT=8081
MONGO_URI=mongodb://localhost:27017/message-service
RABBITMQ_URI=amqp://guest:guest@localhost:5672/
```

### 4. Run the Service with hot-reloading

We use air to automatically reload the Go server when you change files. This drastically improves developer experience.

Install air if you haven't already:
```bash
go install github.com/cosmtrek/air@latest
```

Then run the service:
```bash
air
```
air will:

Watch .go files inside the project

Automatically recompile and restart the Gin server

Load environment variables from .env before running

🐳 Option: Run with Docker (dev mode)
If you're using Docker Compose for development:

Ensure docker-compose.override.yaml is present with volume mounts

Use this command:
```bash
docker-compose up --build
```
