package main

import (
	"database/sql"
	"fmt"

	"github.com/farhanmuftihilmy/go-fiber-rest-api/book"
	"github.com/gofiber/fiber/v2"

	_ "github.com/go-sql-driver/mysql"
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

// func initDatabase() {

// }

func main() {
	app := fiber.New()
	// initDatabase() // open connection
	// defer database.DBConn.Close()
	fmt.Println("Go MySQL Tutorial")

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306/test_go_db)")
	// var err error
	// database.DBConn, err == gorm.Open("sqlite3", "books.db")
	// if err != nil {
	// 	panic("Failed to connect to database")
	// }
	// fmt.Println("Database connection successfully opened")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println("Database connection successfully opened")

	setupRoutes(app)

	app.Listen(":3000")
}
