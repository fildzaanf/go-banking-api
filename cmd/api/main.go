package main

import (
	"context"
	"go-banking-api/infrastructure/config"
	"go-banking-api/infrastructure/database"
	ar "go-banking-api/internal/account/router"
	cr "go-banking-api/internal/customer/router"
	tr "go-banking-api/internal/transaction/router"
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
	customer := e.Group("/customers")
	cr.CustomerRouter(customer, db)

	account := e.Group("/accounts")
	ar.AccountRouter(account, db)

	transaction := e.Group("/transactions")
	tr.TransactionRouter(transaction, db)
}

func main() {

	middleware.InitLogger()

	godotenv.Load()
	config, err := config.LoadConfig()
	if err != nil {
		logrus.Fatalf("[ERROR] failed to load configuration: %v", err)
	}

	pdb := database.ConnectPostgreSQL()

	e := echo.New()

	middleware.RemoveTrailingSlash(e)
	e.Use(middleware.Logger) 
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
		logrus.Info("[INFO] Server is running on address ", address)
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
		logrus.Warn("[WARN] Shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := e.Shutdown(ctx); err != nil {
			logrus.Errorf("[ERROR] Error shutting down server: %v", err)
		}
		close(errChan)
	}()

	select {
	case err := <-errChan:
		logrus.Fatalf("[CRITICAL] Server error: %v", err)
	case <-time.After(1 * time.Second):
		logrus.Info("[INFO] Server is running smoothly...")
	}

	wg.Wait()
	logrus.Info("[INFO] Server has been shut down gracefully.")
}
