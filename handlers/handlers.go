package handlers

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

func GetRoot(c *fiber.Ctx) error {
	return c.SendFile("public/index.html")
}

func GetStatic(c *fiber.Ctx) error {
	filePath := c.Params("*")

	if _, err := os.Stat("public/" + filePath); os.IsNotExist(err) {
		return c.SendString("Cannot GET /" + filePath)
	}

	return c.SendFile("public/" + filePath)
}

func PostCalculate(c *fiber.Ctx) error {
	expression := string(c.Body())

	answer, err := evaluateExpression(expression)
	if err != nil {
		return c.Status(400).SendString("Invalid Expression")
	}

	return c.SendString(answer)
}

// func calculate(c *fiber.Ctx) error {
//		response := struct {
//			Expression string `json:"expression"`
//		}{}
//		if err := json.Unmarshal(c.Body(), &response); err != nil {
//			return c.Status(400).SendString("Invalid request body")
//		}
//
//		expression := response.Expression
//
//		answer, err := evaluateExpression(expression)
//		if err != nil {
//			return c.Status(400).SendString("Invalid expression")
//		}
//
//		return c.SendString(answer)
//	})
