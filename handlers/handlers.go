package handlers

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/steelthedev/supra-paints/utils"
)

type AppHandler struct {
	EmailService *utils.SmtpServer
}

func (h AppHandler) IndexHandler(ctx *fiber.Ctx) error {
	context := fiber.Map{}
	return ctx.Render("index", context)
}

func (h AppHandler) AboutHanlder(ctx *fiber.Ctx) error {
	context := fiber.Map{}
	return ctx.Render("about", context)
}

func (h AppHandler) TeamHanlder(ctx *fiber.Ctx) error {
	context := fiber.Map{}
	return ctx.Render("team", context)
}

func (h AppHandler) ServicesHanlder(ctx *fiber.Ctx) error {
	context := fiber.Map{}
	return ctx.Render("services", context)
}

func (h AppHandler) ContactHanlder(ctx *fiber.Ctx) error {
	context := fiber.Map{}
	return ctx.Render("contact", context)
}

func (h AppHandler) SendContactMail(ctx *fiber.Ctx) error {
	fullName := ctx.FormValue("full_name")
	clientMsg := ctx.FormValue("client_msg")
	email := ctx.FormValue("email")
	phone, _ := strconv.Atoi(ctx.FormValue("phone"))
	subject := "You have a new message Notification"

	// check if require field is not empty and the character is greater than 1
	if utils.IsEmpty(email) {
		slog.Info(" Field Email, cannot be empty ")
		return ctx.Redirect("/contact")
	}
	if utils.IsEmpty(phone) {
		slog.Info(" Fields Phone, cannot be empty ")
		return ctx.Redirect("/contact")
	}
	if utils.IsEmpty(clientMsg) {
		slog.Info(" Fields client message, cannot be empty ")
		return ctx.Redirect("/contact")
	}

	// check emaill is valid
	if !utils.IsValidEmail(email) {
		slog.Info("Invalid Email")
		return ctx.Redirect("/contact")
	}

	// send the email from contact  page
	msg := []byte(fmt.Sprintf("You have a new message from %s\nEmail: %s \n Phone %d \n\n Message: %s ", fullName, email, phone, clientMsg))
	to := []string{(os.Getenv("DEFAULT_EMAIL_ADDRESS"))}
	from := string(os.Getenv("EMAIL_HOST_USER"))

	if err := h.EmailService.SendMail(to, from, subject, msg); err != nil {
		return err
	}
	return ctx.Redirect("/contact")
}
