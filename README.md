

---

```md
# go-user-api

A simple backend REST API built in Go using GoFiber.  
This project manages users with name and date of birth and calculates the user’s age dynamically at runtime.

The project focuses on clean backend structure, type-safe database access, logging, middleware, Docker support, and basic unit testing.

## What this project does

- Create, read, update, and delete users
- Store user name and date of birth in the database
- Calculate age dynamically using Go’s time package
- Return clean JSON responses
- Follow a layered backend architecture

## Tech used

- Go
- GoFiber
- MySQL
- SQLC
- Uber Zap
- go-playground/validator
- Docker and Docker Compose

## Project structure

```

cmd/server/main.go
config/
db/
internal/
├── handler/
├── service/
├── repository/
├── routes/
├── middleware/
├── models/
└── logger/

````

## Database

Table name: users

Columns:
- id (primary key)
- name (text, not null)
- dob (date, not null)

Age is not stored in the database.  
It is calculated when the user data is fetched.

## API endpoints

### Create user

POST /users

```json
{
  "name": "Alice",
  "dob": "1990-05-10"
}



````

### Get user by id

GET /users/:id

Returns user details along with calculated age.

### Update user

PUT /users/:id

### Delete user

DELETE /users/:id

Returns HTTP 204 on success.

### List users (pagination supported)

GET /users?page=1&limit=10

## Pagination

Pagination is implemented for the list users API using page and limit query parameters.

Example:

```
/users?page=1&limit=5
```

## Middleware

The project uses custom middleware for:

* Request ID generation (added as X-Request-ID header)
* Logging request method, path, and execution time

## Validation and logging

* Request payloads are validated using go-playground/validator
* Logs are written using Uber Zap in a structured format

## Running the project locally

Requirements:

* Go 1.21 or higher
* MySQL
* SQLC

Steps:

```bash
go mod tidy
go run cmd/server/main.go
```

The server will start on port 8080.

## Running with Docker

```bash
docker-compose up --build
```

This starts both the API and the MySQL database.

## Tests

Unit tests are added for the age calculation logic.

```bash
go test ./...
```



## Author

Lavanya M R
GitHub: [https://github.com/lavanyamr0306](https://github.com/lavanyamr0306)

````

---



