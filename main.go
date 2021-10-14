package main

import (
	"fmt"

	"example.com/ajax/session/DBManager"
	"example.com/ajax/session/Routes"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	Routes.StudentRoutes(app.Group("/student"))

}

func main() {
	initState := DBManager.InitCollections()

	if initState {
		fmt.Println("[OK]")
	} else {
		fmt.Println("[FAIL]")
		return
	}
	app := fiber.New()

	SetupRoutes(app)

	app.Listen(":3000")
}
