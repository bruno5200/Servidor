package main

import (
	"errors"

	"github.com/bruno5200/Servidor/api"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if _, ok := err.(*fiber.Error); ok {
				return errors.New("this is managed error")
			}
			return errors.New("this is managed error")
		},
	})
	//middlewares
	app.Use(recover.New())
	app.Use(logger.New())

	api.RoutesUp(app)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(`FAKE-JWT-TOKEN
/v1/login:

POST:

200:

{
	"id_agencia": 1,
	"usuario": {
		"password": null,
		"username": "admin",
		"authorities": [
			{
				"authority": "ROLE_ADMIN"
			}
		],
		"accountNonExpired": true,
		"accountNonLocked": true,
		"credentialsNonExpired": true,
		"enabled": true
	},
	"recursos": [
		{
			"title": "Configuración",
			"icon": "fa fa-cog",
			"deploymentOrder": 1,
			"resourceLoginQueryDtoList": [
				{
					"title": "Dominios",
					"route": "/domains"
				}
			]
		},
		{
			"title": "Usuarios",
			"icon": "fa fa-cog",
			"deploymentOrder": 2,
			"resourceLoginQueryDtoList": [
				{
					"title": "Usuarios",
					"route": "/users"
				},
				{
					"title": "Roles",
					"route": "/roles"
				},
				{
					"title": "Menús",
					"route": "/resources"
				}
			]
		},
		{
			"title": "Reporte",
			"icon": "fa fa-cog",
			"deploymentOrder": 4,
			"resourceLoginQueryDtoList": [
				{
					"title": "Reportes",
					"route": "/reports"
				}
			]
		}
	],
	"mensaje": "Hola admin, has iniciado sesión con éxito!",
	"telefono": "292292",
	"permisos": [
		{
			"title": "Seleccione el acceso a la información permitida",
			"permissionGroupCode": "PGR01",
			"rolePermissionQueryDtoList": [
				{
					"id": 1,
					"title": "Reporte pdf",
					"permissionCode": "PR001",
					"assign": true
				},
				{
					"id": 2,
					"title": "Reporte csv",
					"permissionCode": "PR002",
					"assign": true
				}
			]
		}
	],
	"nombre": "ISRAEL GUILLERMO QUISPE MAMANI",
	"email": "iquispe@datec.com.bo",
	"agencia": "Oficina Central Banco Union ",
	"token": "eyJhbGciOiJIUzUxMiJ9.eyJhdXRob3JpdGllcyI6Ilt7XCJhdXRob3JpdHlcIjpcIlJPTEVfQURNSU5cIn1dIiwic3ViIjoiYWRtaW4iLCJpYXQiOjE2NTQwMjk2MzUsImV4cCI6MTY1NDA0MzYzNX0.LSgsI0ucjQ-u80KslU7Lmd70ugI1ShTungdhSSKqLtIXWOJb6WrsYpn2RD2bKxF6w2a0Zt-81Sy1YZ5IMclLLw"
}`)
	})
	app.Listen(":3000")
}
