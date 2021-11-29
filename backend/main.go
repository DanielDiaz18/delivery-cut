package main

import (
	"log"
	"net/http"
	"os"

	"github.com/DanielDiaz18/rappi-cut/backend/data"
	"github.com/DanielDiaz18/rappi-cut/backend/handlers"
	"github.com/joho/godotenv"
	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Panic("Error loading .env file")
	}

	if err := mgm.SetDefaultConfig(nil, os.Getenv("mongoDb"), options.Client().ApplyURI(os.Getenv("mongoUri"))); err != nil {
		log.Panic(err)
	}

	e := echo.New()
	e.Validator = data.NewValidator()
	e.Use(
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}\n",
		}),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{http.MethodOptions, http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		}),
	)

	ur := data.NewUsersRepo()
	pr := data.NewParnerRepo()
	or := data.NewOrderRepo()
	sh := handlers.NewServerHandler(ur, pr, or)

	e.GET("/ws", handlers.NewWebSocketHandler)
	router := e.Group("/api/v1")
	{
		router.POST("/login", sh.Login)
		router.POST("/signup", sh.CreateUser)

		ar := router.Group("")
		{
			config := middleware.JWTConfig{
				Claims:     &data.CustomClaims{},
				SigningKey: []byte(os.Getenv("JWT_SECRET")),
			}
			ar.Use(middleware.JWTWithConfig(config))
			ar.GET("/user", sh.GetUser)
			ar.GET("/parners", sh.GetParners)
			ar.GET("/orders", sh.GetOrders)
			ar.POST("/order", sh.CreateOrder)
			ar.POST("/admin/signup", sh.CreateUser)
		}

	}

	e.Logger.Fatal(e.Start(":8000"))
}
