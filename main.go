package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/zeimedee/goDocker-postgres/database"
	"github.com/zeimedee/goDocker-postgres/handlers"
)

func Handler(app *fiber.App) {
	app.Get("/hello", handlers.Hello)
	app.Get("/all", handlers.All)
	app.Get("/add", handlers.Add)
	app.Get("/update", handlers.Update)
	app.Get("/delete", handlers.Delete)

}

func main() {
	database.ConnectDb()
	app := fiber.New()

	Handler(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
	log.Fatal(app.Listen(":4000"))
}
