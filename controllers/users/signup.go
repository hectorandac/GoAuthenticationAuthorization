package users

import (
	"strings"

	"github.com/go-martini/martini"
	"github.com/go-playground/validator"
	"github.com/hectorandac/AuthenticationAuthorization/models"
	"github.com/hectorandac/AuthenticationAuthorization/services"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
)

func SignUp(params martini.Params, userInfo models.UserSignUpRequest, r render.Render, db *mgo.Database) {
	validate = validator.New()
	validationError := validate.Struct(userInfo)

	if validationError != nil {
		json := map[string]interface{}{"error": strings.Split(validationError.Error(), "\n")}
		r.JSON(400, json)
		return
	}

	new_user := models.User{}
	new_user.Name = userInfo.Name
	new_user.Email = userInfo.Email
	new_user.EncryptedPassword, _ = services.HashPassword(userInfo.Password)
	err := db.C("users").Insert(new_user)

	if err != nil {
		json := map[string]interface{}{"error": err}
		r.JSON(400, json)
		return
	}

	json := map[string]interface{}{"user": new_user, "status": "successful"}
	r.JSON(200, json)
}
