package main

import (
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"github.com/steelthedev/supra-paints/connections"
	"github.com/steelthedev/supra-paints/data"
	"github.com/steelthedev/supra-paints/handlers"
	"github.com/steelthedev/supra-paints/middlewares"
	"github.com/steelthedev/supra-paints/utils"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Info(err.Error())
	}

	viewsEngine := html.New("./static", ".html")

	app := fiber.New(fiber.Config{
		Views:             viewsEngine,
		ViewsLayout:       "layout/main",
		PassLocalsToViews: true,
	})

	app.Static("/static", "./assets")
	app.Use(middlewares.CheckHtmxRequest())

	//smtp handler

	smtpService := utils.SmtpServer{
		Host:     os.Getenv(string("EMAIL_HOST")),
		Port:     465,
		Username: os.Getenv(string("EMAIL_HOST_USER")),
		Password: os.Getenv(string("EMAIL_HOST_PASSWORD")),
	}

	if err != nil {
		slog.Info(err.Error())
	}

	dbURL := os.Getenv("DB_URL")

	db := connections.InitDB(dbURL)
	// Services

	userService := data.NewUserService(db)

	appHandler := handlers.AppHandler{
		EmailService: &smtpService,
	}

	authHandler := handlers.NewAuthHandler(userService)

	app.Get("/", appHandler.IndexHandler)
	app.Get("/about", appHandler.AboutHanlder)
	app.Get("/team", appHandler.TeamHanlder)
	app.Get("/contact", appHandler.ContactHanlder)
	app.Get("/services", appHandler.ServicesHanlder)

	// POST Requests

	app.Post("/send-contact-mail", appHandler.SendContactMail)

	// Auth
	app.Get("/signup", authHandler.CreateNewAdmin)
	app.Listen(":3000")
}
