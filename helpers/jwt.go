package helpers

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func GenerateToken(id uint, email string) string {
// Create the Claims
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		log.Fatalln(err)
	}

	return signedToken
}

func VerifyToken(ctx *fiber.Ctx) interface{} {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims
}