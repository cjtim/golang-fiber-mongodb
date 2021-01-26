package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func HomeController(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "Hello, world"})
}
func PingController(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"test": "thank"})
}
func PostController(c *fiber.Ctx) error {
	body := new(postBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}
	c.JSON(body)
	return nil
}

type postBody struct {
	Name   string    `json:"name"`
	ArrInt postBody2 `json:"arrInt"`
}

type postBody2 struct {
	Arr []int `json:"arr"`
	Len int   `json:"len"`
}
