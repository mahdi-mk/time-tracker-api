package auth

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdi-mk/time-tracker/app/models"
	"github.com/mahdi-mk/time-tracker/app/requests"
	"github.com/mahdi-mk/time-tracker/support/hash"
	"github.com/mahdi-mk/time-tracker/support/jwt"
	"github.com/mahdi-mk/time-tracker/support/validator"
	"gorm.io/gorm"
)

type AuthController struct {
	db *gorm.DB
}

func MakeController(db *gorm.DB) AuthController {
	return AuthController{
		db: db,
	}
}

// Register registers a new user to the system
func (cont *AuthController) Register(c *fiber.Ctx) error {
	request := new(requests.RegisterUserRequest)

	if err := validator.ValidateRequest(c, request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	var user models.User

	err := cont.db.First(&user, "email = ?", request.Email).Error

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

	cont.db.Create(&user)

	token, _ := jwt.GenerateToken(&jwt.TokenPayload{UserID: user.ID})

	return c.JSON(fiber.Map{
		"user":  user,
		"token": token,
	})
}

// Login logins an existing user to the system
func (cont *AuthController) Login(c *fiber.Ctx) error {
	request := new(requests.LoginUserRequest)

	if err := validator.ValidateRequest(c, request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	var user models.User
	err := cont.db.First(&user, "email = ?", request.Email).Error

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
