package users

import (
	"net/http"
	"strings"

	"github.com/go-martini/martini"
	"github.com/hectorandac/AuthenticationAuthorization/models"
	"github.com/hectorandac/AuthenticationAuthorization/services"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func UserInfo(params martini.Params, r render.Render, db *mgo.Database, req *http.Request) {
	claim, err := services.GetClaim(req)

	if err != nil {
		json := map[string]interface{}{"error": strings.Split(err.Error(), "\n")}
		r.JSON(400, json)
		return
	}

	user := models.User{}
	oId := bson.ObjectIdHex(claim.UserID)
	db.C("users").FindId(oId).One(&user)

	json := map[string]interface{}{"user": user}
	r.JSON(200, json)
}
