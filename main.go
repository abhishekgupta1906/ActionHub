package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)
type Action struct {
	ID int `json:"id"`
	Completed bool `json:"completed"`
	Body string `json:"body"`
}
func main() {
	fmt.Println("Hello, world")
	app :=fiber.New()
	actions := []Action{}
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Server Runnning"})
	})
	app.Post("/api/actions", func(c *fiber.Ctx) error {
		action := new(Action)
		if err := c.BodyParser(action); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		if action.Body==""{
			return c.Status(400).JSON(fiber.Map{"error": "Body is required"})
		}
		action.ID = len(actions) + 1
		actions = append(actions, *action)
		return c.Status(201).JSON(action)
	})
	
	log.Fatal(app.Listen(":3000"))
}