package main

import (
	"net/http"

	"github.com/kaushikkumarbora/TurnedIn/controller"
	"github.com/kaushikkumarbora/TurnedIn/model"
	"github.com/kaushikkumarbora/TurnedIn/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	db := storage.NewDB()
	defer db.Close()

	//Group
	r := e.Group("/u")
	//Setup JWT Middleware
	config := middleware.JWTConfig{
		Claims:     &model.Auth{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", landing)
	e.GET("/login", controller.Login)          //FORM DATA
	r.GET("/name", controller.GetName)         //QUERY PARAM
	r.GET("/", controller.GetUser)             //TOKEN
	r.GET("/search", controller.GetUsers)      //QUERY PARAM
	r.GET("/resume", controller.GetResume)     //TOKEN
	r.GET("/requests", controller.GetRequests) //TOKEN

	r.DELETE("/delete", controller.DeleteUser) //TOKEN

	r.PUT("/resume", controller.PutResume)        //FORM DATA
	r.PUT("/update", controller.Update)           //FORM DATA
	r.PUT("/accept", controller.AcceptConnection) //QUERY PARAM

	r.POST("/connect", controller.SendConnection) //QUERY PARAM
	e.POST("/signup", controller.Signup)          //FORM DATA

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func landing(c echo.Context) error {
	return c.String(http.StatusOK, "TurnedIn")
}
