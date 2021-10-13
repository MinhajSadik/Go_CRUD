package Routes

import (
	"example.com/ajax/session/Controllers"
	"github.com/gofiber/fiber/v2"
)

func StudentRoutes(route fiber.Router) {
	route.Post("/new", Controllers.StudentNew)
	route.Post("/get_all", Controllers.StudentGetAll)
	route.Post("/modify", Controllers.StudentModify)
}
