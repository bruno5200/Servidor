package api

import (
	"github.com/bruno5200/Servidor/login"
	
	"github.com/gofiber/fiber/v2"
)

func RoutesUp(app *fiber.App) {
	app.Post("/v1/login", login.Login)
}
