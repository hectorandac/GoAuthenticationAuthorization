package models

import (
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id                bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name              string        `json:"name" form:"name" binding:"required" bson:"name" validate:"required"`
	Email             string        `json:"email" form:"email" binding:"required" bson:"email" validate:"required"`
	EncryptedPassword string        `json:"ecrypted_password" form:"ecrypted_password" binding:"required" bson:"ecrypted_password"`
}

type UserSignUpRequest struct {
	Name     string `json:"name" form:"name" binding:"required" bson:"name" validate:"required"`
	Email    string `json:"email" form:"email" binding:"required" bson:"email" validate:"required"`
	Password string `json:"password" form:"password" binding:"required" bson:"password" validate:"required"`
}

type UserSignInRequest struct {
	Email    string `json:"email" form:"email" binding:"required" bson:"email" validate:"required"`
	Password string `json:"password" form:"password" binding:"required" bson:"password" validate:"required"`
}

type AuthTokenClaims struct {
	TokenUUID string `json:"tid"`
	UserID    string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"mail"`
	jwt.StandardClaims
}
