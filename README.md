```markdown
# Go - Backend Development Task

# 🧪 User with DOB and Calculated Age

## 📌 Objective

Build a RESTful API using Go to manage users with their `name` and `dob` (date of birth). The API should calculate and return a user’s age dynamically when fetching user details.

---

## 🗂️ Project Structure

Please follow the directory structure below:

```

/cmd/server/main.go
/config/
/db/migrations/
/db/sqlc/<generated>
/internal/
├── handler/
├── repository/
├── service/
├── routes/
├── middleware/
├── models/
└── logger/

````

---

## 🔧 Tech Stack

- [GoFiber](https://gofiber.io/)
- SQL (MySQL or PostgreSQL) + [SQLC](https://sqlc.dev/)
- [Uber Zap](https://github.com/uber-go/zap) for logging
- [go-playground/validator](https://github.com/go-playground/validator) for input validation

---

## 📊 Required Table

`users` table schema:

| Field | Type | Constraints |
| --- | --- | --- |
| id | SERIAL | PRIMARY KEY |
| name | TEXT | NOT NULL |
| dob | DATE | NOT NULL |

---

## 🔄 API Endpoints

### Create User

**POST** `/users`

**Request:**

```json
{
  "name": "Alice",
  "dob": "1990-05-10"
}
````

**Response:**

```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10"
}
```

---

### Get User by ID

**GET** `/users/:id`

**Response:**

```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 35
}
```

---

### Update User

**PUT** `/users/:id`

**Request:**

```json
{
  "name": "Alice Updated",
  "dob": "1991-03-15"
}
```

**Response:**

```json
{
  "id": 1,
  "name": "Alice Updated",
  "dob": "1991-03-15"
}
```

---

### Delete User

**DELETE** `/users/:id`

**Response:**

* HTTP `204 No Content`

---

### List All Users

**GET** `/users`

**Response:**

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

## ✅ Requirements

* Store `dob` in the database
* Return `age` calculated dynamically using Go's `time` package
* Use SQLC for generating DB access layer
* Validate inputs with `go-playground/validator`
* Log key actions using Uber Zap
* Clean HTTP status codes and error handling

---

