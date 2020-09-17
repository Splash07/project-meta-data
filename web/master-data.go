package web

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	ghandler "gitlab.ghn.vn/online/common/gstuff/handler"
	"gitlab.com/Splash07/project-meta-data/web/route"
)

// MasterData ..
var MasterData gMasterData

// gMasterData consumer
type gMasterData struct{}

func (gMasterData) Start() error {

	// Init db
	cfg.MongoV2.Get("master-data").Init()

	// Echo instance
	e := echo.New()
	e.Validator = ghandler.NewValidator()
	e.HTTPErrorHandler = ghandler.Error

	// Middlewares
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())

	// Routes => handler
	route.MasterData(e)

	// Start server
	go func() {
		if err := e.Start(":" + cfg.Port["master-data"]); err != nil {
			log.Println("⇛ shutting down the server")
			log.Println(fmt.Sprintf("%v\n", err.Error()))
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM)
	signal.Notify(quit, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	return nil
}
