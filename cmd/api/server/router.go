package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Server) setupRouter(e *echo.Echo) {
	v1 := e.Group("/v1")
	v1.GET("health", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Health Check OK")
	})
	v1.POST("/user/it/register", s.baseHandler.RunAction(s.userHandler.RegisterIT))
	v1.POST("/user/it/login", s.baseHandler.RunAction(s.userHandler.LoginIT))
	v1.POST("/user/nurse/login", s.baseHandler.RunAction(s.userHandler.LoginNurse))
	v1.POST("/user/nurse/register", s.baseHandler.RunActionAuth(s.userHandler.RegisterNurse))
}
