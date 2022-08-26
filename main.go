package main

import (
	"JWT-GoFiber/Database"
	"JWT-GoFiber/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)
// routes.Authorize([]int{2}),
func RoutesSetup(app *fiber.App){
	app.Get("/api/fetch/roles" , routes.GetRoles)
	app.Post("/api/create/roles", routes.CreateRoles)
	app.Get("/api/fetch/permission", routes.GetPermission)
	app.Get("/api/fetch/permission/:id", routes.FetchRolesForUser)
	app.Post("/api/create/permission", routes.CreatePermission)
	app.Post("/api/register", routes.Register)
	app.Post("/api/login", routes.Login)
	
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	// app.Use(cors.New(cors.Config{
	// 	AllowCredentials: true,
	// }))
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,

	}))


	RoutesSetup(app)

	

	log.Fatal(app.Listen(":8000"))
}