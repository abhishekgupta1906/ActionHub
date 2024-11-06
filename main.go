package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Action struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("Hello, world")
	app := fiber.New()
	actions := []Action{}
	err :=godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	app.Get("/api/actions", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(actions)
	})
	app.Post("/api/actions", func(c *fiber.Ctx) error {
		action := new(Action)
		if err := c.BodyParser(action); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		if action.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Body is required"})
		}
		action.ID = len(actions) + 1
		actions = append(actions, *action)
		return c.Status(201).JSON(action)
	})
	// update action
	// fm.sprint to convert int to string
	app.Patch("/api/actions/:id", func(c *fiber.Ctx) error {
		id :=c.Params("id")
		for i,action := range actions{
			if fmt.Sprint(action.ID) == id{
				
				actions[i].Completed = !action.Completed
				return c.Status(200).JSON(actions[i])
			}
		}
		return c.Status(404).JSON(fiber.Map{"error":"Action not found"})

	})
	// delete action
	app.Delete("/api/actions/:id", func(c *fiber.Ctx) error {
		id :=c.Params("id")
		for i,action := range actions{
			if fmt.Sprint(action.ID) == id{
				
				actions=append(actions[:i],actions[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"success":true})	
			}
		}
        return c.Status(404).JSON(fiber.Map{"error":"Action not found"})
	})


	log.Fatal(app.Listen(":"+PORT))
}
