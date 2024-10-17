package controllers

import (
	//"go/token"

	"github.com/gofiber/fiber/v2"
	"github.com/lokesh2201013/config"
	"github.com/lokesh2201013/middleware"
	"github.com/lokesh2201013/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)
func Signup(c *fiber.Ctx) error{
	var data struct {
		Name            string `json:"name"`
		Email           string `json:"email"`
		Password        string `json:"password"`
		UserType        string `json:"user_type"` // Admin or Applicant
		 Profile        models.Profile `gorm:"foreignKey:UserID"`
		//Address         string `json:"address"`
	}

	if err:=c.BodyParser(&data); err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var existingUser models.User
	if err := config.DB.Where("email = ?", data.Email).First(&existingUser).Error; err == nil {
		// Email already exists
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email already in use",
		})
	}


	passwordHash,_:= bcrypt.GenerateFromPassword([]byte(data.Password),bcrypt.DefaultCost)

	user:=models.User{
		Name:            data.Name,
		Email:           data.Email,
		PasswordHash:    string(passwordHash),
		UserType:        data.UserType,
		Profile:          data.Profile,
		//Address:         data.Address,
	}

	config.DB.Create(&user)

	token, err := middleware.GenerateJWT(user.ID, user.UserType)

 
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate JWT",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})

}
func Login(c *fiber.Ctx) error{
	var data struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if err:=c.BodyParser(&data);err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":err.Error()})
	}

	var user models.User

	if err:= config.DB.Where("email=?",data.Email).First(&user).Error; err!=nil{
		if err ==gorm.ErrRecordNotFound{
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"Invalid email or password"})
		}

		return  c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})

	}

	if err:= bcrypt.CompareHashAndPassword([]byte(user.PasswordHash),[]byte(data.Password)); err!=nil{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"Invalid email or password"})
	}

	token, err := middleware.GenerateJWT(user.ID, user.UserType)


	if err!=nil{
		return  c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}

	return  c.Status(fiber.StatusOK).JSON(fiber.Map{"token":token})

}