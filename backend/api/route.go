package api

import (
	"backend/api/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v2"
	"log"
	"os"
)

func Setup(app *fiber.App) {

	// allow cors when in dev mode
	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return os.Getenv("ENVIRONMENT") == "development"
		},
	}))
	// open apis
	app.Post("/api/v1/login", handler.Login)
	// app.Post("/api/v1/register", handler.Register)
	// protected apis
	protected := app.Group("/api/v1/protected")
	protected.Use(jwtware.New(jwtware.Config{
		SigningMethod: "HS512",
		SigningKey:    []byte(os.Getenv("SECRETKEY")),
	}))
	protected.Get("/stats", handler.GetStats)

	// static root
	app.Static("/", "./public")
	app.Get("*", func(c *fiber.Ctx) error {
		return c.SendFile("./public/index.html")
	})
	if err := app.ListenTLS(":8443", "server.crt", "server.key"); err != nil {
		log.Fatalln(err)
	}
}
