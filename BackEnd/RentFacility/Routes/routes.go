package routes

import (
	controllers "rentFacility/Controllers"

	"github.com/gofiber/fiber/v2"
)

func Routing(app *fiber.App) {
	rent := app.Group("rentFacility")
	rent.Get("/", controllers.GetAllRentFacility)
	rent.Post("/create", controllers.CreateRentFacility)
	rent.Put("/update/:id", controllers.UpdateRoomBookingByUser)
	rent.Delete("/delete/:id", controllers.DeleteRentFacility)
}
