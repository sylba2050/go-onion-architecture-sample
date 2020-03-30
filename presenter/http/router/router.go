package router

import (
	"github.com/labstack/echo"
	"github.com/sylba2050/go-onion-architecture-sample/presenter/http/handler"
)

// NewRouter Routerの設定を行います.
func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.POST("/users", h.CreateUser)
	e.GET("/users", h.GetUsers)
	e.GET("/users/:id", h.GetUser)
	e.PUT("/users/:id", h.UpdateUser)
	e.DELETE("/users/:id", h.DeleteUser)
}
