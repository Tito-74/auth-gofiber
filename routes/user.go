package routes

import (
	database "JWT-GoFiber/Database"
	"JWT-GoFiber/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)




func CreateRoles(c *fiber.Ctx) error {
db := database.Database.Db
var roles models.Roles
if err := c.BodyParser(&roles); err != nil {
	return c.Status(400).JSON("Bad request")
}

	db.Create(&roles)
	// database.Database.Db.Create(&roles)
	return c.JSON(roles)
}

func GetRoles(c *fiber.Ctx)error{
	roles :=[]models.Roles{}

	database.Database.Db.Find(&roles)

	return c.JSON(roles)
	
}




func CreatePermission(c *fiber.Ctx) error {
	db := database.Database.Db
	var permission models.Permission
	if err := c.BodyParser(&permission); err != nil {
		return c.Status(400).JSON("Bad request")
	}
	
		db.Create(&permission)
		// database.Database.Db.Create(&roles)
		return c.JSON(permission)
	}
	
	func GetPermission(c *fiber.Ctx)error{
		permission :=[]models.Permission{}
	
		database.Database.Db.Preload("Role").Find(&permission)
	
		return c.JSON(permission)
		
	}

	func FetchRolesForUser(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.JSON(err)
		}
		permission :=[]models.Permission{}

		database.Database.Db.Preload("Role").Find(&permission, "user_id =?", id)
		// for i :=0; i<=permission.Len(); i++ {
			var roleint []uint
			for j :=0; j< len(permission); j++ {
				fmt.Println(permission[j].Role.Id)
				roleint = append(roleint, permission[j].Role.Id)

			fmt.Println("hello ", roleint)
		}
			

			

		
		return c.Status(200).JSON(permission)

	}