package Controllers

import (
	"context"
	"encoding/json"

	"example.com/ajax/session/DBManager"
	"example.com/ajax/session/Models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
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

func StudentGetAll(c *fiber.Ctx) error {
	collection := DBManager.SystemCollections.Student

	results:= []bson.M{}

	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.Status(500)
		return err
	}
	defer cur.Close(context.Background())

	cur.All(context.Background(), &results)

	response, _ := json.Marshal(bson.M{
		"results": results,
	})
	c.Set("Content-Type", "application/json")
	c.Status(200).Send(response)

	return nil
}