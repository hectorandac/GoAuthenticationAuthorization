package middlewares

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/hectorandac/AuthenticationAuthorization/services"
	"gopkg.in/mgo.v2"
)

var ProtectedPaths = []string{
	"/users/me",
}

func Session() martini.Handler {
	return func(c martini.Context, res http.ResponseWriter, req *http.Request, db *mgo.Database) {

		if !contains(ProtectedPaths, req.RequestURI) {
			return
		}

		_, err := services.GetClaim(req)
		if err != nil {
			res.WriteHeader(http.StatusUnauthorized)
		}
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
