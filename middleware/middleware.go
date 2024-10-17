package middleware

import (
	//"go/token"
	//"os"
	"time"
      "fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware() fiber.Handler{
	return func(c *fiber.Ctx) error{

		tokenString:=c.Get("Authorization")
	    
		if tokenString==""{
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token format",
			})
		}

		token ,err := jwt.Parse(tokenString , func(token *jwt.Token)(interface {},error){
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "unexpected signing method")
			}
			return []byte("mysecretkey123"),nil
		})
		if err != nil || !token.Valid{
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": fmt.Sprintf("Token error: %v", err),
			})
		}
		
        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "Invalid token claims",
            })
        }
          
		userID, ok := claims["user_id"].(float64) // JWT stores numbers as float64
if !ok {
    return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
        "error": "Missing user_id in token",
    })
}
        // Extract user_type from JWT claims and set it in Locals
        userType, ok := claims["user_type"].(string)
        if !ok {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "Missing user_type in token",
            })
        }
		c.Locals("user_id", uint(userID))
        c.Locals("user_type", userType)
        return c.Next()

	}
}

func GenerateJWT(userID uint , userType string)(string ,error){
	claims:= jwt.MapClaims{
		"user_id":userID,
		"user_type": userType, 
		"exp": time.Now().Add(time.Hour*24).Unix(),

	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("mysecretkey123"))
}
func AdminOnly(next fiber.Handler) fiber.Handler {
    return func(c *fiber.Ctx) error {
        // Log the value to check what is stored
        fmt.Println("user_type:", c.Locals("user_type"))

        userType, ok := c.Locals("user_type").(string)
        if !ok {
            // Log the issue if type assertion fails
            fmt.Println("Type assertion failed for user_type")
            return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Admins only"})
        }

        if userType != "Admin" {
            return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Admins only"})
        }

        return next(c)
    }
}


func ApplicantOnly(next fiber.Handler) fiber.Handler {
    return func(c *fiber.Ctx) error {
        userType,ok := c.Locals("user_type").(string)
        if !ok||userType != "Applicant" {
            return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Applicants only"})
        }
        return next(c)
    }
}
