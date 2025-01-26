# TaskFlow
TaskFlow is a Go-based application designed to manage and automate tasks efficiently.

## Installation Guide

### 1. Get the Code
```sh
git clone https://github.com/vsennikov/taskFlow
cd taskFlow
```

### 2. Environment Setup
Create `.env` file in the root directory with these settings:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=yourusername
DB_PASSWORD=yourpassword
DB_NAME=taskflow
JWT_SECRET=yourjwtsecret
```

### 3. Database Setup

First, install PostgreSQL:
- Ubuntu: `sudo apt-get install postgresql`
- MacOS: `brew install postgresql`
- Windows: Download from [PostgreSQL website](https://www.postgresql.org/download/windows/)

Initialize the database:
```sh
psql -U postgres
CREATE DATABASE taskflow;
CREATE USER yourusername WITH ENCRYPTED PASSWORD 'yourpassword';
GRANT ALL PRIVILEGES ON DATABASE taskflow TO yourusername;
```

Create required tables:
```sql
CREATE TABLE users (
	 id SERIAL PRIMARY KEY,
	 username VARCHAR(50) UNIQUE NOT NULL,
	 password_hash VARCHAR(255) NOT NULL,
	 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tasks (
	 id SERIAL PRIMARY KEY,
	 user_id INTEGER REFERENCES users(id),
	 title VARCHAR(100) NOT NULL,
	 description TEXT,
	 status VARCHAR(20) DEFAULT 'pending',
	 due_date TIMESTAMP,
	 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 4. Run the Application
```sh
go run main.go
```

### 5. Available Endpoints

Authentication:
- `POST /api/v1/auth/register` - Register a new user
- `POST /api/v1/auth/login` - Login and receive JWT token

Tasks:
- `GET /api/v1/tasks` - List all tasks
- `POST /api/v1/tasks` - Create a new task
- `GET /api/v1/tasks/{id}` - Get task details
- `PUT /api/v1/tasks/{id}` - Update task
- `DELETE /api/v1/tasks/{id}` - Delete task

Note: Use the JWT token from login in the `Authorization` header for authenticated endpoints.

### API Documentation
You can find the complete API documentation in the `taskflow-api.json` file. This file can be imported into Postman for easier API testing and exploration.

To import into Postman:
1. Open Postman
2. Click "Import" button
3. Select the `taskflow-api.json` file
4. All endpoints will be available in a new collection