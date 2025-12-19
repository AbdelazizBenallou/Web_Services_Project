This README explains **architecture, services, messaging, APIs, environment variables, and how to run everything step-by-step** based strictly on your codebase .

---

# Black Project – Microservices E-Commerce Backend

A **Go-based microservices backend** implementing a simple e-commerce workflow using **Clean Architecture**, **PostgreSQL**, and **RabbitMQ** for **event-driven communication**.

---

## 📌 Architecture Overview

This project is composed of **three independent microservices**:

| Service             | Responsibility                           | Port   |
| ------------------- | ---------------------------------------- | ------ |
| **User Service**    | User registration, authentication, roles | `8080` |
| **Order Service**   | Order creation & lifecycle               | `8081` |
| **Product Service** | Products, categories & inventory         | `8082` |

### Communication

* **HTTP REST** for synchronous operations
* **RabbitMQ (Topic Exchange)** for async event handling

### Key Events

* `user.registered`
* `order.created`
* `inventory.reserved`
* `inventory.failed`

---

## 🧱 Project Structure

```
black_project/
├── services/
│   ├── user_service/
│   ├── order_service/
│   └── product_service/
└── .gitignore
```

Each service follows **Clean Architecture**:

```
service/
├── config/        # DB config
├── delivery/      # HTTP handlers & routes
├── domain/        # Entities & business rules
├── repository/    # Database layer
├── usecase/       # Application logic
├── messaging/     # RabbitMQ producers/consumers
└── main.go        # Service entrypoint
```

---

## 🔁 Event Flow (End-to-End)

1. **User registers**

   * User Service publishes `user.registered`
   * Order Service stores user in `user_view`

2. **Order is created**

   * Order Service publishes `order.created`

3. **Inventory reservation**

   * Product Service consumes `order.created`
   * If stock OK → publishes `inventory.reserved`
   * If stock fails → publishes `inventory.failed`

4. **Order final state**

   * Order Service updates order to:

     * `CONFIRMED` or
     * `CANCELLED`

---

## 🛠 Tech Stack

* **Language**: Go 1.22+
* **Database**: PostgreSQL
* **Message Broker**: RabbitMQ
* **Architecture**: Clean Architecture
* **Security**:

  * Argon2 password hashing
  * Constant-time password comparison

---

## ⚙️ Environment Variables

All services require the following:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=black_project
DB_SSLMODE=disable

RABBITMQ_URL=amqp://guest:guest@localhost:5672/
```

> Each service may use **its own database schema**.

---

## 🗄 Database Tables (Required)

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

> Tables must exist before running services.

---

## ▶️ How to Run the Project

### 1️⃣ Start Dependencies

#### PostgreSQL

```bash
docker run -d \
  --name postgres \
  -e POSTGRES_PASSWORD=postgres \
  -p 5432:5432 \
  postgres
```

#### RabbitMQ

```bash
docker run -d \
  --name rabbitmq \
  -p 5672:5672 \
  -p 15672:15672 \
  rabbitmq:3-management
```

RabbitMQ UI → [http://localhost:15672](http://localhost:15672)
User: `guest` | Password: `guest`

---

### 2️⃣ Run Services (in separate terminals)

#### User Service

```bash
cd services/user_service
go run main.go
```

#### Order Service

```bash
cd services/order_service
go run main.go
```

#### Product Service

```bash
cd services/product_service
go run main.go
```

---

## 🔌 API Endpoints

### 🧑 User Service (`:8080`)

| Method | Endpoint      | Description   |
| ------ | ------------- | ------------- |
| POST   | `/register`   | Register user |
| POST   | `/login`      | Login         |
| GET    | `/users/{id}` | Get user      |
| GET    | `/users`      | List users    |
| GET    | `/health`     | Health check  |

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

Example payload:

```json
{
  "user_id": 1,
  "items": [
    { "product_id": 1, "quantity": 2, "price": 100 }
  ]
}
```

---

## 🔐 Security Notes

* Passwords hashed with **Argon2id**
* No plaintext passwords stored
* Constant-time comparison prevents timing attacks
* Sensitive fields removed from responses

---

## 🚀 Future Improvements

* JWT authentication
* API Gateway
* Docker Compose
* Observability (Prometheus + Grafana)
* Distributed tracing
* Saga pattern enhancements

---

## 📄 License

MIT License

---

If you want, I can also:

* ✅ Create **Docker Compose**
* ✅ Add **Swagger/OpenAPI**
* ✅ Write **SQL migrations**
* ✅ Convert this into **monorepo CI/CD**

Just tell me 👍
