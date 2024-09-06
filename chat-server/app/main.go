package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"net/http"

	"chat-server/chat"
	"chat-server/internal/repository/rdb"

	"chat-server/internal/rest"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

var address = flag.String("addr", ":8081", "http service address")

func main() {
	flag.Parse()
	db, err := rdb.OpenDB()
	if err != nil {
		logrus.Error("error in DB", err)
		return
	}

	chatroomRepo := rdb.NewChatRepository(db)
	chatService := chat.NewChatService(chatroomRepo)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.DELETE, echo.PATCH, echo.OPTIONS},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		ExposeHeaders:    []string{echo.HeaderAccessControlAllowCredentials},
		AllowCredentials: true,
	}))

	v1 := e.Group("/v1")
	anonymouseRoute := v1.Group("")

	rest.NewChatroomHandler(anonymouseRoute, chatService)

	// When server is started, register in API server
	println("Register chat server to API server")
	body := struct {
		Url string `json:"url"`
	}{Url: "http://localhost:8081/v1"}
	buf, _ := json.Marshal(body)
	res, _ := http.Post("http://localhost:8080/v1/chat-servers", "application/json", bytes.NewBuffer(buf))
	println("Received", res.StatusCode)
	if res.StatusCode != http.StatusOK {
		panic("Error in system!!!")
	}

	e.Logger.Fatal(e.Start(*address))
}
