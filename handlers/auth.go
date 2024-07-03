package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/steelthedev/supra-paints/data"
	"github.com/steelthedev/supra-paints/types"
	"github.com/sujit-baniya/flash"
)

type AuthHandler struct {
	us *data.UserService
}

func NewAuthHandler(us *data.UserService) *AuthHandler {
	return &AuthHandler{
		us: us,
	}
}

func (a AuthHandler) CreateNewAdmin(ctx *fiber.Ctx) error {

	var params types.UserSignUpRequest

	if err := ctx.BodyParser(&params); err != nil {
		return flash.WithError(ctx, fiber.Map{"message": "Invalid Body Request"}).Redirect("/signup")
	}

	user := data.User{
		Email:    params.Email,
		Password: params.Password,
	}

	if err := a.us.CreateUser(user); err != nil {
		return flash.WithError(ctx, fiber.Map{"message": "Could not sign up"}).Redirect("/signup")
	}

	return ctx.Redirect("/")
}
