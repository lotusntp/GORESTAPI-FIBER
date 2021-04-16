package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/lotusntp/go-fiber/api/book"
	"github.com/lotusntp/go-fiber/database"
)

func helloworld(c *fiber.Ctx)  {
	c.Send("Hello, World")
}

func setupRoutes(app *fiber.App)  {
	app.Get("/api/v1/book",book.GetBooks)
	app.Get("/api/v1/book/:id",book.GetBook)
	app.Post("/api/v1/book",book.NewBook)
	app.Delete("/api/v1/book/:id",book.DeleteBook)
}

func initDatabase()  {

	var err error
	database.DBConn,err = gorm.Open("sqlite3","book.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connection successfully opened")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}


func main()  {

	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()


	setupRoutes(app)

	app.Listen("8080")

}