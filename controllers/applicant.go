package controllers

import (
	//"bytes"
	//"encoding/json"
	//"io"
	//"mime/multipart"
	//"net/http"
	//"path/filepath"
	//"strings"
	   //"time"

	"github.com/gofiber/fiber/v2"
	//"github.com/klauspost/compress/gzhttp/writer"
	"github.com/lokesh2201013/config"
	"github.com/lokesh2201013/models"
)
func ListJobs(c *fiber.Ctx) error {
    var jobs []models.Job

    if err := config.DB.Find(&jobs).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch jobs"})
    }

    return c.Status(fiber.StatusOK).JSON(jobs)
}

func ApplyForJob(c *fiber.Ctx) error {
    jobID := c.Query("job_id")
    //userID := c.Locals("user_id").(uint)

    var job models.Job
    if err := config.DB.First(&job, jobID).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Job not found"})
    }

    job.TotalApplications++
    config.DB.Save(&job)

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Job applied successfully"})
}
