package main

import (
	"fmt"
	stdlog "log"
	"os"
	"ps-halo-suster/cmd/api/server"
	"ps-halo-suster/configs"
	imagehandler "ps-halo-suster/internal/image/handler"
	imageservice "ps-halo-suster/internal/image/service"
	medicalhandler "ps-halo-suster/internal/medical/handler"
	patientrepository "ps-halo-suster/internal/medical/patient/repository"
	patientservice "ps-halo-suster/internal/medical/patient/service"
	recordrepository "ps-halo-suster/internal/medical/record/repository"
	recordservice "ps-halo-suster/internal/medical/record/service"
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
	params         map[string]string
	baseHandler    *bhandler.BaseHTTPHandler
	userHandler    *userhandler.UserHandler
	medicalHandler *medicalhandler.MedicalHandler
	imageHandler   *imagehandler.ImageHandler
	cfg            *configs.MainConfig
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
		medicalHandler,
		imageHandler,
		port,
	)

	return httpServer.Run()
}

func dbInitConnection() *sqlx.DB {
	return psqlqgen.Init(cfg)
}

func initInfra() {
	cfg = configs.Init()
	db := dbInitConnection()

	userRepository := userrepository.NewUserRepositoryImpl(db)
	userService := userservice.NewUserServiceImpl(userRepository)
	userHandler = userhandler.NewUserHandler(userService)
	imageService := imageservice.NewImageService(cfg)
	imageHandler = imagehandler.NewImageHandler(imageService)

	patientRepository := patientrepository.NewMedicalPatientRepositoryImpl(db)
	recordRepository := recordrepository.NewMedicalRecordRepositoryImpl(db)
	patientService := patientservice.NewMedicalPatientServiceImpl(patientRepository)
	recordService := recordservice.NewMedicalRecordServiceImpl(userRepository, patientRepository, recordRepository)
	medicalHandler = medicalhandler.NewMedicalHandler(userService, patientService, recordService)

}
