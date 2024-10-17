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

## Setup Instructions

1. **Clone the Repository**

   ```bash
   git clone https://github.com/yourusername/yourrepository.git
   cd yourrepository
   ```

2. **Install Dependencies**

   ```bash
   go mod download
   ```

3. **Configure the Database**

   - Update the `config` package to include your database connection details.
   - Example using GORM:

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

   ```bash
   go run main.go
   ```

   - The server will start on `http://localhost:8080`.

## API Endpoints

### Authentication

#### Sign Up

- **URL**: `/signup`
- **Method**: `POST`
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

**Example using `curl`**

```bash
curl -X POST http://localhost:8080/signup \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "johndoe@example.com",
    "password": "password123",
    "user_type": "Applicant"
  }'
```

**Response**

```json
{
  "token": "your_jwt_token"
}
```

---

#### Log In

- **URL**: `/login`
- **Method**: `POST`
- **Description**: Log in with email and password to receive a JWT token.

**Request Body**

```json
{
  "email": "johndoe@example.com",
  "password": "password123"
}
```

**Example using `curl`**

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "johndoe@example.com",
    "password": "password123"
  }'
```

**Response**

```json
{
  "token": "your_jwt_token"
}
```

---

### Admin Routes

*Note: All Admin routes require a valid JWT token from an Admin user. Include the token in the `Authorization` header as `Bearer your_jwt_token`.*

#### Create Job

- **URL**: `/admin/job`
- **Method**: `POST`
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

**Example using `curl`**

```bash
curl -X POST http://localhost:8080/admin/job \
  -H "Authorization: Bearer your_jwt_token" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Software Engineer",
    "description": "Responsible for building scalable web applications.",
    "company_name": "TechCorp"
  }'
```

**Response**

```json
{
  "message": "Job created successfully",
  "job": {
    "ID": 1,
    "CreatedAt": "2024-10-17T12:00:00Z",
    "UpdatedAt": "2024-10-17T12:00:00Z",
    "DeletedAt": null,
    "title": "Software Engineer",
    "description": "Responsible for building scalable web applications.",
    "posted_on": "2024-10-17",
    "total_applications": 0,
    "company_name": "TechCorp",
    "posted_by": 1
  }
}
```

---

#### Get Job Details

- **URL**: `/admin/job/:job_id`
- **Method**: `GET`
- **Description**: Admins can retrieve details of a specific job.

**Headers**

```http
Authorization: Bearer your_jwt_token
```

**Example using `curl`**

```bash
curl -X GET http://localhost:8080/admin/job/1 \
  -H "Authorization: Bearer your_jwt_token"
```

**Response**

```json
{
  "ID": 1,
  "CreatedAt": "2024-10-17T12:00:00Z",
  "UpdatedAt": "2024-10-17T12:00:00Z",
  "DeletedAt": null,
  "title": "Software Engineer",
  "description": "Responsible for building scalable web applications.",
  "posted_on": "2024-10-17",
  "total_applications": 0,
  "company_name": "TechCorp",
  "posted_by": 1
}
```

---

#### Get All Applicants

- **URL**: `/admin/applicants`
- **Method**: `GET`
- **Description**: Admins can retrieve a list of all applicants.

**Headers**

```http
Authorization: Bearer your_jwt_token
```

**Example using `curl`**

```bash
curl -X GET http://localhost:8080/admin/applicants \
  -H "Authorization: Bearer your_jwt_token"
```

**Response**

```json
[
  {
    "ID": 2,
    "CreatedAt": "2024-10-17T12:10:00Z",
    "UpdatedAt": "2024-10-17T12:10:00Z",
    "DeletedAt": null,
    "name": "Jane Smith",
    "email": "janesmith@example.com",
    "user_type": "Applicant",
    "Profile": {
      // Profile details
    }
  },
  // ... other applicants
]
```

---

#### Get Applicant Details

- **URL**: `/admin/applicant/:applicant_id`
- **Method**: `GET`
- **Description**: Admins can retrieve the profile details of a specific applicant.

**Headers**

```http
Authorization: Bearer your_jwt_token
```

**Example using `curl`**

```bash
curl -X GET http://localhost:8080/admin/applicant/2 \
  -H "Authorization: Bearer your_jwt_token"
