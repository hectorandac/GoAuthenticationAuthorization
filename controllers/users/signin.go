package users

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-martini/martini"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/hectorandac/AuthenticationAuthorization/models"
	"github.com/hectorandac/AuthenticationAuthorization/services"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func SignIn(params martini.Params, userInfo models.UserSignInRequest, r render.Render, db *mgo.Database) {
	validate = validator.New()
	validationError := validate.Struct(userInfo)

	if validationError != nil {
		json := map[string]interface{}{"error": strings.Split(validationError.Error(), "\n")}
		r.JSON(400, json)
		return
	}

	user := models.User{}
	err := db.C("users").Find(bson.M{"email": userInfo.Email}).One(&user)
	if err != nil {
		json := map[string]interface{}{"error": "User not found"}
		r.JSON(400, json)
		return
	}

	if !services.CheckPasswordHash(userInfo.Password, user.EncryptedPassword) {
		json := map[string]interface{}{"error": "Wrong credentials"}
		r.JSON(400, json)
		return
	}

	at := models.AuthTokenClaims{
		TokenUUID: uuid.NewString(),
		UserID:    user.Id.Hex(),
		Name:      user.Name,
		Email:     user.Email,
	}

	atoken := jwt.NewWithClaims(jwt.SigningMethodHS256, &at)
	signedAuthToken, tokenErr := atoken.SignedString([]byte("SecretCode"))
	if tokenErr != nil {
		json := map[string]interface{}{"error": tokenErr.Error()}
		r.JSON(400, json)
		return
	}

	json := map[string]interface{}{"user": user, "token": signedAuthToken}
	r.JSON(200, json)

}
