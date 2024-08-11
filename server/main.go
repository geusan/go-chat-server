package main

import (
	"flag"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"api-server/auth"
	"api-server/chat"
	"api-server/internal/repository/rdb"

	"api-server/internal/rest"

	_ "api-server/docs"

	localMiddleware "api-server/internal/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var address = flag.String("addr", ":8080", "http service address")

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v1
func main() {
	flag.Parse()
	db, err := gorm.Open(sqlite.Open("sqlite3.db"), &gorm.Config{})
	if err != nil {
		logrus.Error("error in DB", err)
		return
	}

	userRepo := rdb.NewUserRepository(db)
	chatroomRepo := rdb.NewChatroomRepository(db)
	chatService := chat.NewChatService(userRepo, chatroomRepo)
	authService := auth.NewAuthService(userRepo)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "./templates")
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	v1 := e.Group("/v1")
	rest.NewAuthHandler(v1, authService)
	localMiddleware.UseAuthMiddleware(v1)
	rest.NewChatroomHandler(v1, chatService)

	e.Logger.Fatal(e.Start(*address))
}
