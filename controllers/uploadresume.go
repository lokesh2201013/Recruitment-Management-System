package controllers

import (
	//"bytes"
	"encoding/json"
	"io"
	//"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/lokesh2201013/config"
	"github.com/lokesh2201013/models"
)

type ResumeResponse struct {
	Name      string   `json:"name"`
	Address   string   `json:"address"`
	Email     string   `json:"email"`
	Phone     string   `json:"phone"`
	Skills    []string `json:"skills"`
	Education []struct {
		Name  string   `json:"name"`
		Dates []string `json:"dates"`
	} `json:"education"`
	Experience []string `json:"experience"`
}

func UploadResume(c *fiber.Ctx) error {
	// Get user ID from context
	userID := c.Locals("user_id").(uint)

	// Get resume file from request
	file, err := c.FormFile("resume")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Resume file is required"})
	}

	// Validate file extension
	if ext := filepath.Ext(file.Filename); ext != ".pdf" && ext != ".docx" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Only pdf and docx are allowed"})
	}

	// Read file data
	fileData, err := file.Open()
if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to read file"})
}
defer fileData.Close()


req, err := http.NewRequest("POST", "https://api.apilayer.com/resume_parser/upload", fileData)
if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create request"})
}
req.Header.Set("Content-Type", "application/octet-stream")
req.Header.Set("apikey", "0bWeisRWoLj3UdXt3MXMSMWptYFIpQfS")

client := &http.Client{}
resp, err := client.Do(req)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to call resume parser API"})
	}
	defer resp.Body.Close()


	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return c.Status(resp.StatusCode).JSON(fiber.Map{"error": string(body)})
	}

	
	var resumeData ResumeResponse
	if err := json.NewDecoder(resp.Body).Decode(&resumeData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse resume data"})
	}

	// Create and save profile in the database
	profile := models.Profile{
		UserID:     userID,
		ResumeFile: file.Filename,
		Skills:     strings.Join(resumeData.Skills, ", "),
		Education:  formatEducation(resumeData.Education),
		Experience: strings.Join(resumeData.Experience, ", "),
		Phone:      resumeData.Phone,
	}
	if err := config.DB.Create(&profile).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save profile"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Resume uploaded successfully", "profile": profile})
}


func formatEducation(edu []struct {
	Name  string   `json:"name"`
	Dates []string `json:"dates"`
}) string {
	var education []string
	for _, e := range edu {
		education = append(education, e.Name+" ("+strings.Join(e.Dates, ", ")+")")
	}
	return strings.Join(education, ", ")
}
