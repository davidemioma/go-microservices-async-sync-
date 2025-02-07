package utils

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func RespondWithJSON(c *fiber.Ctx, code int, payload interface{}) {
    err := c.JSON(payload)

	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)

		c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})

		return
	}

	c.Status(code).JSON(payload)
}

func RespondWithError(c *fiber.Ctx, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}
	
	RespondWithJSON(c, code, errorResponse{
		Error: msg,
	})
}

