package server

import "github.com/gofiber/fiber/v2"

// Routes содержит список всех доступных путей
func Routes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!\n")
	})
	app.Get("/tasks", ShowTasks)
	app.Post("/tasks", CreateTask)
	app.Put("/tasks/:id", UpdateTask)
	app.Delete("/tasks/:id", DeleteTask)
}
