package main

import (
	"github.com/ShijuPJohn/quizapp-backend/routers"
	"github.com/ShijuPJohn/quizapp-backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
)

func main() {
	utils.ConnectDb()
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	routers.SetupRoutes(app)
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}
