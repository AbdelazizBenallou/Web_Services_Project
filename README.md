---

# Black Project – Go Microservices E-Commerce Backend

A **Go-based microservices backend** implementing an event-driven e-commerce system using **Clean Architecture**, **PostgreSQL**, **RabbitMQ**, and **Docker Compose**.

This project demonstrates **service decoupling**, **asynchronous messaging**, and **domain-driven design principles**.

---

## 📌 High-Level Overview

The system is composed of **three independent microservices**:

| Service             | Responsibility                           | Port   |
| ------------------- | ---------------------------------------- | ------ |
| **User Service**    | User registration, authentication, roles | `8080` |
| **Order Service**   | Order creation & lifecycle management    | `8081` |
| **Product Service** | Products, categories & inventory         | `8082` |

### Communication Patterns

* **HTTP REST APIs** for synchronous operations
* **RabbitMQ (topic exchange)** for asynchronous event communication

---

## 🧱 Architecture

Each service follows **Clean Architecture**:

```
service/
├── config/        # Database configuration
├── delivery/      # HTTP handlers & routes
├── domain/        # Entities & business rules
├── repository/    # Persistence layer
├── usecase/       # Application logic
├── messaging/     # RabbitMQ producers & consumers
└── main.go        # Service entrypoint
```

### Key Design Principles

* Dependency inversion
* Explicit domain rules
* Event-driven communication
* No shared databases between services

---

## 🔁 Event Flow

### 1️⃣ User Registration

* User registers via **User Service**
* `user.registered` event is published
* Order Service consumes the event and stores a read-only user view

### 2️⃣ Order Creation

* Order Service validates user existence
* Order is created with status `PENDING_INVENTORY`
* `order.created` event is published

### 3️⃣ Inventory Reservation

* Product Service consumes `order.created`
* Attempts to reserve stock

  * Success → publishes `inventory.reserved`
  * Failure → publishes `inventory.failed`

### 4️⃣ Order Finalization

* Order Service updates order status:

  * `CONFIRMED` on success
  * `CANCELLED` on failure

---

## 🛠 Tech Stack

* **Language**: Go (1.22+)
* **Database**: PostgreSQL
* **Message Broker**: RabbitMQ
* **Containerization**: Docker & Docker Compose
* **Security**:

  * Argon2id password hashing
  * Constant-time password comparison

---

## ▶️ How to Run (Docker Compose)

### ✅ Prerequisites

* Docker (v20+)
* Docker Compose (v2+)

Verify installation:

```bash
docker --version
docker compose version
```

---

### 🚀 Start the Application

From the **project root** (where `docker-compose.yml` is located):

```bash
docker compose up --build
```

This command will:

* Build all Go services
* Start PostgreSQL
* Start RabbitMQ
* Start all microservices
* Create a shared Docker network

---

### 🧯 Stop the Application

```bash
docker compose down
```

Remove volumes (database reset):

```bash
docker compose down -v
```

---

## 🌐 Exposed Ports

| Service         | Port    |
| --------------- | ------- |
| User Service    | `8080`  |
| Order Service   | `8081`  |
| Product Service | `8082`  |
| RabbitMQ UI     | `15672` |
| PostgreSQL      | `5432`  |

RabbitMQ Management UI:
👉 [http://localhost:15672](http://localhost:15672)
**Username:** `guest`
**Password:** `guest`

---

## 🔧 Environment Configuration

All environment variables are defined inside `docker-compose.yml`.

Each service uses:

```env
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=black_project
DB_SSLMODE=disable

RABBITMQ_URL=amqp://guest:guest@rabbitmq:5672/
```

> ⚠️ `postgres` and `rabbitmq` are **Docker service names**, not `localhost`.

---

## 🗄 Database Tables

### User Service

* `users`
* `profiles`
* `roles`

### Order Service

* `orders`
* `order_items`
* `user_view`

### Product Service

* `categories`
* `products`
* `stock`

> Tables must be created via migrations or init scripts before production use.

---

## 🔌 API Endpoints

### 🧑 User Service (`:8080`)

| Method | Endpoint      | Description     |
| ------ | ------------- | --------------- |
| POST   | `/register`   | Register a user |
| POST   | `/login`      | User login      |
| GET    | `/users/{id}` | Get user by ID  |
| GET    | `/users`      | List all users  |
| GET    | `/health`     | Health check    |

---

### 📦 Product Service (`:8082`)

| Method | Endpoint                    | Description          |
| ------ | --------------------------- | -------------------- |
| POST   | `/categories`               | Create category      |
| GET    | `/categories`               | List categories      |
| POST   | `/products`                 | Create product       |
| GET    | `/products`                 | List products        |
| GET    | `/categories/{id}/products` | Products by category |

---

### 🧾 Order Service (`:8081`)

| Method | Endpoint  | Description  |
| ------ | --------- | ------------ |
| POST   | `/orders` | Create order |

**Example Request**

```json
{
  "user_id": 1,
  "items": [
    {
      "product_id": 1,
      "quantity": 2,
      "price": 100
    }
  ]
}
```

---

## 🔐 Security Notes

* Passwords are hashed using **Argon2id**
* No plaintext passwords are stored or returned
* Constant-time comparison prevents timing attacks
* Sensitive fields removed from API responses

---

## 🧪 Useful Docker Commands

```bash
docker compose ps
docker compose logs -f
docker compose restart order_service
```

---

## 🚀 Future Improvements

* JWT authentication & authorization
* API Gateway
* Database migrations
* Observability (Prometheus, Grafana)
* Distributed tracing
* Saga pattern enhancements
* Kubernetes deployment

---

## 📄 License

MIT License

---

