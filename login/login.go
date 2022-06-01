package login

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type login struct {
	IDAgencia int    `json:"id_agencia"`
	Nombre    string `json:"nombre"`
	Email     string `json:"email"`
	Agencia   string `json:"agencia"`
	Token     string `json:"token"`
}
type allLogins []login

var logins = allLogins{
	{
		IDAgencia: 1,
		Nombre:    "ISRAEL GUILLERMO QUISPE MAMANI",
		Email:     "iquispe@datec.com.bo",
		Agencia:   "Oficina Central Banco Union ",
		Token:     "eyJhbGciOiJIUzUxMiJ9.eyJhdXRob3JpdGllcyI6Ilt7XCJhdXRob3JpdHlcIjpcIlJPTEVfQURNSU5cIn1dIiwic3ViIjoiYWRtaW4iLCJpYXQiOjE2NTQwMjk2MzUsImV4cCI6MTY1NDA0MzYzNX0.LSgsI0ucjQ-u80KslU7Lmd70ugI1ShTungdhSSKqLtIXWOJb6WrsYpn2RD2bKxF6w2a0Zt-81Sy1YZ5IMclLLw",
	},
}

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type allUsers []user

var users = allUsers{
	{
		Username: "admin",
		Password: "12345",
	},
}

func Login(c *fiber.Ctx) error {

	var newUser user

	u := new(user)

	if err := c.BodyParser(u); err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success":     false,
			"description": "Request Invalida",
		})
	}

	err := json.Unmarshal(c.Body(), &newUser)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success":     false,
			"description": "Request Invalida",
		})
	}
	if newUser.Username != users[len(users)-1].Username && newUser.Password != users[len(users)-1].Password {
		return c.Status(400).JSON(&fiber.Map{
			"success":     false,
			"description": "Request Invalida",
		})
	} else {
		return c.Status(200).JSON(&fiber.Map{
			"": logins,
		})
	}
}
