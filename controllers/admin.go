package controllers

import (
	/*"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"*/
	"time"

	"github.com/gofiber/fiber/v2"
	//"github.com/klauspost/compress/gzhttp/writer"
	"github.com/lokesh2201013/config"
	"github.com/lokesh2201013/models"
)

func CreateJob(c *fiber.Ctx) error{
	var job models.Job

	if err:= c.BodyParser(&job); err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid input"})

	}

	userID:=c.Locals("user_id").(uint)
	job.PostedBy=userID
	job.PostedOn=time.Now().Format("2006-01-02")
     
	config.DB.Create(&job)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Job created successfully", "job": job})

}

func GetJob(c *fiber.Ctx) error {
    jobID := c.Params("job_id")

    var job models.Job
    if err := config.DB.First(&job, jobID).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Job not found"})
    }

    return c.Status(fiber.StatusOK).JSON(job)
}


func GetAllApplicants(c *fiber.Ctx) error{
	var applicants []models.User

	if err:= config.DB.Where("user_type=?","Applicant").Find(&applicants).Error;err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Failed to fetch applicants"})
	}

	return c.Status(fiber.StatusOK).JSON(applicants)
}

func GetApplicantDetails(c *fiber.Ctx)error{
	applicantID:=c.Params("applicant_id")
	var profile models.Profile

	if err := config.DB.Where("user_id = ?", applicantID).First(&profile).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Applicant not found"})
    }

    return c.Status(fiber.StatusOK).JSON(profile)
}

