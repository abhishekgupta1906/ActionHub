package main

import (
	"context"
	"fmt"
	"log"
	"os"
	

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Action struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

var collection *mongo.Collection

func main() {
	fmt.Println("Server,Running")
	if(os.Getenv("ENV")!="production"){
		// cannot load env file in production environment
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	MONGODB_URL := os.Getenv("MONGODB_URL")
	//  connect to mongodb
	clientoptions := options.Client().ApplyURI(MONGODB_URL)
	client, err := mongo.Connect(context.Background(), clientoptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())
	
	err = client.Ping(context.Background(), nil)
	fmt.Println("Connected to MongoDB")
	collection = client.Database("actionhub").Collection("actions")

	app := fiber.New()
	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins:     "http://localhost:5173",
	// 	AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
	// 	AllowHeaders:     "Origin, Content-Type, Accept",
       
        
	// }))

	app.Get("/api/actions", getActions)
	app.Post("/api/actions", createAction)
	app.Patch("/api/actions/:id", updateAction)
	app.Delete("/api/actions/:id", deleteAction)
	port := os.Getenv("PORT")
	if port == "" {
		port="3000"
	}
	if(os.Getenv("ENV")=="production"){
		app.Static("/", "./client/dist")
	}
	log.Fatal(app.Listen("0.0.0.0:" + port))
}

func getActions(c *fiber.Ctx) error {
	var actions []Action
	cursor, err := collection.Find(context.Background(), bson.M{})
	// bson.M is used for filtering
	if err != nil {
		c.Status(500).SendString(err.Error())
		return err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		// cursor collection pr iterate kr rha h
		var action Action
		if err := cursor.Decode(&action); err != nil {
			return err
		}

		actions = append(actions, action)
	}
	return c.JSON(actions)
}
func createAction(c *fiber.Ctx) error {
	action := new(Action)
	if err := c.BodyParser(action); err != nil {
		return err
	}
	if action.Body == "" {
		c.Status(400).JSON(fiber.Map{"error": "Body cannot be empty"})
		return nil
	}
	insertResult, err := collection.InsertOne(context.Background(), action)
	if err != nil {

	}
	action.ID = insertResult.InsertedID.(primitive.ObjectID)
	return c.Status(201).JSON(action)
}

func updateAction(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	filter := bson.M{"_id": objectID}

	var action Action
	err = collection.FindOne(context.Background(), filter).Decode(&action)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Action not found"})
	}

	update := bson.M{"$set": bson.M{"completed": !action.Completed}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update action"})
	}

	return c.JSON(fiber.Map{"message": "Action updated successfully"})
}

func deleteAction(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	filter := bson.M{"_id": objectID}
	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Action deleted successfully",
		"success": true,
	})
}
