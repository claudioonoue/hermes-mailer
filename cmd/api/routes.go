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

var routes = []route{
	{
		Method:  http.MethodGet,
		Path:    "/",
		Handler: checkAPI,
	},
	{
		Method:  http.MethodPost,
		Path:    "/mail",
		Handler: sendMail,
	},
}
