package main

import (
	"net/http"

	"github.com/kaushikkumarbora/TurnedIn/controller"
	"github.com/kaushikkumarbora/TurnedIn/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	db := storage.NewDB()
	defer db.Close()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", landing)
	e.GET("/users", controller.GetUsers)
	e.GET("/login", controller.GetUsers)
	e.GET("/search", controller.GetUsers)

	e.DELETE("/delete", controller.GetUsers)

	e.PUT("/resume", controller.GetUsers)

	e.POST("/connect", controller.GetUsers)
	e.POST("/update", controller.GetUsers)
	e.POST("/signup", controller.GetUsers)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func landing(c echo.Context) error {
	return c.String(http.StatusOK, "TurnedIn")
}
