package book

import (
	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	return c.SendString("All Books")
}

func GetBook(c *fiber.Ctx) error {
	return c.SendString("A Single Books")
}

func NewBook(c *fiber.Ctx) error {
	return c.SendString("Adds a new book")
}

func DeleteBook(c *fiber.Ctx) error {
	return c.SendString("Deletes a book")
}

func UpdateBook(c *fiber.Ctx) error {
	return c.SendString("All Books")
}
