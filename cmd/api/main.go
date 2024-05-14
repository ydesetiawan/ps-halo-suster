package main

import (
	"fmt"
	stdlog "log"
	"os"
	"ps-halo-suster/cmd/api/server"
	"ps-halo-suster/configs"

	userhandler "ps-halo-suster/internal/user/handler"
	userrepository "ps-halo-suster/internal/user/repository"
	userservice "ps-halo-suster/internal/user/service"

	bhandler "ps-halo-suster/pkg/base/handler"
	"ps-halo-suster/pkg/logger"
	psqlqgen "ps-halo-suster/pkg/psqlqgen"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

var port int

var httpCmd = &cobra.Command{
	Use:   "http [OPTIONS]",
	Short: "Run HTTP API",
	Long:  "Run HTTP API for SCM",
	RunE:  runHttpCommand,
}

var (
	params      map[string]string
	baseHandler *bhandler.BaseHTTPHandler
	userHandler *userhandler.UserHandler
)

func init() {
	httpCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to run the HTTP server")
}

func main() {
	if err := httpCmd.Execute(); err != nil {
		slog.Error(fmt.Sprintf("Error on command execution: %s", err.Error()))
		os.Exit(1)
	}
}

func logLevel() slog.Level {
	switch strings.ToLower(os.Getenv("LOG_LEVEL")) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func initLogger() {
	{

		log, err := logger.SlogOption{
			Resource: map[string]string{
				"service.name":        "halo suster app",
				"service.ns":          "halo suster",
				"service.instance_id": "random-uuid",
				"service.version":     "v.0",
				"service.env":         "staging",
			},
			ContextExtractor:   nil,
			AttributeFormatter: nil,
			Writer:             os.Stdout,
			Leveler:            logLevel(),
		}.NewSlog()
		if err != nil {
			err = fmt.Errorf("prepare logger error: %w", err)
			stdlog.Fatal(err) // if logger cannot be prepared (commonly due to option value error), use std logger.
			return
		}

		// Set logger as global logger.
		slog.SetDefault(log)
	}
}

func runHttpCommand(cmd *cobra.Command, args []string) error {
	initLogger()
	initInfra()

	httpServer := server.NewServer(
		baseHandler,
		userHandler,
		port,
	)

	return httpServer.Run()
}

func dbInitConnection() *sqlx.DB {
	return psqlqgen.Init(configs.Init())
}

func initInfra() {
	db := dbInitConnection()

	userRepository := userrepository.NewUserRepositoryImpl(db)
	userService := userservice.NewUserServiceImpl(userRepository)
	userHandler = userhandler.NewUserHandler(userService)

}
