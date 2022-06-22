package main

import (
	"context"
	"fmt"
	"gym-membership/api"
	"gym-membership/app/modules"
	"gym-membership/config"
	"gym-membership/database"
	"net/http"
	"os"
	"time"

	"log"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

var banner = `

 ██████╗ ██╗   ██╗███╗   ███╗      ███╗   ███╗███████╗███╗   ███╗██████╗ ███████╗██████╗  ██████╗██╗  ██╗██╗██████╗
██╔════╝ ╚██╗ ██╔╝████╗ ████║      ████╗ ████║██╔════╝████╗ ████║██╔══██╗██╔════╝██╔══██╗██╔════╝██║  ██║██║██╔══██╗
██║  ██╗  ╚████╔╝ ██╔████╔██║█████╗██╔████╔██║█████╗  ██╔████╔██║██████╦╝█████╗  ██████╔╝╚█████╗ ███████║██║██████╔╝
██║  ╚██╗  ╚██╔╝  ██║╚██╔╝██║╚════╝██║╚██╔╝██║██╔══╝  ██║╚██╔╝██║██╔══██╗██╔══╝  ██╔══██╗ ╚═══██╗██╔══██║██║██╔═══╝
╚██████╔╝   ██║   ██║ ╚═╝ ██║      ██║ ╚═╝ ██║███████╗██║ ╚═╝ ██║██████╦╝███████╗██║  ██║██████╔╝██║  ██║██║██║
 ╚═════╝    ╚═╝   ╚═╝     ╚═╝      ╚═╝     ╚═╝╚══════╝╚═╝     ╚═╝╚═════╝ ╚══════╝╚═╝  ╚═╝╚═════╝ ╚═╝  ╚═╝╚═╝╚═╝

v1.0.0-alpha 
`

func main() {
	log.Println(banner)

	//load config if available or set to default
	config := config.GetConfig()

	dbCon := database.CreateDatabaseConnection(config)
	defer dbCon.CloseConnection()

	e := echo.New()
	e.Pre(echoMiddleware.RemoveTrailingSlash())
	e.Use(echoMiddleware.CORS())
	e.HideBanner = true

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Let's Get Fit and Healthy")
	})

	controllers := modules.RegisterModules(dbCon, config)
	api.RegistrationPath(e, controllers, config)

	go func() {
		address := fmt.Sprintf("0.0.0.0:%d", config.App.Port)
		if err := e.Start(address); err != nil {
			log.Fatal(err)
			panic(err)
		}
	}()

	quit := make(chan os.Signal)
	// signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
