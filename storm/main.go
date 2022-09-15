package main

import (
	"log"
	"os"
	controller "storm/controller/fiber"
	"storm/models"
	"storm/repository"
	"storm/services"

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

	db, err := repository.ConnectMysqlDB(config)
	if err != nil {
		log.Fatal("could not load database")
	}

	err = models.MigratePersonalRecords(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	service := services.New(db)

	fiberCtrl := controller.NewFiber()

	fiberCtrl.SetupRoutes(service)

	fiberCtrl.Listen(":8081")

	//service.SetupRoutes(fiberCtrl.App)

	/*
			//fiberController.App.Get("/swagger/*", swagger.HandlerDefault) // default
			//app := fiber.New()
		  //app.Get("/swagger/*", swagger.HandlerDefault) // default

			service := &services.Repository{
				DB: db,
			}

			// app.Use(cors.New(cors.Config{
			// 	AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
			// }))

			service.SetupRoutes(app)

			app.Listen(":8080")
	*/
}
