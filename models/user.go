package models

import (
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type User struct{
	ID    	 uint `json:"id" gorm:"primaryKey"`
	Name  	 string `json:"name"`
	Email 	 string `json:"email" gorm:"unique"`
	Password []byte `json:"password"`
	Permissions    []Permission 
}


type Roles struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	
}

type Permission struct{
	ID uint `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time
	RoleRefer uint `json:"role_id"`
	Role  Roles `gorm:"foreignKey:RoleRefer;"`
	UserID uint `json:"user_id"`
}

type Claims struct {
	jwt.StandardClaims
	Roles    []int
}