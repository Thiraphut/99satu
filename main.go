package main

import (
	"99living-go/database"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func helloWord(c *fiber.Ctx) error {
	return c.SendString("Hello world")
  }


func initDatabase() {
	err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	
	database.ConnectDB()

}


func main() {

	app := fiber.New()
	initDatabase()
	app.Get("/", helloWord)
	setupRoutes(app)
	app.Listen(":8000")
}
