package Controllers

import (
	"context"
	"encoding/json"

	"example.com/ajax/session/DBManager"
	"example.com/ajax/session/Models"
	"github.com/gofiber/fiber/v2"
)

func StudentNew(c *fiber.Ctx) error {
	collection := DBManager.SystemCollections.Student

	var self Models.Student
	c.BodyParser(&self)

	err:= self.Validate()
	if err != nil {
		c.Status(500)
		return err
	}

	res, err:= collection.InsertOne(context.Background(), self)
	
	if err != nil {
		c.Status(500)
		return err
	}
	response, _ := json.Marshal(res)
	c.Status(200).Send(response)

	return nil
} 