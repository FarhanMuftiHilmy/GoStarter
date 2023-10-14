package main

import (
	// "database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/farhanmuftihilmy/go-fiber-rest-api/book"
	"github.com/farhanmuftihilmy/go-fiber-rest-api/database"
	"github.com/gofiber/fiber/v2"
	// _ "github.com/go-sql-driver/mysql"
	// "gorm.io/driver/mysql"
)

func hellowWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("mysql", "root:hilmy148@/test_go_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Database connection successfully opened")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Migrated")
}

func main() {
	app := fiber.New()

	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(":3000")
}
