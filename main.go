package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cjtim/golang-fiber-mongodb/api"
	"github.com/cjtim/golang-fiber-mongodb/datasource"
	"github.com/cjtim/golang-fiber-mongodb/middlewares"
	"github.com/cjtim/golang-fiber-mongodb/models"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	var m *models.Models
	go func() {
		DBchannel := make(chan *mongo.Client)
		go datasource.MongoClient(DBchannel) // GoRoutine connectDB
		DBclient := <-DBchannel
		m = models.GetModels(DBclient)
		fmt.Println(m.Production.FindAll())
		fmt.Println(m.User.FindAll())
	}()
	log.Fatal(startServer())
}

func startServer() error {
	PORT, isFound := os.LookupEnv("PORT")
	if !isFound {
		PORT = "8080"
	}
	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandling,
	})
	app.Use(middlewares.Cors())
	api.Route(app) // setup router path
	return app.Listen(":" + PORT)
}
