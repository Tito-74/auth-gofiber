package routes

import (
	database "JWT-GoFiber/Database"
	"JWT-GoFiber/models"
	"errors"

	// "encoding/base64"
	// "encoding/json"
	// "reflect"

	// "encoding/json"
	// "strings"

	// "strconv"
	"fmt"
	"time"

	// "logger"

	"github.com/dgrijalva/jwt-go/v4"
	// "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)



const SecretKey = "secret"

// type typetime time.Time

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}
	database.Database.Db.Create(&user)
	return c.JSON(user)
}
// ##############################################
func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User

	database.Database.Db.Where("email = ?", data["email"]).First(&user)

	if user.ID == 0 {
		c.Status(404)
		fmt.Printf("Wrong user...\n")
		return c.JSON(fiber.Map{
			"Message": "Wrong Credentials",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(400)
		fmt.Printf("Wrong password...\n")
		return c.JSON(fiber.Map{
			"Message": "Wrong Credentials",
		})
		
	}

	

	// var usersP models.UserPayload 

	// permission :=[]models.Permission{}

	// database.Database.Db.Preload("Permissions").Find(&user, "id =?", user.ID)
	// var roleint []uint
	// for j :=0; j< len(user.Permissions); j++ {
	// 	fmt.Println(user.Permissions[j].RoleRefer)
	// 	roleint = append(roleint, user.Permissions[j].RoleRefer)

	// fmt.Println("hello ", roleint)
	// }

	
	// usersP.Name = user.Name
	// usersP.Email = user.Email
	// usersP.Password = user.Password
	// usersP.Roles = roleint





	
	token, err := CreateJwtToken(&user)
	if err != nil {
		return err

	}
// 	}
// 	cookie := fiber.Cookie{
// 	Name:"jwt",
// 	Value: token, 
// 	Expires: time.Unix(0,ExpiresAt),
// 	HTTPOnly: true,
// }
// 	c.Cookie(&cookie)
	
  DecodeToken(token)
	// c.Cookie(cookie)
	// return c.JSON(fiber.Map{
	// 	"Message":"Success"})

	return c.JSON(fiber.Map{"token": token})
	
	

}

func CreateJwtToken(user *models.User) (string, error) {


	database.Database.Db.Preload("Permissions").Find(&user, "id =?", user.ID)
	var roleint []uint
	for j :=0; j< len(user.Permissions); j++ {
		fmt.Println(user.Permissions[j].RoleRefer)
		roleint = append(roleint, user.Permissions[j].RoleRefer)

	fmt.Println("hello ", roleint)
	}
	
	ExpiresAt := time.Now().Add(time.Hour * 24).Unix()

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = user.Email

	claims["expires"] = ExpiresAt

	claims["roles"] = roleint

	t, err:= token.SignedString([]byte(SecretKey))

	if err != nil {

		return "", err
	}

	return t, nil

}


func DecodeToken(tokens string) []int {
	token := tokens
	claims := &models.Claims{}
	jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error){
		return nil, errors.New("invalid token")
	})
	fmt.Println("claims.Roles", claims.Roles[0])
	return claims.Roles
}

// func handler() func(*fiber.Ctx) {
// 	return func(c *fiber.Ctx) {

func Authorize(validRoles []int) func( *fiber.Ctx){
	return func(c *fiber.Ctx) {
		token := c.Query("api")
		roles:= DecodeToken(token)

		// if roles == nil {
		// 	return nil
		// }

		role := roles
		validation := make(map[int]int)
		for _, val := range role {
			validation[val] = 0
		}

		for _, val := range validRoles {
			if _, ok := validation[val]; !ok {
				fmt.Println("error")
				// return nil
			}
		}
		// return nil
	}
}
	// token := tokens

// 	claims,err := token.Claims.(jwt.MapClaims)

	
	
// }

// func VerifyToken(token string)bool{  
// 	claims := &models.UserPayload{}  

// 	jwt.Parse(token, claims, func(token *jwt.Token) (interface{}, error){
		
// 	})

// 		if err != nil {
// 			return false
// 		}

	// token, err := jwt.Parse(token func(token *jwt.Token)(interface{}, error){
	// 	if err != nil {
	// 		return err
	// 	}
		// if _, isValid := token.Method.(*jwt.SigningMethodHS256); !isValid{
		// 	return nil, fmt.Errorf("Unexpected signing method: #{token.Header[alg]}")
		// }
	// })
	
// })
// 	return true
// }

// // var claimsVal []string
// var result = strings.Split(token, ".")
	
// fmt.Println("splited token", result[1])

// decoded, err := base64.RawURLEncoding.DecodeString(result[1])
// 	if err != nil {
// 			panic(err)
// 	}
// 	xt := reflect.TypeOf(decoded).Kind()
// 	fmt.Println("decoded:", string(decoded), xt )

// 	vvv :=  string(decoded)
// 	fmt.Println("data:", vvv[1])



	

