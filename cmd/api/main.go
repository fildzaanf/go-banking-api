package main

import (
	"context"
	"go-banking-api/infrastructure/database"
	"go-banking-api/infrastructure/config"
	"go-banking-api/pkg/middleware"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	
}

func main() {
	godotenv.Load()
	config, err := config.LoadConfig()
	if err != nil {
		logrus.Fatalf("failed to load configuration: %v", err)
	}

	pdb := database.ConnectPostgreSQL()
	e := echo.New()

	middleware.RemoveTrailingSlash(e)
	middleware.Logger(e)
	middleware.RateLimiter(e)
	middleware.Recover(e)
	middleware.CORS(e)

	SetupRoutes(e, pdb)

	host := config.SERVER.SERVER_HOST
	port := config.SERVER.SERVER_PORT

	address := host + ":" + port

	errChan := make(chan error, 1)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		logrus.Info("server is running on address ", address)
		if err := e.Start(address); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

		<-quit
		logrus.Warn("shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := e.Shutdown(ctx); err != nil {
			logrus.Errorf("error shutting down server: %v", err)
		}
		close(errChan)
	}()

	select {
	case err := <-errChan:
		logrus.Fatalf("server error: %v", err)
	case <-time.After(1 * time.Second):
		logrus.Info("server is running smoothly...")
	}

	wg.Wait()
	logrus.Info("server has been shut down gracefully.")
}
