package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	//"github.com/klauspost/compress/gzhttp/writer"
	"github.com/lokesh2201013/config"
	"github.com/lokesh2201013/models"
)

func UploadResume(c *fiber.Ctx) error{
	userID:=c.Locals("user_id").(uint)
    
	file,err:= c.FormFile("resume")
	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Resume file is required"})
	}

	ext :=filepath.Ext(file.Filename)
	if ext != ".pdf" && ext!=".docx"{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Only pdf and docx are allowed"})
	}

	fileData ,err := file.Open()

	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Failed to read file"})

	}

	body := &bytes.Buffer{}
	writer:=multipart.NewWriter(body)
	part,_:=writer.CreateFormFile("file",file.Filename)

	io.Copy(part,fileData)
	writer.Close()

	request,_:=http.NewRequest("POST","https://api.apilayer.com/resume_parser/upload",body)
	request.Header.Add("Content-Type",writer.FormDataContentType())
	request.Header.Add("apikey", "0bWeisRWoLj3UdXt3MXMSMWptYFIpQfS")

	client:=&http.Client{}

	response,_:=client.Do(request)

	defer response.Body.Close()

	var resumeData map[string]interface{}

	json.NewDecoder(response.Body).Decode(&resumeData)

	var profile models.Profile
    profile.UserID = userID
    profile.ResumeFile = file.Filename
    profile.Skills = strings.Join(resumeData["skills"].([]string), ", ")
    profile.Education = strings.Join(resumeData["education"].([]string), ", ")
    profile.Experience = strings.Join(resumeData["experience"].([]string), ", ")

    config.DB.Create(&profile)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message":"Resume uploaded successfully ","profile":profile})
}