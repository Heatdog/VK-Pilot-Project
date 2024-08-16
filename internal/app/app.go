package app

import (
	"VK-Pilot-Project/internal/config"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

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

	router := makeMuxRouter(conf.Server.Port)

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
