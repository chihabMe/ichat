package router

import (
	"fmt"

	"github.com/chihabMe/ichat/server/handler"
	"github.com/gofiber/fiber/v2"
)


func SetupAccountsRoutes(app fiber.Router){
	 router := app.Group("/accounts")
	 router.Get("",handler.GetAllAccounts)
	 router.Post("/register",handler.Register)
	 fmt.Println(("regeared accounts routes successfully"))
}