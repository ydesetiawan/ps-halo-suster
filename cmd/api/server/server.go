package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	userhandler "ps-halo-suster/internal/user/handler"
	bhandler "ps-halo-suster/pkg/base/handler"
)

type Server struct {
	baseHandler *bhandler.BaseHTTPHandler
	userHandler *userhandler.UserHandler
	echo        *echo.Echo
	port        int
}

func NewServer(
	bHandler *bhandler.BaseHTTPHandler,
	userHandler *userhandler.UserHandler,
	port int,
) Server {
	return Server{
		baseHandler: bHandler,
		userHandler: userHandler,
		echo:        echo.New(),
		port:        port,
	}
}

func (s *Server) Run() error {
	slog.Info(fmt.Sprintf("Starting HTTP server at :%d ...", s.port))
	e := echo.New()

	//e.Validator = &helpers.CustomValidator{Validator: validator.New()}
	//e.HTTPErrorHandler = helpers.ErrorHandler
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Secure())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true}))

	s.setupRouter(e)

	return e.Start(fmt.Sprintf(":%d", s.port))
}
