package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"tanahore/configs"
	"tanahore/internal/app"
	"tanahore/internal/infrastructure"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	mysql, err := configs.InitConfig()
	if err != nil {
		logrus.Fatal("error loading congfiguration : ", err.Error())
	}

	db, err := infrastructure.NewMySQLConnection(&mysql.MySQL)
	if err != nil {
		logrus.Fatal("cannot connect to mysql : ", err.Error())
	}

	validate := validator.New()
	e := echo.New()

	app.InitApp(db, validate, e)

	e.GET("/", func(c echo.Context) error {
		file, err := os.ReadFile("./web/static/index.html")
		if err != nil {
			fmt.Println("Error reading file:", err)
			return c.String(http.StatusInternalServerError, "Error reading HTML file")
		}
		return c.HTMLBlob(http.StatusOK, file)
	})

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		},
	))

	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Error starting server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Shutdown Echo gracefully
	if err := e.Shutdown(context.Background()); err != nil {
		logrus.Fatal("Error shutting down server:", err)
	}

	logrus.Info("Server shut down gracefully")
}
