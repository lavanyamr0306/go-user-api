

# **User Management Backend (Go)**

This project is a backend service built using **Go** to manage user information.  
The user’s **age** is calculated dynamically from the date of birth and is **not stored** in the database.

The project focuses on **clean architecture**, **layered design**, and **production-ready backend practices**.

---

## **Project Overview**

The application provides REST APIs to **create, read, update, delete, and list users**.  
Each user has a **name** and **date of birth** stored in a **MySQL database**.  
**Age** is computed at runtime using Go’s `time` package to ensure accuracy.

---

## **Technologies Used**

- **Go (Fiber)**  
- **MySQL**  
- **SQLC**  
- **go-playground/validator**  
- **Uber Zap**  
- **Docker & Docker Compose**  

---

## **Project Structure**



/cmd/server/main.go
/config/
/db/migrations/
/db/sqlc/
/internal/
handler/
repository/
service/
routes/
middleware/
models/
logger/


This structure keeps the code **modular, maintainable, and readable**.

---

## **Database Design**

The application uses a single `users` table:

- **id** – primary key  
- **name** – user name  
- **dob** – date of birth  

The **age** is **not stored** and is calculated dynamically when retrieving user data.

---

## **API Endpoints**

### **Create User**

**POST** `/users`

Request example:



{
"name": "Alice",
"dob": "1990-05-10"
}


---

### **Get User by ID**

**GET** `/users/:id`  

Returns user details along with **calculated age**.

---

### **Update User**

**PUT** `/users/:id`

---

### **Delete User**

**DELETE** `/users/:id`  

Returns **HTTP 204** on success.

---

### **List Users (Pagination Supported)**

**GET** `/users?page=1&limit=10`  

Supports **pagination** using `page` and `limit` query parameters.

Example:



/users?page=1&limit=5


---

## **Middleware**

Custom middleware includes:

- **Request ID generation** (`X-Request-ID` header)  
- **Request duration logging**  
- **Centralized validation and error handling**  

---

## **Validation and Logging**

- Requests are validated using **go-playground/validator**  
- Logs are written using **Uber Zap** in a structured format  

---

## **Running the Project Locally**

### Requirements

- **Go 1.21 or higher**  
- **MySQL**  
- **SQLC**  

### Steps



go mod tidy
go run cmd/server/main.go


The server starts on **port 8080**.

---

## **Running with Docker**



docker-compose up --build


This starts both the **API** and **MySQL database**.

---

## **Testing**

Unit tests are included for the **age calculation logic**.

Run tests using:



go test ./...



## **Author**

**Lavanya M R**  
GitHub: [https://github.com/lavanyamr0306](https://github.com/lavanyamr0306)
