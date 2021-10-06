package main

import (
	"fmt"

	"example.com/ajax/session/Controllers"
	"example.com/ajax/session/DBManager"
	"github.com/gofiber/fiber/v2"
)

func main() {
	initState := DBManager.InitCollections()

	if initState {
		fmt.Println("[OK]")
	}else{
		fmt.Println("[FAIL]")
		return
	}
	app := fiber.New()
	

	app.Post("/student/new", Controllers.StudentNew)
	app.Post("/student/get_all", Controllers.StudentGetAll)
	app.Post("/student/modify", Controllers.StudentModify)


	app.Listen(":3000")
}