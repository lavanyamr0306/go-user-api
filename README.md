

```md
# Go User API 🚀

A RESTful User Management API built using **GoFiber**, **SQLC**, and **MySQL**, which stores user details (`name`, `dob`) and dynamically calculates the **age** using Go’s `time` package.

This project follows **clean architecture**, includes **Docker support**, **pagination**, **middleware**, and **unit tests**, making it production-ready and interview-friendly.

---

## 📌 Features

- CRUD operations for Users
- Dynamic age calculation (not stored in DB)
- Clean layered architecture (handler, service, repository)
- SQLC for type-safe database access
- Input validation using go-playground/validator
- Structured logging using Uber Zap
- Pagination support for listing users
- Docker & Docker Compose support
- Middleware for:
  - Request ID injection
  - Request duration logging
- Unit tests for age calculation

---

## 🗂️ Project Structure

```

go-user-api/
├── cmd/server/main.go
├── config/
├── db/
│   ├── migrations/
│   └── sqlc/
├── internal/
│   ├── handler/
│   ├── service/
│   ├── repository/
│   ├── routes/
│   ├── middleware/
│   ├── models/
│   └── logger/
├── docker-compose.yaml
├── Dockerfile
├── sqlc.yaml
├── go.mod
└── go.sum

````

---

## 🔧 Tech Stack

- **Language:** Go
- **Framework:** GoFiber
- **Database:** MySQL
- **ORM:** SQLC
- **Logging:** Uber Zap
- **Validation:** go-playground/validator
- **Containerization:** Docker & Docker Compose
- **Testing:** Go testing package

---

## 🗃️ Database Schema

### `users` table

| Field | Type | Constraints |
|------|------|-------------|
| id | SERIAL | Primary Key |
| name | TEXT | NOT NULL |
| dob | DATE | NOT NULL |

---

## 🔄 API Endpoints

### ➕ Create User
**POST** `/users`

```json
{
  "name": "Alice",
  "dob": "1990-05-10"
}
````

---

### 📄 Get User by ID

**GET** `/users/:id`

```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 35
}
```

---

### ✏️ Update User

**PUT** `/users/:id`

```json
{
  "name": "Alice Updated",
  "dob": "1991-03-15"
}
```

---

### ❌ Delete User

**DELETE** `/users/:id`

**Response:** `204 No Content`

---

### 📃 List Users (Pagination)

**GET** `/users?page=1&limit=10`

```json
[
  {
    "id": 1,
    "name": "Alice",
    "dob": "1990-05-10",
    "age": 34
  }
]
```

---

## 🧮 Age Calculation Logic

* Age is **calculated dynamically**
* Uses Go’s `time` package
* Age is NOT stored in the database
* Unit tested for correctness

---

## 🧪 Run Unit Tests

```bash
go test ./...
```

---

## 🐳 Docker Setup

### Build & Run using Docker Compose

```bash
docker-compose up --build
```

* API runs on: `http://localhost:8080`
* MySQL runs on: `localhost:3306`

---

## 🛡️ Middleware

### Request ID Middleware

* Adds `X-Request-ID` to every response

### Request Logger Middleware

* Logs:

  * HTTP method
  * Path
  * Request duration
  * Request ID

---

## 📦 SQLC

Generate SQLC code using:

```bash
sqlc generate
```

---

## 🧠 Key Learnings

* Clean architecture in Go
* Type-safe DB access with SQLC
* Dockerizing Go applications
* Writing middleware in Fiber
* Pagination implementation
* Unit testing business logic

---

## 👤 Author

**Lavanya M R**
GitHub: [https://github.com/lavanyamr0306](https://github.com/lavanyamr0306)

---


