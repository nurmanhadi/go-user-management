# User Management Service

A robust user management microservice built with Go for handling user registration, profile management, and authentication in a microservices architecture.

## Tech Stack

- **Language:** Go 1.21+
- **Database:** PostgreSQL 15+
- **Cache:** Memcached
- **Message Broker:** LavinMQ / RabbitMQ
- **Containerization:** Docker & Docker Compose
- **Orchestration:** Kubernetes

## Prerequisites

- Go 1.21 or higher
- Docker & Docker Compose
- PostgreSQL 15+
- Memcached
- LavinMQ or RabbitMQ

## Quick Start

### 1. Clone the Repository

```bash
git clone https://github.com/nurmanhadi/go-user-management.git
cd user-management
```

### 2. Configure Environment Variables

```bash
cp .env.example .env
```

Edit `.env` with your configuration:

```bash
# Database Configuration
DB_HOST=localhost
DB_PORT=5433
DB_USERNAME=user
DB_PASSWORD=user
DB_NAME=user_management

# Cache Configuration
CACHE_HOST=localhost
CACHE_PORT=11212

# Message Broker Configuration
BROKER_HOST=localhost
BROKER_PORT=5672
BROKER_USERNAME=guest
BROKER_PASSWORD=guest
BROKER_VHOST=someone
```

### 3. Start Services with Docker Compose

```bash
docker-compose up -d
```

### 4. Run Database Migrations

```bash
make migrate-up
```

## Development

### Run Locally

```bash
go run cmd/main.go
```

### Run Tests

```bash
go test ./...
```

### Build Binary

```bash
go build -o bin/user-service cmd/main.go
```

## API Endpoints

### Update User Profile

**Endpoint:** `PUT /api/users/{id}`

**Request Body:**

```json
{
  "first_name": "testing",
  "last_name": "testing juga",
  "email": "test@gmail.com",
  "phone": "085734621178",
  "gender": "male",
  "bio": "testing",
  "description": "test",
  "birth_date": "2002-11-12"
}
```

**Response Codes:**
- `200` - OK (User profile updated successfully)
- `400` - Bad Request (Invalid input data)
- `404` - Not Found (User not found)
- `409` - Conflict (Email or phone already exists)

### Get User by Username

**Endpoint:** `GET /api/users/{username}`

**Response Body:**

```json
{
  "data": {
    "id": 3,
    "auth_id": "29c8d0b6-737b-494e-b8f5-934e64dd7ea7",
    "username": "test12",
    "name": {
      "first_name": null,
      "last_name": null
    },
    "contact": {
      "email": null,
      "phone": null
    },
    "about": {
      "bio": null,
      "description": null,
      "birth_date": null,
      "gender": null
    },
    "verification": {
      "email_verified_at": null,
      "phone_verified_at": null
    },
    "avatar_url": null,
    "created_at": "2025-11-13T22:49:49.549084+07:00",
    "updated_at": "2025-11-13T22:49:49.562617+07:00"
  },
  "path": "/api/users/test12"
}
```

**Response Codes:**
- `200` - OK (User found and returned)
- `404` - Not Found (User not found)
- `500` - Internal Server Error

### Service Get User by Id

**Endpoint:** `GET /api/users/services/{id}`

**Response Body:**

```json
{
  "data": {
    "id": 3,
    "auth_id": "29c8d0b6-737b-494e-b8f5-934e64dd7ea7",
    "username": "test12",
    "name": {
      "first_name": null,
      "last_name": null
    },
    "contact": {
      "email": null,
      "phone": null
    },
    "about": {
      "bio": null,
      "description": null,
      "birth_date": null,
      "gender": null
    },
    "verification": {
      "email_verified_at": null,
      "phone_verified_at": null
    },
    "avatar_url": null,
    "created_at": "2025-11-13T22:49:49.549084+07:00",
    "updated_at": "2025-11-13T22:49:49.562617+07:00"
  },
  "path": "/api/users/test12"
}
```

**Response Codes:**
- `200` - OK (User found and returned)
- `404` - Not Found (User not found)
- `500` - Internal Server Error

## Event-Driven Architecture

### Subscribed Events

The service subscribes to events from the message broker published by other microservices.

#### User Registered Event

**Event Name:** `user.registered`  
**Queue:** `user.registered`

Triggered when a new user successfully registers.

**Payload:**

```json
{
  "event": "user.registered",
  "timestamp": "2025-11-12T10:30:00Z",
  "data": {
    "user_id": "uuid-here",
    "username": "username",
    "registered_at": "2025-11-12T10:30:00Z"
  }
}
```

#### User Avatar Updated Event

**Event Name:** `user.avatar`  
**Queue:** `user.avatar`

Triggered when a user's avatar is uploaded or updated.

**Payload:**

```json
{
  "event": "user.avatar",
  "timestamp": "2025-11-12T10:30:00Z",
  "data": {
    "user_id": 1,
    "avatar_url": "http://something"
  }
}
```

## Security Considerations

- Enable TLS/SSL for production deployments
- Use HTTPS for all API endpoints
- Validate and sanitize all user inputs
- Implement rate limiting on authentication endpoints
- Use environment variables for sensitive configuration (API keys, database credentials)
- Implement proper authentication and authorization mechanisms
- Enable CORS only for trusted domains in production
- Keep dependencies updated regularly

## Project Structure

```
.
├── cmd/
│   └── main.go
├── internal/
│   ├── api/
│   ├── service/
│   ├── repository/
│   └── models/
├── docker-compose.yml
├── Dockerfile
├── Makefile
├── go.mod
├── go.sum
└── README.md
```

## License

This project is licensed under the MIT License.

## Author

**Nurman Hadi**  
Backend Developer (Golang, Microservices)  
GitHub: [@nurmanhadi](https://github.com/nurmanhadi)