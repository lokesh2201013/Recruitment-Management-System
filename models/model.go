package models

import "gorm.io/gorm"

// User model for Admins and Applicants
type User struct {
    gorm.Model
    Name           string `json:"name"`
    Email          string `json:"email" gorm:"unique"`
    PasswordHash   string `json:"-"`
    UserType       string `json:"user_type"`  // Either "Admin" or "Applicant"
    Profile        Profile `gorm:"foreignKey:UserID"`
}

// Profile model for Applicants
type Profile struct {
    gorm.Model
    UserID         uint
    ResumeFile     string `json:"resume_file"`
    Skills         string `json:"skills"`
    Education      string `json:"education"`
    Experience     string `json:"experience"`
    Phone          string `json:"phone"`
}

// Job model for job postings
type Job struct {
    gorm.Model
    Title          string `json:"title"`
    Description    string `json:"description"`
    PostedOn       string `json:"posted_on"`
    TotalApplications int  `json:"total_applications"`
    CompanyName    string `json:"company_name"`
    PostedBy       uint   `json:"posted_by"`  // Admin's User ID
}
