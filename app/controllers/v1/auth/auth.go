package auth

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdi-mk/time-tracker/app/models"
	"github.com/mahdi-mk/time-tracker/database"
	"github.com/mahdi-mk/time-tracker/utils"
	"github.com/mahdi-mk/time-tracker/utils/hash"
	"github.com/mahdi-mk/time-tracker/utils/jwt"
	"gorm.io/gorm"
)

// Register registers a new user to the system
func Register(c *fiber.Ctx) error {
	request := new(models.RegisterUserRequest)

	if err := utils.ValidateRequest(c, request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	var user models.User

	err := database.DB.First(&user, "email = ?", request.Email).Error

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email already in use",
		})
	}

	user = models.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  hash.Generate(request.Password),
	}

	database.DB.Create(&user)

	token, _ := jwt.GenerateToken(&jwt.TokenPayload{UserID: user.ID})

	return c.JSON(fiber.Map{
		"user":  user,
		"token": token,
	})
}

// Login logins an existing user to the system
func Login(c *fiber.Ctx) error {
	request := new(models.LoginUserRequest)

	if err := utils.ValidateRequest(c, request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	var user models.User
	err := database.DB.First(&user, "email = ?", request.Email).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	if err := hash.Verify(user.Password, request.Password); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	token, _ := jwt.GenerateToken(&jwt.TokenPayload{UserID: user.ID})

	return c.JSON(fiber.Map{
		"user":  user,
		"token": token,
	})
}
