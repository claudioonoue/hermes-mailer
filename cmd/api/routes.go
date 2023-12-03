package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type route struct {
	Method  string
	Path    string
	Handler func(c *fiber.Ctx) error
}

func (a *App) GetRoutes() []route {
	return []route{
		{
			Method:  http.MethodGet,
			Path:    "/",
			Handler: a.CheckAPI,
		},
		{
			Method:  http.MethodPost,
			Path:    "/mail",
			Handler: a.SendMail,
		},
	}
}
