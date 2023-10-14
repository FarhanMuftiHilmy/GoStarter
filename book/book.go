package book

import (
	"github.com/farhanmuftihilmy/go-fiber-rest-api/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	return c.JSON(books)
	// return c.SendString("All Books")
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book []Book
	db.Find(&book, id)
	return c.JSON(book)
	// return c.SendString("A Single Books")
}

func NewBook(c *fiber.Ctx) error {
	db := database.DBConn

	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(503).Send([]byte(err.Error()))
	}
	// var book Book
	// book.Title = "1984"
	// book.Author = "George Well"
	// book.Rating = 5
	db.Create(&book)
	return c.JSON(book)
	// return c.SendString("Adds a new book")
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.First(&book, id)
	if book.Title == "" {
		return c.Status(500).SendString("No book found with given id")
	}
	db.Delete(&book)
	return c.SendString("Book successfully deleted")
	// return c.SendString("Deletes a book")
}

func UpdateBook(c *fiber.Ctx) error {
	return c.SendString("All Books")
}
