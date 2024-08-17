package app

import (
	"VK-Pilot-Project/internal/config"
	migration "VK-Pilot-Project/internal/migrations/tarantool"
	tarantoolrepo "VK-Pilot-Project/internal/repository/users/tarantool"
	loginservice "VK-Pilot-Project/internal/services/login"
	"VK-Pilot-Project/internal/services/token/jwt"
	"VK-Pilot-Project/internal/transport/handlers/login"
	"VK-Pilot-Project/internal/transport/middleware"
	tarantoolclient "VK-Pilot-Project/pkg/clients/tarantool"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	_ "VK-Pilot-Project/docs" // docs

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// swag init --pd -g internal/app/app.go

// @title Tarantool API
// @description API для tarantool

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name token
func Run() error {
	logger := makeLogger(slog.LevelDebug, os.Stdout)

	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		err := fmt.Errorf("no env for config path")
		logger.Error(err.Error())
		return err
	}

	conf, err := config.New(path)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	logger.Debug("config", slog.Any("struct", conf))

	context := context.Background()

	conn, err := tarantoolclient.NewClient(context, conf.Tarantool)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	defer conn.Close()

	repo, err := tarantoolrepo.New(logger, conn)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	if err := migration.Init(context, repo); err != nil {
		logger.Error(err.Error())
		return err
	}

	mid := middleware.New(logger)

	loginService := loginservice.New(logger, repo)
	tokenService := jwt.New(conf.Tokens.Key)

	loginHandler := login.New(logger, loginService, mid, tokenService)

	router := makeMuxRouter(conf.Server.Port)
	loginHandler.HandleRoute(router)

	server := &http.Server{
		Handler:           router,
		Addr:              fmt.Sprintf("%s:%d", conf.Server.IP, conf.Server.Port),
		ReadHeaderTimeout: 3 * time.Second,
	}

	logger.Debug("server started")

	if err := server.ListenAndServe(); err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func makeLogger(level slog.Level, outFile *os.File) *slog.Logger {
	opt := &slog.HandlerOptions{
		AddSource: true,
		Level:     level,
	}
	return slog.New(slog.NewJSONHandler(outFile, opt))
}

func makeMuxRouter(port int) *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", port)),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	return router
}
