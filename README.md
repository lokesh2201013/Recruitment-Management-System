To make the **Recruitment-Management-System** compatible with both Windows and Linux systems, we need to ensure that commands and configurations work seamlessly on both operating systems. Below are the modifications to the setup and instructions to accommodate both platforms.

---

# Recruitment-Management-System

This document provides a step-by-step guide on how to set up and use the Job Application API. The API allows users to sign up as either an Admin or an Applicant, post and apply for jobs, and manage user profiles.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Setup Instructions](#setup-instructions)
- [API Endpoints](#api-endpoints)
  - [Authentication](#authentication)
    - [Sign Up](#sign-up)
    - [Log In](#log-in)
  - [Admin Routes](#admin-routes)
    - [Create Job](#create-job)
    - [Get Job Details](#get-job-details)
    - [Get All Applicants](#get-all-applicants)
    - [Get Applicant Details](#get-applicant-details)
  - [Applicant Routes](#applicant-routes)
    - [Upload Resume](#upload-resume)
    - [List Jobs](#list-jobs)
    - [Apply for Job](#apply-for-job)
- [Testing the API with Postman](#testing-the-api-with-postman)
- [Notes](#notes)

## Prerequisites

- **Go**: Install [Go](https://golang.org/dl/) (version 1.16 or later).
- **Database**: Set up a PostgreSQL database or adjust the configuration for your preferred database.
- **Tools**: Install [Postman](https://www.postman.com/downloads/) or [curl](https://curl.se/download.html) for testing API endpoints.
- **OS Compatibility**: Ensure you're using compatible commands for Windows or Linux.

## Setup Instructions

1. **Clone the Repository**

   For **Windows** users, run this in **Git Bash** or **PowerShell**. For **Linux**, use the terminal.

   **Windows:**

   ```bash
   git clone https://github.com/yourusername/yourrepository.git
   cd yourrepository
   ```

   **Linux:**

   ```bash
   git clone https://github.com/yourusername/yourrepository.git
   cd yourrepository
   ```

2. **Install Dependencies**

   For **both** Windows and Linux, you can use the same `go mod download` command:

   ```bash
   go mod download
   ```

3. **Configure the Database**

   Ensure that your database connection string is correctly configured for both environments.

   **Linux** systems typically connect to the database via `localhost` or `127.0.0.1`.

   **Windows** users connecting to PostgreSQL should make sure the server is accessible via a local connection (the same as in Linux).

   Example using GORM for both:

   ```go
   // config/db.go
   package config

   import (
       "gorm.io/driver/postgres"
       "gorm.io/gorm"
   )

   var DB *gorm.DB

   func ConnectDB() {
       dsn := "host=localhost user=postgres password=yourpassword dbname=yourdb port=5432 sslmode=disable"
       db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
       if err != nil {
           panic("Failed to connect to database!")
       }

       db.AutoMigrate(&models.User{}, &models.Profile{}, &models.Job{})
       DB = db
   }
   ```

4. **Run the Application**

   For **Windows**:

   Open **PowerShell** or **Command Prompt** and run:

   ```bash
   go run main.go
   ```

   For **Linux**:

   Open a terminal and run:

   ```bash
   go run main.go
   ```

   In both cases, the server will start at http://localhost:8080.

## API Endpoints

### Authentication

#### Sign Up

- **URL**: /signup
- **Method**: POST
- **Description**: Register a new user as an Admin or Applicant.

**Request Body**

```json
{
  "name": "John Doe",
  "email": "johndoe@example.com",
  "password": "password123",
  "user_type": "Applicant" // or "Admin"
}
```

**Example using curl**

For **Windows**, replace `\'` with `\"` or use double quotes in PowerShell:

```bash
curl -X POST http://localhost:8080/signup -H "Content-Type: application/json" -d "{\"name\":\"John Doe\", \"email\":\"johndoe@example.com\", \"password\":\"password123\", \"user_type\":\"Applicant\"}"
```

For **Linux**, use single quotes as normal:

```bash
curl -X POST http://localhost:8080/signup -H "Content-Type: application/json" -d '{"name":"John Doe", "email":"johndoe@example.com", "password":"password123", "user_type":"Applicant"}'
```

---

#### Log In

- **URL**: /login
- **Method**: POST
- **Description**: Log in with email and password to receive a JWT token.

**Request Body**

```json
{
  "email": "johndoe@example.com",
  "password": "password123"
}
```

**Example using curl**

For **Windows**:

```bash
curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d "{\"email\":\"johndoe@example.com\", \"password\":\"password123\"}"
```

For **Linux**:

```bash
curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"email":"johndoe@example.com", "password":"password123"}'
```

---

### Admin Routes

*Note: All Admin routes require a valid JWT token from an Admin user. Include the token in the Authorization header as Bearer your_jwt_token.*

#### Create Job

- **URL**: /admin/job
- **Method**: POST
- **Description**: Admins can create new job postings.

**Headers**

```http
Authorization: Bearer your_jwt_token
Content-Type: application/json
```

**Request Body**

```json
{
  "title": "Software Engineer",
  "description": "Responsible for building scalable web applications.",
  "company_name": "TechCorp"
}
```

**Example using curl**

For **Windows**:

```bash
curl -X POST http://localhost:8080/admin/job -H "Authorization: Bearer your_jwt_token" -H "Content-Type: application/json" -d "{\"title\":\"Software Engineer\",\"description\":\"Responsible for building scalable web applications.\",\"company_name\":\"TechCorp\"}"
```

For **Linux**:

```bash
curl -X POST http://localhost:8080/admin/job -H "Authorization: Bearer your_jwt_token" -H "Content-Type: application/json" -d '{"title":"Software Engineer","description":"Responsible for building scalable web applications.","company_name":"TechCorp"}'
```

---

#### Other Admin and Applicant Routes

For all remaining routes (such as `Get Job Details`, `Get All Applicants`, etc.), the same considerations for command formatting apply. When using **curl** on Windows, remember to use double quotes (`"`) instead of single quotes (`'`).

---

## Testing the API with Postman

1. **Import the API Collection**

   - Create a new collection in Postman and add the endpoints as described above.

2. **Set Up Environment Variables**

   - Create environment variables for base_url (e.g., http://localhost:8080) and token.

3. **Sign Up and Log In**

   - Use the **Sign Up** endpoint to create a new user.
   - Use the **Log In** endpoint to obtain a JWT token.
   - Set the token environment variable with the obtained JWT token.

4. **Add Authorization Header**

   - For authenticated routes, add an Authorization header with the value Bearer {{token}}.

5. **Test Admin and Applicant Routes**

   - Use the appropriate JWT tokens for Admin and Applicant roles when testing the respective routes.

## Notes

- **OS-specific Considerations**: When switching between Windows and Linux, ensure that paths (like file paths for resume uploads) are adjusted accordingly.
- **JWT Secret Key**: Ensure that the JWT secret key (mysecretkey123) is kept secure and ideally stored in environment variables, not hardcoded.
- **Database Configuration**: Ensure the PostgreSQL service is running and accessible on your operating system.
