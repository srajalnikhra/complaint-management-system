# Complaint Management System Architecture

## Project Structure

The project follows a layered architecture where every layer has a single responsibility.

```text
Client
    │
    ▼
Routes
    │
    ▼
Middleware
    │
    ▼
Controllers
    │
    ▼
Services
    │
    ▼
Repositories
    │
    ▼
PostgreSQL
```

---

## Request Lifecycle

A typical request passes through the following layers:

1. HTTP Request
2. Route Matching
3. Middleware Execution
4. Controller
5. Service
6. Repository
7. Database
8. JSON Response

---

## Layered Architecture

### Routes

Responsible for registering API endpoints.

### Middleware

Responsible for:

- Authentication
- Authorization
- Logging
- Rate Limiting
- CORS

### Controllers

Handle HTTP requests and responses.

Responsibilities include:

- Parsing request body
- Input validation
- Calling services
- Returning standardized responses

### Services

Contain business logic.

Responsibilities include:

- Validation
- Business rules
- Calling repositories

### Repositories

Handle database operations.

Responsibilities include:

- CRUD operations
- SQL queries
- Database abstraction

### Database

PostgreSQL stores all persistent application data.

---

## Authentication Flow

```text
User Login
     │
     ▼
Password Verification
     │
     ▼
JWT Generation
     │
     ▼
Client Stores Token
     │
     ▼
Authorization Header
     │
     ▼
Auth Middleware
     │
     ▼
Protected Endpoint
```

---

## Folder Responsibilities

| Folder | Responsibility |
| --------- | ---------------- |
| config | Application & JWT configuration |
| controllers | HTTP handlers |
| database | Database initialization |
| dto | Request and response DTOs |
| middleware | Authentication, Authorization, Logging, CORS, Rate Limiting |
| models | Database models |
| repositories | Database access |
| routes | Route registration |
| services | Business logic |
| utils | Helper functions |
| validation | Request validation |

---

## Design Principles

The project follows the following principles:

- Separation of Concerns
- Layered Architecture
- Dependency Injection (through constructors)
- Reusable Components
- Standardized API Responses
- Centralized Validation
- Repository Pattern
