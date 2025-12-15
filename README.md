
# User Management Backend (Go)

This project is a backend service developed in Go to manage user data, where a user's age is calculated dynamically from the date of birth instead of being stored in the database.

The design emphasizes a clean project structure, clear separation of concerns, and reliable handling of data.

---

## Project Overview

The application stores user information, including date of birth, in a MySQL database.  
Instead of storing age, it is computed at runtime whenever user data is retrieved, ensuring it is always accurate.

---

## Technologies Used

- **Go (Fiber)** – backend framework  
- **MySQL** – relational database  
- **SQLC** – type-safe database access  
- **go-playground/validator** – request validation  
- **Uber Zap** – structured logging  

---

## Project Structure

```

/cmd/server/main.go
/config/
/db/migrations/
/db/sqlc/
/internal/
├── handler/
├── repository/
├── service/
├── routes/
├── middleware/
├── models/
└── logger/

````

The structure ensures responsibilities are clearly separated and the codebase is easy to maintain.

---

## Database Design

The application uses a single `users` table with the following fields:

- `id` – primary key  
- `name` – user name  
- `dob` – date of birth  

The user's **age is not stored** in the database; it is calculated dynamically using Go’s `time` package whenever needed.

---

## Setup Summary

To run this project:

- Install **Go** (version 1.21 or higher)  
- Install and configure **MySQL**  
- Set the required environment variables  
- Apply database schema using migration files  

**Key Features:**

- Dynamic age calculation without storing redundant data  
- Type-safe database interactions with SQLC  
- Input validation at the API boundary  
- Structured logging using Uber Zap  
- Clear separation between HTTP, service, and data layers  

---

## What This Project Demonstrates

- Backend development with Go and Fiber  
- Integration with a MySQL database  
- Maintainable and clean project architecture  
- Practical backend best practices  

API functionality has been tested locally using Postman.  
The server runs locally and handles all user-related operations efficiently.

---

## Running the Application

After completing setup, start the server with:

```bash
go run cmd/server/main.go
````

```


```
