package main

import (
	"YeonwooSung/search-and-go/db"
	"YeonwooSung/search-and-go/middlewares"
	"YeonwooSung/search-and-go/routes"
	"YeonwooSung/search-and-go/utils"

	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	env := godotenv.Load()
	if env != nil {
		panic("cannot find environment variables")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = ":4000"
	} else {
		port = ":" + port
	}

	app := fiber.New(fiber.Config{
		IdleTimeout: 5 * time.Second,
	})

	middlewares.SetMiddlewares(app)
	db.InitDB()
	routes.SetRoutes(app)
	utils.StartCronJobs()

	// Start our server and listen for a shutdown
	go func() {
		if err := app.Listen(port); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c // Block the main thread until interupted
	app.Shutdown()
	fmt.Println("shutting down server")
}
