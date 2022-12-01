package main

import (
	"fmt"
	"log"
	"os"
	controllerfiber "storm/controller/fiber"
	"storm/models"
	"storm/repository"
	"storm/services"

	"github.com/gofiber/fiber/v2"

	_ "github.com/arsmn/fiber-swagger/v2/example/docs"

	"github.com/joho/godotenv"
)

// @title Fiber Example API - Storm
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	//db connect
	db, err := repository.ConnectMysqlDB(config)
	if err != nil {
		log.Fatal("could not load database")
	}

	err = models.MigratePersonalRecords(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	//service create
	service := services.New(db)

	//fiber controller engine
	fiberApp := controllerfiber.NewFiber()
	fiberApp.SetupRoutes(service)

	//middleware test
	fiberApp.App.Use(func(c *fiber.Ctx) error {
		fmt.Println("fiber middleware")
		return c.Next()
	})

	// Render index template
	fiberApp.App.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Sports Score Management System!",
			"Name":  "Kyungmun, lim",
		})
	})

	fiberApp.Listen(":8081")

	//gin controller engine
	//ginApp := gin.NewGin()
	//ginApp.SetupRoutes(service)
	//ginApp.Listen(":8082")
}
