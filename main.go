package main

import (
	"fmt"
	"log"

	"github.com/cjtim/golang-fiber-mongodb/api"
	"github.com/cjtim/golang-fiber-mongodb/datasource"
	"github.com/cjtim/golang-fiber-mongodb/datasource/collections"
	"github.com/cjtim/golang-fiber-mongodb/middlewares"
	"github.com/cjtim/golang-fiber-mongodb/models"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	var m *models.Models
	go func() {
		DBchannel := make(chan *mongo.Client)
		go datasource.MongoClient(DBchannel) // GoRoutine connectDB
		m = models.GetModels(<-DBchannel)
		dbExample(m)
	}()
	log.Fatal(startServer().Listen(":8080"))
}

func startServer() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandling,
	})
	app.Use(middlewares.Cors())
	api.Route(app) // setup router path
	return app
}

func dbExample(m *models.Models) {
	newUser := &collections.UserScheama{
		Name: "testInsertUser1",
	}
	fmt.Println("--- InsertOne ---")
	m.InsertOne("users", newUser)
	oneUser := &collections.UserScheama{}
	m.FindOne("users", oneUser, bson.M{"name": newUser.Name})
	fmt.Println(oneUser)

	fmt.Println("--- FindAll --- ")
	// FindAll don't forget Array variable
	allProductions := []collections.ProductionScheama{}
	allUsers := new([]collections.UserScheama)
	allUnitTest := new([]collections.UnitTestSchema)
	m.FindAll("prodcution", allProductions, nil)
	m.FindAll("users", &allUsers, nil)
	m.FindAll("unit_test", &allUnitTest, bson.M{"test": "123456"})
	fmt.Println(allProductions)
	fmt.Println(allUsers)
	fmt.Println(allUnitTest)
	m.Destroy("users", bson.M{
		"name": newUser.Name,
	})
}
