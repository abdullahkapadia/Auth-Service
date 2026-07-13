# Go Authentication Service

A reusable authentication service built with Go that provides a centralized authentication system for web and mobile applications. Instead of implementing user registration and login in every project, applications can connect to this service through REST APIs.

**Note:** This project is currently configured to run on a local server for development and testing.

## Features

* User registration
* User login
* JWT-based authentication
* Protected profile endpoint
* Password hashing using bcrypt
* PostgreSQL database integration
* RESTful API architecture
* Modular project structure for easy maintenance

## Tech Stack

* Go
* Chi Router
* PostgreSQL
* JWT
* bcrypt
* pgx
* godotenv

## Project Structure

```text
auth-service/
├── cmd/
│   └── server/
├── internal/
│   ├── config/
│   ├── database/
│   ├── handler/
│   ├── middleware/
│   ├── model/
│   ├── repository/
│   ├── router/
│   ├── service/
│   └── utils/
├── .env
├── go.mod
└── go.sum
```

## Getting Started

### Clone the repository

```bash
git clone <repository-url>
cd auth-service
```

### Install dependencies

```bash
go mod tidy
```

### Configure environment variables

Create a `.env` file in the project root.

```env
APP_NAME=Auth Service
APP_PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=auth_db

JWT_SECRET=your_secret_key
JWT_EXPIRE=24h
```

### Start PostgreSQL

Create a database named:

```text
auth_db
```

### Run the application

```bash
go run ./cmd/server
```

The server will start on:

```text
http://localhost:8080
```

## API Endpoints

### Register

```http
POST /register
```

Request

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

### Login

```http
POST /login
```

Request

```json
{
  "email": "example@example.com",
  "password": "password123"
}
```

Response

```json
{
  "token": "your_jwt_token"
}
```

### Get Profile

```http
GET /profile
```

Headers

```http
Authorization: Bearer <your_jwt_token>
```

## Authentication Flow

1. Register a new user.
2. Log in using your email and password.
3. Receive a JWT token.
4. Include the token in the `Authorization` header when accessing protected endpoints.
5. The middleware verifies the token before allowing access.

## Reusing This Service

This authentication service is designed to be reused across multiple projects. Any web, mobile, or desktop application can connect to the exposed REST APIs instead of implementing its own authentication system.

During development, the service runs on a local server:

```text
http://localhost:8080
```

After deployment, the same APIs can be hosted on a public domain and shared with multiple applications.

## Future Improvements

* Refresh tokens
* Role-based access control (RBAC)
* Email verification
* Password reset
* Docker support
* API versioning
* Swagger/OpenAPI documentation
* Unit and integration tests
* Rate limiting
* Logging and monitoring

## ScreenShots
<img width="642" height="517" alt="image" src="https://github.com/user-attachments/assets/e6aac324-58d1-438e-9740-a208128c17b7" />
<img width="635" height="162" alt="image" src="https://github.com/user-attachments/assets/950a9096-40f9-489e-84f7-e34931124876" />
<img width="635" height="162" alt="image" src="https://github.com/user-attachments/assets/1883c221-aca8-4398-bb00-182157c2037f" />
<img width="597" height="125" alt="image" src="https://github.com/user-attachments/assets/42e76517-1e97-472d-b6d7-948ce5ad3870" />
<img width="593" height="213" alt="image" src="https://github.com/user-attachments/assets/685f4181-dd05-4b15-9546-35e4fc6eee3a" />

## License

This project is available for learning, experimentation, and further development.
