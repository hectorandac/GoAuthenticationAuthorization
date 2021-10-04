package main

import (
	"github.com/go-martini/martini"
	"github.com/hectorandac/AuthenticationAuthorization/controllers/users"
	"github.com/hectorandac/AuthenticationAuthorization/middlewares"
	"github.com/hectorandac/AuthenticationAuthorization/models"
	"github.com/joho/godotenv"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func main() {
	godotenv.Load()
	m := martini.Classic()
	m.Use(middlewares.MongoDB())
	m.Use(middlewares.Session())
	m.Use(render.Renderer())

	m.Post("/users", binding.Json(models.UserSignUpRequest{}), users.SignUp)
	m.Post("/users/token", binding.Json(models.UserSignInRequest{}), users.SignIn)
	m.Get("/users/me", users.UserInfo)

	m.Run()
}
