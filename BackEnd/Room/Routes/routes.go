package routes

import (
	controllers "room/Controllers"

	"github.com/gofiber/fiber/v2"
)

func Routing(app *fiber.App) {
	kamar := app.Group("Room")
	kamar.Get("/", controllers.GetAllRoom)
	kamar.Post("/create", controllers.CreateRoom)
	kamar.Get("/:id", controllers.GetRoomByID)
	kamar.Put("/update/:id", controllers.UpdateRoom)
	kamar.Delete("/delete/:id", controllers.DeleteRoom)
}
