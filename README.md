# Echo + Machinery + RabbitMQ (Golang)

This project demonstrates a simple architecture using:

- 🐹 [Echo](https://echo.labstack.com/) — Fast and minimalist web framework
- ⚙️ [Machinery](https://github.com/RichardKnop/machinery) — Asynchronous task queue/job queue
- 🐇 [RabbitMQ](https://www.rabbitmq.com/) — Message broker

It provides:

- API server (`/add`, `/multiply`)
- Background workers for handling tasks
- RabbitMQ as broker and result backend

---

## 🐳 Prerequisites

### Install Dependencies

```bash
go mod tidy
```

### Start RabbitMQ via Docker

```bash
docker run -d --name rabbitmq-test \
  -p 5672:5672 \
  -p 15672:15672 \
  rabbitmq:3-management
```

Then visit [http://localhost:15672](http://localhost:15672/) to access the RabbitMQ management UI (user: `guest`, pass: `guest`).

---

## ✨ Usage

### Run the API server

```bash
go run . api
```

> This will start an Echo server on port `:8080` and send task to worker.
> Example :
>
> - `GET /add?a=2&b=3`
> - `GET /multiply?a=2&b=3`

### Run the Producer (optional load testing / spawner)

```bash
go run . producer
```

> This will hit `/add` and `/multiply` endpoints concurrently with random integer values.

### Run the Worker

```bash
go run . worker
```

> This will listen to RabbitMQ queues and execute registered tasks.

---

## 💬 Notes

- This repo only demonstrate 1 worker + 1 queue with multiple tasks, but it can be configured using multiple worker for each tasks.
- Each worker can consume tasks from the same or different queues (via `RoutingKey`).
