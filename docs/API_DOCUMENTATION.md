# Complaint Management System API Documentation

## Overview

The Complaint Management System provides RESTful APIs for user authentication, complaint management, and administrator operations.

All protected endpoints require JWT authentication.

---

## Authentication Flow

1. Register a new account.
2. Login using email and password.
3. Receive a JWT access token.
4. Include the token in the Authorization header.
5. Access protected endpoints.

Example:

Authorization: Bearer <your_jwt_token>

---

## Authorization

There are two roles in the system.

### User

Users can:

- Register
- Login
- Reset password
- Create complaints
- View their own complaints
- Update their own complaints
- Delete their own complaints

### Admin

Administrators can:

- View all complaints
- Update complaint status
- View all users
- Update user roles
- Activate or deactivate users
- Delete users

---

## User APIs

Authentication APIs

- Register
- Login
- Forgot Password
- Verify OTP
- Reset Password

Complaint APIs

- Create Complaint
- Get My Complaints
- Get Complaint By ID
- Update Complaint
- Delete Complaint

---

## Admin APIs

Complaint Management

- Get All Complaints
- Update Complaint Status

User Management

- Get All Users
- Update User Role
- Update User Status
- Delete User

---

## API Response Format

Every API returns a standard JSON response.

Success

```json
{
  "success": true,
  "message": "Operation completed successfully",
  "data": {}
}
```

Error

```json
{
  "success": false,
  "message": "Something went wrong"
}
```

---

## Error Responses

Common error codes include:

- 400 Bad Request
- 401 Unauthorized
- 403 Forbidden
- 404 Not Found
- 409 Conflict
- 500 Internal Server Error

---

## Status Codes

| Code | Description |
| ---- | ----------- |
| 200 | Success |
| 201 | Resource Created |
| 400 | Invalid Request |
| 401 | Authentication Failed |
| 403 | Permission Denied |
| 404 | Resource Not Found |
| 409 | Conflict |
| 500 | Internal Server Error |
