package main

import (
	"flag"

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

// @title Go Chatserver Tutorial API
// @version 1.0
// @description 소켓을 이용한 채팅서버 API 문서
// @termsOfService http://swagger.io/terms/

// @contact.name 담당자 마운틴
// @contact.email dnay2k@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v1
func main() {
	flag.Parse()
	db, err := rdb.OpenDB()
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
	anonymouseRoute := v1.Group("")
	authorizedRoute := v1.Group("")
	rest.NewAuthHandler(anonymouseRoute, authService)
	localMiddleware.UseAuthMiddleware(authorizedRoute, authService)
	rest.NewChatroomHandler(authorizedRoute, chatService, authService)

	e.Logger.Fatal(e.Start(*address))
}
