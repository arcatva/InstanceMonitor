package main

import (
	"backend/api"
	"backend/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(".env loaded")
}

func main() {
	repository.Connect()
	app := fiber.New()
	api.Setup(app)

}