```

**Response**

```json
{
  "ID": 1,
  "CreatedAt": "2024-10-17T12:15:00Z",
  "UpdatedAt": "2024-10-17T12:15:00Z",
  "DeletedAt": null,
  "UserID": 2,
  "resume_file": "JaneResume.pdf",
  "skills": "Go, JavaScript, SQL",
  "education": "University X (2015-2019)",
  "experience": "Company A (2019-2021)",
  "phone": "+1234567890"
}
```

---

### Applicant Routes

*Note: All Applicant routes require a valid JWT token from an Applicant user. Include the token in the `Authorization` header as `Bearer your_jwt_token`.*

#### Upload Resume

- **URL**: `/uploadResume`
- **Method**: `POST`
- **Description**: Applicants can upload their resume, which will be parsed and stored.

**Headers**

```http
Authorization: Bearer your_jwt_token
Content-Type: multipart/form-data
```

**Form Data**

- **Key**: `resume`
- **Value**: Select a PDF or DOCX file from your computer.

**Example using `curl`**

```bash
curl -X POST http://localhost:8080/uploadResume \
  -H "Authorization: Bearer your_jwt_token" \
  -F "resume=@/path/to/your/resume.pdf"
```

**Response**

```json
{
  "message": "Resume uploaded successfully",
  "profile": {
    "ID": 1,
    "CreatedAt": "2024-10-17T12:20:00Z",
    "UpdatedAt": "2024-10-17T12:20:00Z",
    "DeletedAt": null,
    "UserID": 2,
    "resume_file": "resume.pdf",
    "skills": "Go, JavaScript, SQL",
    "education": "University X (2015-2019)",
    "experience": "Company A (2019-2021)",
    "phone": "+1234567890"
  }
}
```

---

#### List Jobs

- **URL**: `/jobs`
- **Method**: `GET`
- **Description**: Applicants can view a list of available jobs.

**Headers**

- No authentication required.

**Example using `curl`**

```bash
curl -X GET http://localhost:8080/jobs
```

**Response**

```json
[
  {
    "ID": 1,
    "CreatedAt": "2024-10-17T12:00:00Z",
    "UpdatedAt": "2024-10-17T12:00:00Z",
    "DeletedAt": null,
    "title": "Software Engineer",
    "description": "Responsible for building scalable web applications.",
    "posted_on": "2024-10-17",
    "total_applications": 1,
    "company_name": "TechCorp",
    "posted_by": 1
  },
  // ... other jobs
]
```

---

#### Apply for Job

- **URL**: `/jobs/apply`
- **Method**: `GET`
- **Description**: Applicants can apply for a job by specifying the job ID.

**Headers**

```http
Authorization: Bearer your_jwt_token
```

**Query Parameters**

- `job_id`: The ID of the job to apply for.

**Example using `curl`**

```bash
curl -X GET http://localhost:8080/jobs/apply?job_id=1 \
  -H "Authorization: Bearer your_jwt_token"
```

**Response**

```json
{
  "message": "Job applied successfully"
}
```

---

## Testing the API with Postman

1. **Import the API Collection**

   - Create a new collection in Postman and add the endpoints as described above.

2. **Set Up Environment Variables**

   - Create environment variables for `base_url` (e.g., `http://localhost:8080`) and `token`.

3. **Sign Up and Log In**

   - Use the **Sign Up** endpoint to create a new user.
   - Use the **Log In** endpoint to obtain a JWT token.
   - Set the `token` environment variable with the obtained JWT token.

4. **Add Authorization Header**

   - For authenticated routes, add an `Authorization` header with the value `Bearer {{token}}`.

5. **Test Admin and Applicant Routes**

   - Use the appropriate JWT tokens for Admin and Applicant roles when testing the respective routes.

## Notes

- **JWT Secret Key**: Ensure that the JWT secret key (`mysecretkey123`) is kept secure and ideally stored in environment variables, not hardcoded.
- **API Key for Resume Parsing**: Replace the placeholder API key (`0bWeisRWoLj3UdXt3MXMSMWptYFIpQfS`) with your actual API key from the resume parsing service.
- **Error Handling**: The API returns error messages in JSON format. Ensure to handle these appropriately in your client application.
- **Database Migrations**: The `config.DB.AutoMigrate()` function will automatically create tables based on your models. Ensure your models are defined correctly.

---

Feel free to reach out if you have any questions or need further assistance with the API!
