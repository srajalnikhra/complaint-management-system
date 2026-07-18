# Complaint Management System Database Documentation

## Overview

The Complaint Management System uses PostgreSQL as its relational database.

The database currently consists of two primary tables:

- Users
- Complaints

A one-to-many relationship exists between users and complaints.

---

## Entity Relationship

```text
Users
-----
id (PK)
name
email
password
role
is_active
otp_hash
otp_expires_at
created_at
updated_at
      │
      │ 1
      │
      │
      ▼
Complaints
----------
id (PK)
user_id (FK)
title
description
status
created_at
updated_at
```

---

## Users Table

Stores all registered users.

| Column | Type | Description |
| --- | --- | --- |
| id | Integer | Primary Key |
| name | Varchar | User name |
| email | Varchar | Unique email |
| password | Text | Hashed password |
| role | Varchar | user / admin |
| is_active | Boolean | Account status |
| otp_hash | Text | Hashed OTP |
| otp_expires_at | Timestamp | OTP expiration time |
| created_at | Timestamp | Creation time |
| updated_at | Timestamp | Last update time |

---

## Complaints Table

Stores complaints created by users.

| Column | Type | Description |
| --- | --- | --- |
| id | Integer | Primary Key |
| user_id | Integer | Foreign Key → Users.id |
| title | Varchar | Complaint title |
| description | Text | Complaint description |
| status | Varchar | Complaint status |
| created_at | Timestamp | Creation time |
| updated_at | Timestamp | Last update time |

---

## Relationship

One user can create multiple complaints.

```text
One User
   │
   │ 
   │
   ▼
Complaints
  Many
```

---

## Authentication Data

Authentication-related information stored in the database includes:

- Hashed password
- User role
- Account status
- OTP hash
- OTP expiration time

JWT tokens are **not stored** in the database.

---

## Complaint Status

Each complaint has one status.

Supported statuses include:

- Pending
- In Progress
- Resolved

---

## Database Constraints

The project uses several constraints to maintain data integrity.

Examples include:

- Primary Keys
- Foreign Keys
- Unique Email
- NOT NULL fields

---

## Security Considerations

The database stores only hashed passwords.

OTP values are stored as hashes and automatically expire after a configured duration.

JWT access tokens are generated during login and validated on every protected request. They are never persisted in the database.

---

## Database Design Principles

The database is designed to be:

- Normalized
- Simple
- Easy to extend
- Consistent
- Secure
